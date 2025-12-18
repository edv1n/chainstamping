package stamper

import (
	"context"
	"fmt"

	"github.com/edv1n/chainstamping/internal/core/chainstamp"
	"github.com/edv1n/chainstamping/internal/core/git"
)

type Stamper interface {
	Stamp(ctx context.Context, commitHash string) error
	Status(ctx context.Context, commitHash string) error
}

type stamper struct {
	cs chainstamp.Service
	gs git.Service
}

func NewStamper() (Stamper, error) {
	gs, err := git.NewService()
	if err != nil {
		return nil, fmt.Errorf("failed to create git service: %w", err)
	}

	cs, err := newChainstampService(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to create chainstamp service: %w", err)
	}

	return &stamper{
		cs: cs,
		gs: gs,
	}, nil
}

func (s stamper) Stamp(ctx context.Context, commitHash string) error {
	commit, err := s.gs.GetCommit(commitHash)
	if err != nil {
		return fmt.Errorf("failed to get commit: %w", err)
	}

	t, err := s.cs.StampCommit(ctx, commit.Hash, commit.Tree, commit.Parents)
	if err != nil {
		return fmt.Errorf("failed to stamp commit: %w", err)
	}

	fmt.Printf("Commit %s successfully stamped at %s (%d)\n", commit.Hash, t.String(), t.Unix())

	return nil
}

func (s stamper) Status(ctx context.Context, commitHash string) error {
	commit, err := s.gs.GetCommit(commitHash)
	if err != nil {
		return fmt.Errorf("failed to get commit: %w", err)
	}

	t, err := s.cs.GetTimestamp(ctx, commit.Hash, commit.Tree, commit.Parents)
	if err != nil {
		return fmt.Errorf("failed to get timestamp: %w", err)
	}

	if t == nil {
		fmt.Printf("Commit %s is not stamped\n", commit.Hash)
		return nil
	}

	fmt.Printf("Commit %s stamped at %s (%d)\n", commit.Hash, t.String(), t.Unix())

	return nil
}
