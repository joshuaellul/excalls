// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package testexcall

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// TestEXCALLABI is the input ABI used to generate the binding from.
const TestEXCALLABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oracleRefNr\",\"type\":\"uint256\"}],\"name\":\"BetPlaced\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"alwaysWinGithub\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"alwaysWinLocal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"beginAlwaysWinOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forPunter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"oracleRef\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"answer\",\"type\":\"bool\"}],\"name\":\"continueAlwaysWinOracle\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"excall\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"get0s\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"get1s\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAAA\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getECT\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oracleAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"pending\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOracleNr\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"response\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"add\",\"type\":\"address\"}],\"name\":\"setOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sometimesWin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"winnings\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// TestEXCALLFuncSigs maps the 4-byte function signature to its string representation.
var TestEXCALLFuncSigs = map[string]string{
	"78e4d58c": "alwaysWinGithub()",
	"d7f8db5c": "alwaysWinLocal()",
	"7163c7d7": "beginAlwaysWinOracle()",
	"d7c0d248": "continueAlwaysWinOracle(address,uint256,bool)",
	"025a061b": "excall()",
	"40d68738": "get0s()",
	"cc832f28": "get1s()",
	"2d3085f6": "getAAA()",
	"a59bffe0": "getECT()",
	"a89ae4ba": "oracleAddress()",
	"15167c03": "pending(address,uint256)",
	"54624615": "pendingOracleNr()",
	"7a7f01a7": "response()",
	"7adbf973": "setOracle(address)",
	"09554752": "sometimesWin()",
	"ea3a1499": "winnings(address)",
}

