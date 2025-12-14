package chainstamper

///go:generate vay generate solidity abi ./ChainstampingCommits
//go:generate vay solc --combined-json abi,bin ./contracts/Chainstamper.sol | go tool github.com/ethereum/go-ethereum/cmd/abigen --v2 --combined-json=- --out=generated.go --type Chainstamper
