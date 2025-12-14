package chainstamp

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/edv1n/chainstamping/internal/pkg/chainstampingcommits"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
)

type Service interface {
	Timestamp(ctx context.Context, commitHash string, tree string, parents []string) (*time.Time, error)
	Timestamped(ctx context.Context, commitHash string, tree string, parents []string) (*time.Time, error)
}

type service struct {
	contractAddress common.Address
	contract        *chainstampingcommits.ChainstampingCommits
	instance        *bind.BoundContract
	ec              bind.Backend
	auth            *bind.TransactOpts
	ta              TxAgent
	chainId         *big.Int
}

func NewService(ec bind.Backend, chainId *big.Int, contractAddress common.Address, auth *bind.TransactOpts, ta TxAgent) Service {
	contract := chainstampingcommits.NewChainstampingCommits()

	instance := contract.Instance(ec, contractAddress)

	return &service{
		contractAddress: contractAddress,
		contract:        contract,
		instance:        instance,
		ec:              ec,
		auth:            auth,
		ta:              ta,
	}
}

func (s *service) Timestamp(ctx context.Context, commitHash string, tree string, parents []string) (*time.Time, error) {
	ts, err := s.Timestamped(ctx, commitHash, tree, parents)
	if err != nil {
		return nil, fmt.Errorf("failed to get timestamp: %w", err)
	}

	if ts != nil {
		return nil, ErrCommitAlreadyTimestamped
	}

	evc := make(chan *chainstampingcommits.ChainstampingCommitsCommitTimestamped)

	sub, err := bind.WatchEvents(s.instance, &bind.WatchOpts{Context: ctx}, s.contract.UnpackCommitTimestampedEvent, evc)
	if err != nil {
		return nil, fmt.Errorf("failed to watch events: %w", err)
	}
	defer sub.Unsubscribe()

	if err := s.ta.Timestamp(ctx, s.chainId, commitHash, tree, parents); err != nil {
		return nil, fmt.Errorf("failed to timestamp: %w", err)
	}

	for {
		select {
		case ev := <-evc:
			if ev.Commit.Hash == commitHash && ev.Commit.Tree == tree {
				t := time.Unix(ev.Timestamp.Int64(), 0)
				return &t, nil
			}
		case err := <-sub.Err():
			return nil, fmt.Errorf("event subscription error: %w", err)
		case <-ctx.Done():
			return nil, ctx.Err()
		}
	}
}

func (s *service) Timestamped(ctx context.Context, commitHash string, tree string, parents []string) (*time.Time, error) {
	data, err := s.contract.TryPackTimestamped(chainstampingcommits.Commit{
		Hash:    commitHash,
		Tree:    tree,
		Parents: parents,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to pack timestamped call: %w", err)
	}

	resp, err := bind.Call(s.instance, &bind.CallOpts{Context: ctx}, data, s.contract.UnpackTimestamped)
	if err != nil {
		if err.Error() == "execution reverted: Commit not timestamped" {
			return nil, nil
		}

		return nil, fmt.Errorf("failed to call timestamped: %w", err)
	}

	t := time.Unix(resp.Int64(), 0)

	return &t, nil
}
