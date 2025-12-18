package stamper

import (
	"context"
	"fmt"
	"os/exec"
	"strings"

	"github.com/edv1n/chainstamping/internal/adapter/stampwa"
	"github.com/edv1n/chainstamping/internal/core/chainstamp"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var defaultRpcURL = "wss://ethereum-sepolia-rpc.publicnode.com"
var defaultContractAddress = "0x07A7859491bfDfD10028C965C8C32f8BE88078B7"
var defaultTxAgentURL = "http://localhost:3000"

func newChainstampService(ctx context.Context) (chainstamp.Service, error) {
	rpcURL, err := getGitConfig("rpcurl")
	if err != nil {
		return nil, fmt.Errorf("failed to get chainstamp.rpcurl: %w", err)
	}
	if rpcURL == "" {
		rpcURL = defaultRpcURL
	}

	ec, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to ethereum client: %w", err)
	}

	agentUrl, err := getGitConfig("txagent")
	if err != nil {
		return nil, fmt.Errorf("failed to get chainstamp.agenturl: %w", err)
	}
	if agentUrl == "" {
		agentUrl = defaultTxAgentURL
	}

	txagent, err := stampwa.NewChainStampTxAgent(agentUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to create tx agent: %w", err)
	}

	contractAddressStr, err := getGitConfig("contractaddress")
	if err != nil {
		return nil, fmt.Errorf("failed to get chainstamp.contractaddress: %w", err)
	}
	if contractAddressStr == "" {
		contractAddressStr = defaultContractAddress
	}

	contractAddress := common.HexToAddress(contractAddressStr)

	chainid, err := ec.ChainID(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get chain ID: %w", err)
	}

	return chainstamp.NewService(ec, chainid, contractAddress, nil, txagent), nil
}

func getGitConfig(key string) (string, error) {
	key = fmt.Sprintf("chainstamp.%s", key)

	// Get git config value
	cmd, err := exec.Command("git", "config", "--get", key).Output()
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok && exitError.ExitCode() == 1 {
			// Key not found, return empty string
			return "", nil
		}

		return "", fmt.Errorf("failed to get git config %s: %w", key, err)
	}

	return strings.Trim(string(cmd), " \r\t\n"), nil
}
