package txagent

import (
	"context"
	"fmt"
	"math/big"
	"net/url"

	"github.com/edv1n/chainstamping/internal/core/chainstamp"
	"github.com/ethereum/go-ethereum/common"
)

type txAgent struct {
	url *url.URL
}

var _ chainstamp.TxAgent = (*txAgent)(nil)

func NewTxAgent(rawurl string) (chainstamp.TxAgent, error) {
	parsedUrl, err := url.Parse(rawurl)
	if err != nil {
		return nil, fmt.Errorf("failed to parse url: %w", err)
	}

	return &txAgent{
		url: parsedUrl,
	}, nil
}

func (ta *txAgent) Timestamp(ctx context.Context, chainId *big.Int, contract common.Address, commitHash string, tree string, parents []string) error {
	u := *ta.url

	u.Path = "/stamp"

	q := u.Query()
	q.Set("contract", contract.Hex())
	q.Set("chain", chainId.String())
	q.Set("hash", commitHash)
	q.Set("tree", tree)
	for _, parent := range parents {
		q.Add("parent", parent)
	}

	u.RawQuery = q.Encode()

	fmt.Printf("Please continue at the following URL:\n%s\n", u.String())

	return nil
}
