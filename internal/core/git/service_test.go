package git

import (
	"os/exec"
	"strings"
	"testing"
)

func TestService(t *testing.T) {
	s, err := NewService()
	if err != nil {
		t.Fatalf("Failed to create git service: %v", err)
	}

	t.Run("GetCurrentCommit", func(t *testing.T) {
		setupTestRepo(t)

		hash := gitRevParse(t, "HEAD")

		commit, err := s.GetCurrentCommit()
		if err != nil {
			t.Fatalf("GetCurrentCommit failed: %v", err)
		}

		t.Logf("Current commit: %+v", commit)

		if commit.Hash != hash {
			t.Fatalf("Commit hash mismatch: got '%s', want '%s'", commit.Hash, hash)
		}
	})

	t.Run("GetCommit", func(t *testing.T) {
		setupTestRepo(t)

		firstCommit := gitRevParse(t, "HEAD")

		gitCommit(t, "2nd commit")

		secondCommit := gitRevParse(t, "HEAD")

		t.Logf("First commit: %s", firstCommit)
		gitCatFile(t, firstCommit)

		t.Logf("Second commit: %s", secondCommit)
		gitCatFile(t, secondCommit)

		commit, err := s.GetCommit(secondCommit)
		if err != nil {
			t.Fatalf("GetCommit failed: %v", err)
		}

		t.Logf("Commit %s: %+v", secondCommit, commit)

		if commit.Hash != secondCommit {
			t.Fatalf("Commit hash mismatch: got %s, want %s", commit.Hash, secondCommit)
		}
	})

	t.Run("GetCommitFromNonGitRepoShouldFail", func(t *testing.T) {
		path := t.TempDir()
		t.Logf("Non-git repo path: %s", path)

		t.Chdir(path)

		_, err := s.GetCurrentCommit()
		if err == nil {
			t.Fatalf("Expected error when getting commit from non-git repo, got nil")
		}

		t.Logf("Expected error: %v", err)
	})

	t.Run("GetCommitFromEmptyRepoShouldFail", func(t *testing.T) {
		path := t.TempDir()

		t.Chdir(path)

		gitInit(t)

		_, err := s.GetCurrentCommit()
		if err == nil {
			t.Fatalf("Expected error when getting commit from empty git repo, got nil")
		}

		t.Logf("Expected error: %v", err)
	})
}

func setupTestRepo(t *testing.T) string {
	path := t.TempDir()

	t.Log(path)

	t.Chdir(path)

	pwdCmd := exec.Command("pwd")
	out, err := pwdCmd.Output()
	if err != nil {
		t.Fatalf("Failed to get current directory: %v", err)
	}

	t.Logf("Current directory: %s", string(out))

	gitInit(t)
	gitCommit(t, "Initial commit")
	gitLog(t)
	gitCatFile(t, "HEAD")

	return path
}

func gitInit(t *testing.T) {
	// Initialize git repository
	cmd := exec.Command("git", "init")
	out, err := cmd.Output()
	if err != nil {
		t.Fatalf("Failed to initialize git repository: %v", err)
	}

	t.Logf("Git init output: %s", string(out))
}

func gitCommit(t *testing.T, message string) {
	// Create initial commit
	cmd := exec.Command("git", "commit", "--allow-empty", "-m", message)
	out, err := cmd.Output()
	if err != nil {
		t.Fatalf("Failed to create commit: %v", err)
	}

	t.Logf("Git commit output: %s", string(out))
}

func gitCatFile(t *testing.T, ref string) {
	// cat commit info
	cmd := exec.Command("git", "cat-file", "-p", ref)
	out, err := cmd.Output()
	if err != nil {
		t.Fatalf("Failed to get commit info: %v", err)
	}
	t.Logf("Git cat-file output for %s:\n%s", ref, string(out))
}

func gitLog(t *testing.T) {
	// git log
	cmd := exec.Command("git", "log", "--oneline")
	out, err := cmd.Output()
	if err != nil {
		t.Fatalf("Failed to get git log: %v", err)
	}
	t.Logf("Git log:\n%s", string(out))
}

func gitRevParse(t *testing.T, ref string) string {
	// git rev-parse
	cmd := exec.Command("git", "rev-parse", ref)
	out, err := cmd.Output()
	if err != nil {
		t.Fatalf("Failed to get rev-parse for %s: %v", ref, err)
	}
	rev := string(out)
	t.Logf("Git rev-parse output for %s: %s", ref, rev)

	return strings.Trim(rev, " \r\n\t")
}
