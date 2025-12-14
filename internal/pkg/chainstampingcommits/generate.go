package chainstampingcommits

//go:generate vay generate solidity abi ./ChainstampingCommits
//go:generate go tool github.com/ethereum/go-ethereum/cmd/abigen --v2 --combined-json=./ChainstampingCommits.abi/combined.json --out=generated.go --pkg=chainstampingcommits --type ChainstampingCommits