// TestEXCALLBin is the compiled bytecode used for deploying new contracts.
var TestEXCALLBin = "0x608060405234801561001057600080fd5b50610a14806100206000396000f3fe608060405234801561001057600080fd5b50600436106101005760003560e01c80637a7f01a711610097578063cc832f2811610066578063cc832f281461020d578063d7c0d24814610215578063d7f8db5c14610228578063ea3a14991461023057600080fd5b80637a7f01a7146101a25780637adbf973146101aa578063a59bffe0146101da578063a89ae4ba146101e257600080fd5b806340d68738116100d357806340d6873814610173578063546246151461017b5780637163c7d71461019257806378e4d58c1461019a57600080fd5b8063025a061b14610105578063095547521461012357806315167c031461012d5780632d3085f61461016b575b600080fd5b61010d610250565b60405161011a91906108c9565b60405180910390f35b61012b6102de565b005b61015b61013b36600461080d565b600360209081526000928352604080842090915290825290205460ff1681565b604051901515815260200161011a565b61010d6103cf565b61010d610419565b61018460055481565b60405190815260200161011a565b61012b61045f565b61012b6104f0565b61010d610532565b61012b6101b83660046107f2565b600280546001600160a01b0319166001600160a01b0392909216919091179055565b61010d61053f565b6002546101f5906001600160a01b031681565b6040516001600160a01b03909116815260200161011a565b61010d610585565b61015b610223366004610837565b6105cb565b61012b610708565b61018461023e3660046107f2565b60046020526000908152604090205481565b6001805461025d9061095c565b80601f01602080910402602001604051908101604052809291908181526020018280546102899061095c565b80156102d65780601f106102ab576101008083540402835291602001916102d6565b820191906000526020600020905b8154815290600101906020018083116102b957829003601f168201915b505050505081565b6040805180820190915260208082527f687474703a2f2f6c6f63616c686f73743a383038302f657863616c6c72616e6491810191825261032091600191610746565b506001600081546103309061095c565b811061033e5761033e6109c8565b81546001161561035d5790600052602060002090602091828204019190065b9054901a600160f81b026001600160f81b031916603160f81b14156103a75733600090815260046020526040812080546001929061039c9084906108dc565b909155506103cd9050565b3360009081526004602052604081208054600192906103c790849061091d565b90915550505b565b6040805180820190915260208082527f687474703a2f2f6a6f73687561656c6c756c2e6769746875622e696f2f6161619181019182526060916104159160009190610746565b5090565b6040805180820190915260208082527f687474703a2f2f6c6f63616c686f73743a383038302f657863616c6c307374729181019182526060916104159160009190610746565b33600090815260036020526040812060058054600193918261048083610997565b91905055815260200190815260200160002060006101000a81548160ff021916908315150217905550336001600160a01b03167f43e08d78302cdf9c94e1ffd293c2a3696139ee41cf3199b6f5eed9c7f6cc60906005546040516104e691815260200190565b60405180910390a2565b6040805180820190915260208082527f687474703a2f2f6a6f73687561656c6c756c2e6769746875622e696f2f77696e91810191825261032091600191610746565b6000805461025d9061095c565b6040805180820190915260208082527f687474703a2f2f6a6f73687561656c6c756c2e6769746875622e696f2f6563749181019182526060916104159160009190610746565b6040805180820190915260208082527f687474703a2f2f6c6f63616c686f73743a383038302f657863616c6c317374729181019182526060916104159160009190610746565b6002546000906001600160a01b0316331461062c5760405162461bcd60e51b815260206004820152601760248201527f596f7520617265206e6f7420746865206f7261636c6521000000000000000000604482015260640160405180910390fd5b6001600160a01b038416600090815260036020908152604080832086845290915290205460ff1680156106fb5760018315151415610698576001600160a01b038516600090815260046020526040812080546001929061068d9084906108dc565b909155506106c79050565b6001600160a01b03851660009081526004602052604081208054600192906106c190849061091d565b90915550505b50506001600160a01b03831660009081526003602090815260408083208584529091529020805460ff191690556001610701565b60009150505b9392505050565b6040805180820190915260208082527f687474703a2f2f6c6f63616c686f73743a383038302f657863616c6c31737472918101918252610320916001915b8280546107529061095c565b90600052602060002090601f01602090048101928261077457600085556107ba565b82601f1061078d57805160ff19168380011785556107ba565b828001600101855582156107ba579182015b828111156107ba57825182559160200191906001019061079f565b506104159291505b8082111561041557600081556001016107c2565b80356001600160a01b03811681146107ed57600080fd5b919050565b60006020828403121561080457600080fd5b610701826107d6565b6000806040838503121561082057600080fd5b610829836107d6565b946020939093013593505050565b60008060006060848603121561084c57600080fd5b610855846107d6565b9250602084013591506040840135801515811461087157600080fd5b809150509250925092565b6000815180845260005b818110156108a257602081850181015186830182015201610886565b818111156108b4576000602083870101525b50601f01601f19169290920160200192915050565b602081526000610701602083018461087c565b600080821280156001600160ff1b03849003851316156108fe576108fe6109b2565b600160ff1b8390038412811615610917576109176109b2565b50500190565b60008083128015600160ff1b85018412161561093b5761093b6109b2565b6001600160ff1b0384018313811615610956576109566109b2565b50500390565b600181811c9082168061097057607f821691505b6020821081141561099157634e487b7160e01b600052602260045260246000fd5b50919050565b60006000198214156109ab576109ab6109b2565b5060010190565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052603260045260246000fdfea26469706673582212208ec0ba924edfdebc3c74f693a0d2d4d9fb0bd0891e311c4d63ec876f535a9f9064736f6c63430008060033"

