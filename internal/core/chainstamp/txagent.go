package chainstamp

import (
	"context"
	"math/big"
)

type TxAgent interface {
	Timestamp(ctx context.Context, chainId *big.Int, commitHash string, tree string, parents []string) error
}
