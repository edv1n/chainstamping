// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package chainstamper

import (
	"bytes"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = bytes.Equal
	_ = errors.New
	_ = big.NewInt
	_ = common.Big1
	_ = types.BloomLookup
	_ = abi.ConvertType
)

// Commit is an auto generated low-level Go binding around an user-defined struct.
type Commit struct {
	Hash    string
	Tree    string
	Parents []string
}

// ChainstamperMetaData contains all meta data concerning the Chainstamper contract.
var ChainstamperMetaData = bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"hash\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"hash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"tree\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"parents\",\"type\":\"string[]\"}],\"indexed\":false,\"internalType\":\"structCommit\",\"name\":\"commit\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"CommitTimestamped\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"hash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"tree\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"parents\",\"type\":\"string[]\"}],\"internalType\":\"structCommit\",\"name\":\"commit\",\"type\":\"tuple\"}],\"name\":\"getTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"hash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"tree\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"parents\",\"type\":\"string[]\"}],\"internalType\":\"structCommit\",\"name\":\"commit\",\"type\":\"tuple\"}],\"name\":\"stampCommit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	ID:  "f3eb6d43e3676fefffacad454ad2a5ef66",
	Bin: "0x6080604052348015600e575f5ffd5b50610c8d8061001c5f395ff3fe608060405234801561000f575f5ffd5b5060043610610034575f3560e01c806364804d39146100385780636cd66dcc14610068575b5f5ffd5b610052600480360381019061004d9190610309565b610098565b60405161005f9190610368565b60405180910390f35b610082600480360381019061007d9190610309565b6101f1565b60405161008f9190610368565b60405180910390f35b5f5f826100a49061065c565b73__$ce265267f287671425bc4163bedc6df4c2$__63a9a7eff990916040518263ffffffff1660e01b81526004016100dc91906107e4565b602060405180830381865af41580156100f7573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061011b9190610837565b90505f5f5f8381526020019081526020015f20541461016f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610166906108bc565b60405180910390fd5b5f429050805f5f8481526020019081526020015f208190555083805f019061019791906108e6565b6040516101a5929190610976565b6040518091039020827f39087d708d2af7a673faa2d2b6282304c141de8eb5841b2674a1d53c6bf33d4986846040516101df929190610bc1565b60405180910390a38092505050919050565b5f5f826101fd9061065c565b73__$ce265267f287671425bc4163bedc6df4c2$__63a9a7eff990916040518263ffffffff1660e01b815260040161023591906107e4565b602060405180830381865af4158015610250573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906102749190610837565b90505f5f5f8381526020019081526020015f205490505f81036102cc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102c390610c39565b60405180910390fd5b8092505050919050565b5f604051905090565b5f5ffd5b5f5ffd5b5f5ffd5b5f60608284031215610300576102ff6102e7565b5b81905092915050565b5f6020828403121561031e5761031d6102df565b5b5f82013567ffffffffffffffff81111561033b5761033a6102e3565b5b610347848285016102eb565b91505092915050565b5f819050919050565b61036281610350565b82525050565b5f60208201905061037b5f830184610359565b92915050565b5f5ffd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6103cb82610385565b810181811067ffffffffffffffff821117156103ea576103e9610395565b5b80604052505050565b5f6103fc6102d6565b905061040882826103c2565b919050565b5f5ffd5b5f5ffd5b5f5ffd5b5f67ffffffffffffffff82111561043357610432610395565b5b61043c82610385565b9050602081019050919050565b828183375f83830152505050565b5f61046961046484610419565b6103f3565b90508281526020810184848401111561048557610484610415565b5b610490848285610449565b509392505050565b5f82601f8301126104ac576104ab610411565b5b81356104bc848260208601610457565b91505092915050565b5f67ffffffffffffffff8211156104df576104de610395565b5b602082029050602081019050919050565b5f5ffd5b5f610506610501846104c5565b6103f3565b90508083825260208201905060208402830185811115610529576105286104f0565b5b835b8181101561057057803567ffffffffffffffff81111561054e5761054d610411565b5b80860161055b8982610498565b8552602085019450505060208101905061052b565b5050509392505050565b5f82601f83011261058e5761058d610411565b5b813561059e8482602086016104f4565b91505092915050565b5f606082840312156105bc576105bb610381565b5b6105c660606103f3565b90505f82013567ffffffffffffffff8111156105e5576105e461040d565b5b6105f184828501610498565b5f83015250602082013567ffffffffffffffff8111156106145761061361040d565b5b61062084828501610498565b602083015250604082013567ffffffffffffffff8111156106445761064361040d565b5b6106508482850161057a565b60408301525092915050565b5f61066736836105a7565b9050919050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f6106a08261066e565b6106aa8185610678565b93506106ba818560208601610688565b6106c381610385565b840191505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f6107028383610696565b905092915050565b5f602082019050919050565b5f610720826106ce565b61072a81856106d8565b93508360208202850161073c856106e8565b805f5b85811015610777578484038952815161075885826106f7565b94506107638361070a565b925060208a0199505060018101905061073f565b50829750879550505050505092915050565b5f606083015f8301518482035f8601526107a38282610696565b915050602083015184820360208601526107bd8282610696565b915050604083015184820360408601526107d78282610716565b9150508091505092915050565b5f6020820190508181035f8301526107fc8184610789565b905092915050565b5f819050919050565b61081681610804565b8114610820575f5ffd5b50565b5f815190506108318161080d565b92915050565b5f6020828403121561084c5761084b6102df565b5b5f61085984828501610823565b91505092915050565b5f82825260208201905092915050565b7f436f6d6d697420616c72656164792074696d657374616d7065640000000000005f82015250565b5f6108a6601a83610862565b91506108b182610872565b602082019050919050565b5f6020820190508181035f8301526108d38161089a565b9050919050565b5f5ffd5b5f5ffd5b5f5ffd5b5f5f83356001602003843603038112610902576109016108da565b5b80840192508235915067ffffffffffffffff821115610924576109236108de565b5b6020830192506001820236038313156109405761093f6108e2565b5b509250929050565b5f81905092915050565b5f61095d8385610948565b935061096a838584610449565b82840190509392505050565b5f610982828486610952565b91508190509392505050565b5f5ffd5b5f5ffd5b5f5ffd5b5f5f833560016020038436030381126109b6576109b5610996565b5b83810192508235915060208301925067ffffffffffffffff8211156109de576109dd61098e565b5b6001820236038313156109f4576109f3610992565b5b509250929050565b5f82825260208201905092915050565b5f610a1783856109fc565b9350610a24838584610449565b610a2d83610385565b840190509392505050565b5f5f83356001602003843603038112610a5457610a53610996565b5b83810192508235915060208301925067ffffffffffffffff821115610a7c57610a7b61098e565b5b602082023603831315610a9257610a91610992565b5b509250929050565b5f82825260208201905092915050565b5f819050919050565b5f610abf848484610a0c565b90509392505050565b5f602082019050919050565b5f610adf8385610a9a565b935083602084028501610af184610aaa565b805f5b87811015610b36578484038952610b0b828461099a565b610b16868284610ab3565b9550610b2184610ac8565b935060208b019a505050600181019050610af4565b50829750879450505050509392505050565b5f60608301610b595f84018461099a565b8583035f870152610b6b838284610a0c565b92505050610b7c602084018461099a565b8583036020870152610b8f838284610a0c565b92505050610ba06040840184610a38565b8583036040870152610bb3838284610ad4565b925050508091505092915050565b5f6040820190508181035f830152610bd98185610b48565b9050610be86020830184610359565b9392505050565b7f436f6d6d6974206e6f742074696d657374616d706564000000000000000000005f82015250565b5f610c23601683610862565b9150610c2e82610bef565b602082019050919050565b5f6020820190508181035f830152610c5081610c17565b905091905056fea2646970667358221220fa7be3be6073e5b297e4414645f22de366b4981713c42d7a8b9d1e1c94bfe17764736f6c634300081f0033",
	Deps: []*bind.MetaData{
		&CommitLibraryMetaData,
	},
}

