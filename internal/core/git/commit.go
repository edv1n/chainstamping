package git

type Commit struct {
	Hash    string
	Tree    string
	Parents []string
}
