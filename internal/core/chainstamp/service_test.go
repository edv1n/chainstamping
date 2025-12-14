package chainstamp

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/edv1n/chainstamping/internal/pkg/chainstamper"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient/simulated"
	"github.com/ethereum/go-ethereum/params"
)

var testChainId = params.AllDevChainProtocolChanges.ChainID

type testTxAgent struct {
	t        *testing.T
	instance *bind.BoundContract
	auth     *bind.TransactOpts
	contract *chainstamper.Chainstamper
	sim      *simulated.Backend
}

func newTestTxAgent(t *testing.T, instance *bind.BoundContract, auth *bind.TransactOpts, sim *simulated.Backend) *testTxAgent {
	return &testTxAgent{
		t:        t,
		instance: instance,
		auth:     auth,
		contract: chainstamper.NewChainstamper(),
		sim:      sim,
	}
}

func (ta *testTxAgent) Timestamp(ctx context.Context, chainId *big.Int, commitHash string, tree string, parents []string) error {
	ta.t.Logf("testTxAgent.Timestamp called with chainId=%s, commitHash=%s, tree=%s, parents=%v", chainId.String(), commitHash, tree, parents)

	tx, err := bind.Transact(ta.instance, ta.auth, ta.contract.PackStampCommit(chainstamper.Commit{
		Hash:    commitHash,
		Tree:    tree,
		Parents: parents,
	}))
	if err != nil {
		ta.t.Logf("testTxAgent failed to Transact: %v", err)

		return fmt.Errorf("failed to Transact: %w", err)
	}

	ta.t.Logf("testTxAgent submitted transaction: %s", tx.Hash().Hex())

	ta.sim.Commit()

	if _, err := bind.WaitMined(ctx, ta.sim.Client(), tx.Hash()); err != nil {
		panic(fmt.Errorf("error waiting for tx inclusion: %v", err))
	}

	ta.t.Logf("testTxAgent transaction mined: %s", tx.Hash().Hex())

	return nil
}

func genHash() string {
	tStr := time.Now().String()
	sum := sha1.Sum([]byte(tStr))
	return hex.EncodeToString(sum[:])
}

func TestService(t *testing.T) {
	sim, contractAddress, auth := newContractSimulator(t, testChainId)

	contract := chainstamper.NewChainstamper()

	instance := contract.Instance(sim.Client(), contractAddress)
	ta := newTestTxAgent(t, instance, auth, sim)

	t.Run("Timestamp", func(t *testing.T) {
		s := NewService(sim.Client(), testChainId, contractAddress, auth, ta)
		commitHash := genHash()
		tree := genHash()
		parents := []string{genHash(), genHash()}

		ts, err := s.StampCommit(t.Context(), commitHash, tree, parents)
		if err != nil {
			for err := err; err != nil; err = errors.Unwrap(err) {
				t.Logf("%T: %+v", err, err)
			}

			t.Fatalf("Timestamp failed: %v", err)
		}
		if ts == nil {
			t.Fatalf("Expected timestamp, got nil")
		}

		t.Logf("Timestamped at: %s", ts.String())
	})

	t.Run("Timestamped", func(t *testing.T) {
		s := NewService(sim.Client(), testChainId, contractAddress, auth, ta)
		commitHash := genHash()
		tree := genHash()
		parents := []string{genHash(), genHash()}

		ts, err := s.StampCommit(t.Context(), commitHash, tree, parents)
		if err != nil {
			for err := err; err != nil; err = errors.Unwrap(err) {
				t.Logf("%T: %+v", err, err)
			}

			t.Fatalf("Timestamp failed: %v", err)
		}

		tsed, err := s.GetTimestamp(t.Context(), commitHash, tree, parents)
		if err != nil {
			t.Fatalf("Timestamped failed: %v", err)
		}
		if tsed == nil {
			t.Fatalf("Expected timestamp, got nil")
		}

		t.Logf("Timestamped at: %s", tsed.String())

		if !tsed.Equal(*ts) {
			t.Fatalf("Timestamps do not match: got %s, want %s", tsed.String(), ts.String())
		}
	})
}

func newContractSimulator(t *testing.T, chainID *big.Int) (sim *simulated.Backend, contractAddress common.Address, auth *bind.TransactOpts) {
	// from https://geth.ethereum.org/docs/developers/dapp-developer/native-bindings-v2#blockchain-simulator
	key, err := crypto.GenerateKey()
	if err != nil {
		t.Fatalf("failed to generate key: %v", err)
	}

	auth = bind.NewKeyedTransactor(key, chainID)

	sim = simulated.NewBackend(map[common.Address]types.Account{
		auth.From: {Balance: big.NewInt(9e18)},
	})

	// set up params to deploy an instance of the Chainstamper contract
	deployParams := bind.DeploymentParams{
		Contracts: []*bind.MetaData{&chainstamper.ChainstamperMetaData},
	}

	// use the default deployer: it simply creates, signs and submits the deployment transactions
	deployer := bind.DefaultDeployer(auth, sim.Client())

	// create and submit the contract deployment
	deployRes, err := bind.LinkAndDeploy(&deployParams, deployer)
	if err != nil {
		t.Fatalf("error submitting contract: %v", err)
	}

	address, tx := deployRes.Addresses[chainstamper.ChainstamperMetaData.ID], deployRes.Txs[chainstamper.ChainstamperMetaData.ID]
	contractAddress = address

	// call Commit to make the simulated backend mine a block
	sim.Commit()

	// wait for the pending contract to be deployed on-chain
	if _, err := bind.WaitDeployed(t.Context(), sim.Client(), tx.Hash()); err != nil {
		t.Fatalf("failed waiting for contract deployment: %v", err)
	}

	t.Logf("contract deployment transaction mined: %s", tx.Hash().Hex())
	t.Logf("contract deployed at address 0x%x\n", address)

	return sim, contractAddress, auth
}