// Chainstamper is an auto generated Go binding around an Ethereum contract.
type Chainstamper struct {
	abi abi.ABI
}

// NewChainstamper creates a new instance of Chainstamper.
func NewChainstamper() *Chainstamper {
	parsed, err := ChainstamperMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &Chainstamper{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *Chainstamper) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackGetTimestamp is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6cd66dcc.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getTimestamp((string,string,string[]) commit) view returns(uint256)
func (chainstamper *Chainstamper) PackGetTimestamp(commit Commit) []byte {
	enc, err := chainstamper.abi.Pack("getTimestamp", commit)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetTimestamp is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6cd66dcc.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getTimestamp((string,string,string[]) commit) view returns(uint256)
func (chainstamper *Chainstamper) TryPackGetTimestamp(commit Commit) ([]byte, error) {
	return chainstamper.abi.Pack("getTimestamp", commit)
}

// UnpackGetTimestamp is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x6cd66dcc.
//
// Solidity: function getTimestamp((string,string,string[]) commit) view returns(uint256)
func (chainstamper *Chainstamper) UnpackGetTimestamp(data []byte) (*big.Int, error) {
	out, err := chainstamper.abi.Unpack("getTimestamp", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackStampCommit is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x64804d39.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function stampCommit((string,string,string[]) commit) returns(uint256)
func (chainstamper *Chainstamper) PackStampCommit(commit Commit) []byte {
	enc, err := chainstamper.abi.Pack("stampCommit", commit)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackStampCommit is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x64804d39.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function stampCommit((string,string,string[]) commit) returns(uint256)
func (chainstamper *Chainstamper) TryPackStampCommit(commit Commit) ([]byte, error) {
	return chainstamper.abi.Pack("stampCommit", commit)
}

// UnpackStampCommit is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x64804d39.
//
// Solidity: function stampCommit((string,string,string[]) commit) returns(uint256)
func (chainstamper *Chainstamper) UnpackStampCommit(data []byte) (*big.Int, error) {
	out, err := chainstamper.abi.Unpack("stampCommit", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// ChainstamperCommitTimestamped represents a CommitTimestamped event raised by the Chainstamper contract.
type ChainstamperCommitTimestamped struct {
	Key       [32]byte
	Hash      common.Hash
	Commit    Commit
	Timestamp *big.Int
	Raw       *types.Log // Blockchain specific contextual infos
}

const ChainstamperCommitTimestampedEventName = "CommitTimestamped"

// ContractEventName returns the user-defined event name.
func (ChainstamperCommitTimestamped) ContractEventName() string {
	return ChainstamperCommitTimestampedEventName
}

// UnpackCommitTimestampedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event CommitTimestamped(bytes32 indexed key, string indexed hash, (string,string,string[]) commit, uint256 timestamp)
func (chainstamper *Chainstamper) UnpackCommitTimestampedEvent(log *types.Log) (*ChainstamperCommitTimestamped, error) {
	event := "CommitTimestamped"
	if len(log.Topics) == 0 || log.Topics[0] != chainstamper.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ChainstamperCommitTimestamped)
	if len(log.Data) > 0 {
		if err := chainstamper.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range chainstamper.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// CommitLibraryMetaData contains all meta data concerning the CommitLibrary contract.
var CommitLibraryMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"hash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"tree\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"parents\",\"type\":\"string[]\"}],\"internalType\":\"structCommit\",\"name\":\"commit\",\"type\":\"tuple\"}],\"name\":\"key\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"hash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"tree\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"parents\",\"type\":\"string[]\"}],\"internalType\":\"structCommit\",\"name\":\"commit\",\"type\":\"tuple\"}],\"name\":\"valid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	ID:  "ce265267f287671425bc4163bedc6df4c2",
	Bin: "0x6105b561004d600b8282823980515f1a6073146041577f4e487b71000000000000000000000000000000000000000000000000000000005f525f60045260245ffd5b305f52607381538281f3fe730000000000000000000000000000000000000000301460806040526004361061003f575f3560e01c80638e04beff14610043578063a9a7eff914610073575b5f5ffd5b61005d6004803603810190610058919061017d565b6100a3565b60405161006a91906101de565b60405180910390f35b61008d6004803603810190610088919061017d565b6100dc565b60405161009a919061020f565b60405180910390f35b5f602882805f01906100b59190610234565b90501480156100d5575060288280602001906100d19190610234565b9050145b9050919050565b5f6100e6826100a3565b610125576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161011c906102f0565b60405180910390fd5b81604051602001610136919061055f565b604051602081830303815290604052805190602001209050919050565b5f5ffd5b5f5ffd5b5f5ffd5b5f606082840312156101745761017361015b565b5b81905092915050565b5f6020828403121561019257610191610153565b5b5f82013567ffffffffffffffff8111156101af576101ae610157565b5b6101bb8482850161015f565b91505092915050565b5f8115159050919050565b6101d8816101c4565b82525050565b5f6020820190506101f15f8301846101cf565b92915050565b5f819050919050565b610209816101f7565b82525050565b5f6020820190506102225f830184610200565b92915050565b5f5ffd5b5f5ffd5b5f5ffd5b5f5f833560016020038436030381126102505761024f610228565b5b80840192508235915067ffffffffffffffff8211156102725761027161022c565b5b60208301925060018202360383131561028e5761028d610230565b5b509250929050565b5f82825260208201905092915050565b7f496e76616c696420636f6d6d69740000000000000000000000000000000000005f82015250565b5f6102da600e83610296565b91506102e5826102a6565b602082019050919050565b5f6020820190508181035f830152610307816102ce565b9050919050565b5f5ffd5b5f5ffd5b5f5ffd5b5f5f8335600160200384360303811261033657610335610316565b5b83810192508235915060208301925067ffffffffffffffff82111561035e5761035d61030e565b5b60018202360383131561037457610373610312565b5b509250929050565b5f82825260208201905092915050565b828183375f83830152505050565b5f601f19601f8301169050919050565b5f6103b5838561037c565b93506103c283858461038c565b6103cb8361039a565b840190509392505050565b5f5f833560016020038436030381126103f2576103f1610316565b5b83810192508235915060208301925067ffffffffffffffff82111561041a5761041961030e565b5b6020820236038313156104305761042f610312565b5b509250929050565b5f82825260208201905092915050565b5f819050919050565b5f61045d8484846103aa565b90509392505050565b5f602082019050919050565b5f61047d8385610438565b93508360208402850161048f84610448565b805f5b878110156104d45784840389526104a9828461031a565b6104b4868284610451565b95506104bf84610466565b935060208b019a505050600181019050610492565b50829750879450505050509392505050565b5f606083016104f75f84018461031a565b8583035f8701526105098382846103aa565b9250505061051a602084018461031a565b858303602087015261052d8382846103aa565b9250505061053e60408401846103d6565b8583036040870152610551838284610472565b925050508091505092915050565b5f6020820190508181035f83015261057781846104e6565b90509291505056fea26469706673582212203d3e9d969bcaefdd3fad9c969e515a99c847cb5900d82ad9b116891f07c24f1964736f6c634300081f0033",
}

// CommitLibrary is an auto generated Go binding around an Ethereum contract.
type CommitLibrary struct {
	abi abi.ABI
}

// NewCommitLibrary creates a new instance of CommitLibrary.
func NewCommitLibrary() *CommitLibrary {
	parsed, err := CommitLibraryMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &CommitLibrary{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *CommitLibrary) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackKey is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe0f62495.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function key((string,string,string[]) commit) pure returns(bytes32)
func (commitLibrary *CommitLibrary) PackKey(commit Commit) []byte {
	enc, err := commitLibrary.abi.Pack("key", commit)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackKey is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe0f62495.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function key((string,string,string[]) commit) pure returns(bytes32)
func (commitLibrary *CommitLibrary) TryPackKey(commit Commit) ([]byte, error) {
	return commitLibrary.abi.Pack("key", commit)
}

// UnpackKey is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xe0f62495.
//
// Solidity: function key((string,string,string[]) commit) pure returns(bytes32)
func (commitLibrary *CommitLibrary) UnpackKey(data []byte) ([32]byte, error) {
	out, err := commitLibrary.abi.Unpack("key", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, nil
}

// PackValid is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9a45c3b9.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function valid((string,string,string[]) commit) pure returns(bool)
func (commitLibrary *CommitLibrary) PackValid(commit Commit) []byte {
	enc, err := commitLibrary.abi.Pack("valid", commit)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackValid is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9a45c3b9.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function valid((string,string,string[]) commit) pure returns(bool)
func (commitLibrary *CommitLibrary) TryPackValid(commit Commit) ([]byte, error) {
	return commitLibrary.abi.Pack("valid", commit)
}

// UnpackValid is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x9a45c3b9.
//
// Solidity: function valid((string,string,string[]) commit) pure returns(bool)
func (commitLibrary *CommitLibrary) UnpackValid(data []byte) (bool, error) {
	out, err := commitLibrary.abi.Unpack("valid", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}
