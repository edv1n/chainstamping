package cmd

import (
	"fmt"

	"github.com/edv1n/chainstamping/internal/app/stamper"
)

func newStamper() (stamper.Stamper, error) {
	s, err := stamper.NewStamper()
	if err != nil {
		return nil, fmt.Errorf("failed to create stamper: %w", err)
	}

	return s, nil
}
