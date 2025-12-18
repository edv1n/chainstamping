package chainstamp

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type TxAgent interface {
	Timestamp(ctx context.Context, chainId *big.Int, contract common.Address, commitHash string, tree string, parents []string) error
}
