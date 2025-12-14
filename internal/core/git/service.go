package git

import (
	"fmt"
	"os/exec"
	"strings"
)

type Service interface {
	GetCurrentCommit() (*Commit, error)
	GetCommit(commitHash string) (*Commit, error)
}

type service struct {
	gitexec string
}

func NewService() (Service, error) {
	var path string
	for _, p := range []string{"git"} {
		var err error
		path, err = exec.LookPath(p)
		if err == nil {
			break
		}
	}
	if path == "" {
		return nil, fmt.Errorf("no git executable found in PATH")
	}

	return &service{
		gitexec: path,
	}, nil
}

func (s *service) GetCurrentCommit() (*Commit, error) {
	commitHash, err := s.getHEADCommitHash()
	if err != nil {
		return nil, fmt.Errorf("failed to get current commit hash: %w", err)
	}

	return s.GetCommit(commitHash)
}

func (s *service) GetCommit(commitHash string) (*Commit, error) {
	tree, err := s.getTreeFromRevParseOutput(commitHash)
	if err != nil {
		return nil, fmt.Errorf("failed to get commit tree: %w", err)
	}

	parents, err := s.getParentsFromRevListOutput(commitHash)
	if err != nil {
		return nil, fmt.Errorf("failed to get commit parents: %w", err)
	}

	return &Commit{
		Hash:    commitHash,
		Tree:    tree,
		Parents: parents,
	}, nil
}

func (s service) getTreeFromRevParseOutput(hash string) (string, error) {
	cmd := exec.Command(s.gitexec, "rev-parse", fmt.Sprintf("%s^{tree}", hash))
	out, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to get commit tree: %w", err)
	}

	return strings.Trim(string(out), " \r\t\n"), nil
}

func (s service) getParentsFromRevListOutput(hash string) ([]string, error) {
	cmd := exec.Command(s.gitexec, "rev-list", "--parents", "-n", "1", hash)
	out, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to get commit parents: %w", err)
	}

	str := strings.Trim(string(out), " \r\t\n")
	ss := strings.Split(str, " ")

	return ss[1:], nil
}

func (s service) getHEADCommitHash() (string, error) {
	cmd := exec.Command(s.gitexec, "rev-parse", "HEAD")
	out, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to get HEAD commit hash: %w", err)
	}

	return strings.Trim(string(out), " \r\t\n"), nil
}