// DeployTestEXCALL deploys a new Ethereum contract, binding an instance of TestEXCALL to it.
func DeployTestEXCALL(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TestEXCALL, error) {
	parsed, err := abi.JSON(strings.NewReader(TestEXCALLABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TestEXCALLBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TestEXCALL{TestEXCALLCaller: TestEXCALLCaller{contract: contract}, TestEXCALLTransactor: TestEXCALLTransactor{contract: contract}, TestEXCALLFilterer: TestEXCALLFilterer{contract: contract}}, nil
}

// TestEXCALL is an auto generated Go binding around an Ethereum contract.
type TestEXCALL struct {
	TestEXCALLCaller     // Read-only binding to the contract
	TestEXCALLTransactor // Write-only binding to the contract
	TestEXCALLFilterer   // Log filterer for contract events
}

// TestEXCALLCaller is an auto generated read-only Go binding around an Ethereum contract.
type TestEXCALLCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestEXCALLTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TestEXCALLTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestEXCALLFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TestEXCALLFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestEXCALLSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestEXCALLSession struct {
	Contract     *TestEXCALL       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TestEXCALLCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestEXCALLCallerSession struct {
	Contract *TestEXCALLCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// TestEXCALLTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestEXCALLTransactorSession struct {
	Contract     *TestEXCALLTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// TestEXCALLRaw is an auto generated low-level Go binding around an Ethereum contract.
type TestEXCALLRaw struct {
	Contract *TestEXCALL // Generic contract binding to access the raw methods on
}

// TestEXCALLCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestEXCALLCallerRaw struct {
	Contract *TestEXCALLCaller // Generic read-only contract binding to access the raw methods on
}

// TestEXCALLTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestEXCALLTransactorRaw struct {
	Contract *TestEXCALLTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTestEXCALL creates a new instance of TestEXCALL, bound to a specific deployed contract.
func NewTestEXCALL(address common.Address, backend bind.ContractBackend) (*TestEXCALL, error) {
	contract, err := bindTestEXCALL(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TestEXCALL{TestEXCALLCaller: TestEXCALLCaller{contract: contract}, TestEXCALLTransactor: TestEXCALLTransactor{contract: contract}, TestEXCALLFilterer: TestEXCALLFilterer{contract: contract}}, nil
}

// NewTestEXCALLCaller creates a new read-only instance of TestEXCALL, bound to a specific deployed contract.
func NewTestEXCALLCaller(address common.Address, caller bind.ContractCaller) (*TestEXCALLCaller, error) {
	contract, err := bindTestEXCALL(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestEXCALLCaller{contract: contract}, nil
}

// NewTestEXCALLTransactor creates a new write-only instance of TestEXCALL, bound to a specific deployed contract.
func NewTestEXCALLTransactor(address common.Address, transactor bind.ContractTransactor) (*TestEXCALLTransactor, error) {
	contract, err := bindTestEXCALL(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestEXCALLTransactor{contract: contract}, nil
}

// NewTestEXCALLFilterer creates a new log filterer instance of TestEXCALL, bound to a specific deployed contract.
func NewTestEXCALLFilterer(address common.Address, filterer bind.ContractFilterer) (*TestEXCALLFilterer, error) {
	contract, err := bindTestEXCALL(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestEXCALLFilterer{contract: contract}, nil
}

// bindTestEXCALL binds a generic wrapper to an already deployed contract.
func bindTestEXCALL(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TestEXCALLABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestEXCALL *TestEXCALLRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestEXCALL.Contract.TestEXCALLCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestEXCALL *TestEXCALLRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestEXCALL.Contract.TestEXCALLTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestEXCALL *TestEXCALLRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestEXCALL.Contract.TestEXCALLTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestEXCALL *TestEXCALLCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestEXCALL.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestEXCALL *TestEXCALLTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestEXCALL.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestEXCALL *TestEXCALLTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestEXCALL.Contract.contract.Transact(opts, method, params...)
}

// Excall is a free data retrieval call binding the contract method 0x025a061b.
//
// Solidity: function excall() view returns(bytes)
func (_TestEXCALL *TestEXCALLCaller) Excall(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _TestEXCALL.contract.Call(opts, &out, "excall")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// Excall is a free data retrieval call binding the contract method 0x025a061b.
//
// Solidity: function excall() view returns(bytes)
func (_TestEXCALL *TestEXCALLSession) Excall() ([]byte, error) {
	return _TestEXCALL.Contract.Excall(&_TestEXCALL.CallOpts)
}

// Excall is a free data retrieval call binding the contract method 0x025a061b.
//
// Solidity: function excall() view returns(bytes)
func (_TestEXCALL *TestEXCALLCallerSession) Excall() ([]byte, error) {
	return _TestEXCALL.Contract.Excall(&_TestEXCALL.CallOpts)
}

// OracleAddress is a free data retrieval call binding the contract method 0xa89ae4ba.
//
// Solidity: function oracleAddress() view returns(address)
func (_TestEXCALL *TestEXCALLCaller) OracleAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TestEXCALL.contract.Call(opts, &out, "oracleAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OracleAddress is a free data retrieval call binding the contract method 0xa89ae4ba.
//
// Solidity: function oracleAddress() view returns(address)
func (_TestEXCALL *TestEXCALLSession) OracleAddress() (common.Address, error) {
	return _TestEXCALL.Contract.OracleAddress(&_TestEXCALL.CallOpts)
}

// OracleAddress is a free data retrieval call binding the contract method 0xa89ae4ba.
//
// Solidity: function oracleAddress() view returns(address)
func (_TestEXCALL *TestEXCALLCallerSession) OracleAddress() (common.Address, error) {
	return _TestEXCALL.Contract.OracleAddress(&_TestEXCALL.CallOpts)
}

// Pending is a free data retrieval call binding the contract method 0x15167c03.
//
// Solidity: function pending(address , uint256 ) view returns(bool)
func (_TestEXCALL *TestEXCALLCaller) Pending(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (bool, error) {
	var out []interface{}
	err := _TestEXCALL.contract.Call(opts, &out, "pending", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Pending is a free data retrieval call binding the contract method 0x15167c03.
//
// Solidity: function pending(address , uint256 ) view returns(bool)
func (_TestEXCALL *TestEXCALLSession) Pending(arg0 common.Address, arg1 *big.Int) (bool, error) {
	return _TestEXCALL.Contract.Pending(&_TestEXCALL.CallOpts, arg0, arg1)
}

// Pending is a free data retrieval call binding the contract method 0x15167c03.
//
// Solidity: function pending(address , uint256 ) view returns(bool)
func (_TestEXCALL *TestEXCALLCallerSession) Pending(arg0 common.Address, arg1 *big.Int) (bool, error) {
	return _TestEXCALL.Contract.Pending(&_TestEXCALL.CallOpts, arg0, arg1)
}

// PendingOracleNr is a free data retrieval call binding the contract method 0x54624615.
//
// Solidity: function pendingOracleNr() view returns(uint256)
func (_TestEXCALL *TestEXCALLCaller) PendingOracleNr(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TestEXCALL.contract.Call(opts, &out, "pendingOracleNr")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingOracleNr is a free data retrieval call binding the contract method 0x54624615.
//
// Solidity: function pendingOracleNr() view returns(uint256)
func (_TestEXCALL *TestEXCALLSession) PendingOracleNr() (*big.Int, error) {
	return _TestEXCALL.Contract.PendingOracleNr(&_TestEXCALL.CallOpts)
}

// PendingOracleNr is a free data retrieval call binding the contract method 0x54624615.
//
// Solidity: function pendingOracleNr() view returns(uint256)
func (_TestEXCALL *TestEXCALLCallerSession) PendingOracleNr() (*big.Int, error) {
	return _TestEXCALL.Contract.PendingOracleNr(&_TestEXCALL.CallOpts)
}

// Response is a free data retrieval call binding the contract method 0x7a7f01a7.
//
// Solidity: function response() view returns(string)
func (_TestEXCALL *TestEXCALLCaller) Response(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TestEXCALL.contract.Call(opts, &out, "response")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Response is a free data retrieval call binding the contract method 0x7a7f01a7.
//
// Solidity: function response() view returns(string)
func (_TestEXCALL *TestEXCALLSession) Response() (string, error) {
	return _TestEXCALL.Contract.Response(&_TestEXCALL.CallOpts)
}

// Response is a free data retrieval call binding the contract method 0x7a7f01a7.
//
// Solidity: function response() view returns(string)
func (_TestEXCALL *TestEXCALLCallerSession) Response() (string, error) {
	return _TestEXCALL.Contract.Response(&_TestEXCALL.CallOpts)
}

// Winnings is a free data retrieval call binding the contract method 0xea3a1499.
//
// Solidity: function winnings(address ) view returns(int256)
func (_TestEXCALL *TestEXCALLCaller) Winnings(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TestEXCALL.contract.Call(opts, &out, "winnings", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Winnings is a free data retrieval call binding the contract method 0xea3a1499.
//
// Solidity: function winnings(address ) view returns(int256)
func (_TestEXCALL *TestEXCALLSession) Winnings(arg0 common.Address) (*big.Int, error) {
	return _TestEXCALL.Contract.Winnings(&_TestEXCALL.CallOpts, arg0)
}

// Winnings is a free data retrieval call binding the contract method 0xea3a1499.
//
// Solidity: function winnings(address ) view returns(int256)
func (_TestEXCALL *TestEXCALLCallerSession) Winnings(arg0 common.Address) (*big.Int, error) {
	return _TestEXCALL.Contract.Winnings(&_TestEXCALL.CallOpts, arg0)
}

// AlwaysWinGithub is a paid mutator transaction binding the contract method 0x78e4d58c.
//
// Solidity: function alwaysWinGithub() returns()
func (_TestEXCALL *TestEXCALLTransactor) AlwaysWinGithub(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestEXCALL.contract.Transact(opts, "alwaysWinGithub")
}

// AlwaysWinGithub is a paid mutator transaction binding the contract method 0x78e4d58c.
//
// Solidity: function alwaysWinGithub() returns()
func (_TestEXCALL *TestEXCALLSession) AlwaysWinGithub() (*types.Transaction, error) {
	return _TestEXCALL.Contract.AlwaysWinGithub(&_TestEXCALL.TransactOpts)
}

// AlwaysWinGithub is a paid mutator transaction binding the contract method 0x78e4d58c.
//
// Solidity: function alwaysWinGithub() returns()
func (_TestEXCALL *TestEXCALLTransactorSession) AlwaysWinGithub() (*types.Transaction, error) {
	return _TestEXCALL.Contract.AlwaysWinGithub(&_TestEXCALL.TransactOpts)
}

// AlwaysWinLocal is a paid mutator transaction binding the contract method 0xd7f8db5c.
//
// Solidity: function alwaysWinLocal() returns()
func (_TestEXCALL *TestEXCALLTransactor) AlwaysWinLocal(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestEXCALL.contract.Transact(opts, "alwaysWinLocal")
}

// AlwaysWinLocal is a paid mutator transaction binding the contract method 0xd7f8db5c.
//
// Solidity: function alwaysWinLocal() returns()
func (_TestEXCALL *TestEXCALLSession) AlwaysWinLocal() (*types.Transaction, error) {
	return _TestEXCALL.Contract.AlwaysWinLocal(&_TestEXCALL.TransactOpts)
}

// AlwaysWinLocal is a paid mutator transaction binding the contract method 0xd7f8db5c.
//
// Solidity: function alwaysWinLocal() returns()
func (_TestEXCALL *TestEXCALLTransactorSession) AlwaysWinLocal() (*types.Transaction, error) {
	return _TestEXCALL.Contract.AlwaysWinLocal(&_TestEXCALL.TransactOpts)
}

// BeginAlwaysWinOracle is a paid mutator transaction binding the contract method 0x7163c7d7.
//
// Solidity: function beginAlwaysWinOracle() returns()
func (_TestEXCALL *TestEXCALLTransactor) BeginAlwaysWinOracle(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestEXCALL.contract.Transact(opts, "beginAlwaysWinOracle")
}

// BeginAlwaysWinOracle is a paid mutator transaction binding the contract method 0x7163c7d7.
//
// Solidity: function beginAlwaysWinOracle() returns()
func (_TestEXCALL *TestEXCALLSession) BeginAlwaysWinOracle() (*types.Transaction, error) {
	return _TestEXCALL.Contract.BeginAlwaysWinOracle(&_TestEXCALL.TransactOpts)
}

// BeginAlwaysWinOracle is a paid mutator transaction binding the contract method 0x7163c7d7.
//
// Solidity: function beginAlwaysWinOracle() returns()
func (_TestEXCALL *TestEXCALLTransactorSession) BeginAlwaysWinOracle() (*types.Transaction, error) {
	return _TestEXCALL.Contract.BeginAlwaysWinOracle(&_TestEXCALL.TransactOpts)
}

// ContinueAlwaysWinOracle is a paid mutator transaction binding the contract method 0xd7c0d248.
//
// Solidity: function continueAlwaysWinOracle(address forPunter, uint256 oracleRef, bool answer) returns(bool)
func (_TestEXCALL *TestEXCALLTransactor) ContinueAlwaysWinOracle(opts *bind.TransactOpts, forPunter common.Address, oracleRef *big.Int, answer bool) (*types.Transaction, error) {
	return _TestEXCALL.contract.Transact(opts, "continueAlwaysWinOracle", forPunter, oracleRef, answer)
}

// ContinueAlwaysWinOracle is a paid mutator transaction binding the contract method 0xd7c0d248.
//
// Solidity: function continueAlwaysWinOracle(address forPunter, uint256 oracleRef, bool answer) returns(bool)
func (_TestEXCALL *TestEXCALLSession) ContinueAlwaysWinOracle(forPunter common.Address, oracleRef *big.Int, answer bool) (*types.Transaction, error) {
	return _TestEXCALL.Contract.ContinueAlwaysWinOracle(&_TestEXCALL.TransactOpts, forPunter, oracleRef, answer)
}

// ContinueAlwaysWinOracle is a paid mutator transaction binding the contract method 0xd7c0d248.
//
// Solidity: function continueAlwaysWinOracle(address forPunter, uint256 oracleRef, bool answer) returns(bool)
func (_TestEXCALL *TestEXCALLTransactorSession) ContinueAlwaysWinOracle(forPunter common.Address, oracleRef *big.Int, answer bool) (*types.Transaction, error) {
	return _TestEXCALL.Contract.ContinueAlwaysWinOracle(&_TestEXCALL.TransactOpts, forPunter, oracleRef, answer)
}

// Get0s is a paid mutator transaction binding the contract method 0x40d68738.
//
// Solidity: function get0s() returns(string)
func (_TestEXCALL *TestEXCALLTransactor) Get0s(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestEXCALL.contract.Transact(opts, "get0s")
}

// Get0s is a paid mutator transaction binding the contract method 0x40d68738.
//
// Solidity: function get0s() returns(string)
func (_TestEXCALL *TestEXCALLSession) Get0s() (*types.Transaction, error) {
	return _TestEXCALL.Contract.Get0s(&_TestEXCALL.TransactOpts)
}

// Get0s is a paid mutator transaction binding the contract method 0x40d68738.
//
// Solidity: function get0s() returns(string)
func (_TestEXCALL *TestEXCALLTransactorSession) Get0s() (*types.Transaction, error) {
	return _TestEXCALL.Contract.Get0s(&_TestEXCALL.TransactOpts)
}

// Get1s is a paid mutator transaction binding the contract method 0xcc832f28.
//
// Solidity: function get1s() returns(string)
func (_TestEXCALL *TestEXCALLTransactor) Get1s(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestEXCALL.contract.Transact(opts, "get1s")
}

// Get1s is a paid mutator transaction binding the contract method 0xcc832f28.
//
// Solidity: function get1s() returns(string)
func (_TestEXCALL *TestEXCALLSession) Get1s() (*types.Transaction, error) {
	return _TestEXCALL.Contract.Get1s(&_TestEXCALL.TransactOpts)
}

// Get1s is a paid mutator transaction binding the contract method 0xcc832f28.
//
// Solidity: function get1s() returns(string)
func (_TestEXCALL *TestEXCALLTransactorSession) Get1s() (*types.Transaction, error) {
	return _TestEXCALL.Contract.Get1s(&_TestEXCALL.TransactOpts)
}

// GetAAA is a paid mutator transaction binding the contract method 0x2d3085f6.
//
// Solidity: function getAAA() returns(string)
func (_TestEXCALL *TestEXCALLTransactor) GetAAA(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestEXCALL.contract.Transact(opts, "getAAA")
}

// GetAAA is a paid mutator transaction binding the contract method 0x2d3085f6.
//
// Solidity: function getAAA() returns(string)
func (_TestEXCALL *TestEXCALLSession) GetAAA() (*types.Transaction, error) {
	return _TestEXCALL.Contract.GetAAA(&_TestEXCALL.TransactOpts)
}

// GetAAA is a paid mutator transaction binding the contract method 0x2d3085f6.
//
// Solidity: function getAAA() returns(string)
func (_TestEXCALL *TestEXCALLTransactorSession) GetAAA() (*types.Transaction, error) {
	return _TestEXCALL.Contract.GetAAA(&_TestEXCALL.TransactOpts)
}

// GetECT is a paid mutator transaction binding the contract method 0xa59bffe0.
//
// Solidity: function getECT() returns(string)
func (_TestEXCALL *TestEXCALLTransactor) GetECT(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestEXCALL.contract.Transact(opts, "getECT")
}

// GetECT is a paid mutator transaction binding the contract method 0xa59bffe0.
//
// Solidity: function getECT() returns(string)
func (_TestEXCALL *TestEXCALLSession) GetECT() (*types.Transaction, error) {
	return _TestEXCALL.Contract.GetECT(&_TestEXCALL.TransactOpts)
}

// GetECT is a paid mutator transaction binding the contract method 0xa59bffe0.
//
// Solidity: function getECT() returns(string)
func (_TestEXCALL *TestEXCALLTransactorSession) GetECT() (*types.Transaction, error) {
	return _TestEXCALL.Contract.GetECT(&_TestEXCALL.TransactOpts)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address add) returns()
func (_TestEXCALL *TestEXCALLTransactor) SetOracle(opts *bind.TransactOpts, add common.Address) (*types.Transaction, error) {
	return _TestEXCALL.contract.Transact(opts, "setOracle", add)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address add) returns()
func (_TestEXCALL *TestEXCALLSession) SetOracle(add common.Address) (*types.Transaction, error) {
	return _TestEXCALL.Contract.SetOracle(&_TestEXCALL.TransactOpts, add)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address add) returns()
func (_TestEXCALL *TestEXCALLTransactorSession) SetOracle(add common.Address) (*types.Transaction, error) {
	return _TestEXCALL.Contract.SetOracle(&_TestEXCALL.TransactOpts, add)
}

// SometimesWin is a paid mutator transaction binding the contract method 0x09554752.
//
// Solidity: function sometimesWin() returns()
func (_TestEXCALL *TestEXCALLTransactor) SometimesWin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestEXCALL.contract.Transact(opts, "sometimesWin")
}

// SometimesWin is a paid mutator transaction binding the contract method 0x09554752.
//
// Solidity: function sometimesWin() returns()
func (_TestEXCALL *TestEXCALLSession) SometimesWin() (*types.Transaction, error) {
	return _TestEXCALL.Contract.SometimesWin(&_TestEXCALL.TransactOpts)
}

// SometimesWin is a paid mutator transaction binding the contract method 0x09554752.
//
// Solidity: function sometimesWin() returns()
func (_TestEXCALL *TestEXCALLTransactorSession) SometimesWin() (*types.Transaction, error) {
	return _TestEXCALL.Contract.SometimesWin(&_TestEXCALL.TransactOpts)
}

// TestEXCALLBetPlacedIterator is returned from FilterBetPlaced and is used to iterate over the raw logs and unpacked data for BetPlaced events raised by the TestEXCALL contract.
type TestEXCALLBetPlacedIterator struct {
	Event *TestEXCALLBetPlaced // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TestEXCALLBetPlacedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestEXCALLBetPlaced)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TestEXCALLBetPlaced)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TestEXCALLBetPlacedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestEXCALLBetPlacedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestEXCALLBetPlaced represents a BetPlaced event raised by the TestEXCALL contract.
type TestEXCALLBetPlaced struct {
	From        common.Address
	OracleRefNr *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBetPlaced is a free log retrieval operation binding the contract event 0x43e08d78302cdf9c94e1ffd293c2a3696139ee41cf3199b6f5eed9c7f6cc6090.
//
// Solidity: event BetPlaced(address indexed _from, uint256 oracleRefNr)
func (_TestEXCALL *TestEXCALLFilterer) FilterBetPlaced(opts *bind.FilterOpts, _from []common.Address) (*TestEXCALLBetPlacedIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _TestEXCALL.contract.FilterLogs(opts, "BetPlaced", _fromRule)
	if err != nil {
		return nil, err
	}
	return &TestEXCALLBetPlacedIterator{contract: _TestEXCALL.contract, event: "BetPlaced", logs: logs, sub: sub}, nil
}

// WatchBetPlaced is a free log subscription operation binding the contract event 0x43e08d78302cdf9c94e1ffd293c2a3696139ee41cf3199b6f5eed9c7f6cc6090.
//
// Solidity: event BetPlaced(address indexed _from, uint256 oracleRefNr)
func (_TestEXCALL *TestEXCALLFilterer) WatchBetPlaced(opts *bind.WatchOpts, sink chan<- *TestEXCALLBetPlaced, _from []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _TestEXCALL.contract.WatchLogs(opts, "BetPlaced", _fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestEXCALLBetPlaced)
				if err := _TestEXCALL.contract.UnpackLog(event, "BetPlaced", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBetPlaced is a log parse operation binding the contract event 0x43e08d78302cdf9c94e1ffd293c2a3696139ee41cf3199b6f5eed9c7f6cc6090.
//
// Solidity: event BetPlaced(address indexed _from, uint256 oracleRefNr)
func (_TestEXCALL *TestEXCALLFilterer) ParseBetPlaced(log types.Log) (*TestEXCALLBetPlaced, error) {
	event := new(TestEXCALLBetPlaced)
	if err := _TestEXCALL.contract.UnpackLog(event, "BetPlaced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
