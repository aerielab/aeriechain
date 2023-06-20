package contracts

import "github.com/0xPolygon/polygon-edge/types"

var (
	// ValidatorSetContract is an address of validator set contract
	ValidatorSetContract = types.StringToAddress("0x101")
	// BLSContract is an address of BLS contract
	BLSContract = types.StringToAddress("0x102")
	// MerkleContract is an address of Merkle contract
	MerkleContract = types.StringToAddress("0x103")
	// RewardTokenContract is an address of reward token on child chain
	RewardTokenContract = types.StringToAddress("0x104")
	// RewardPoolContract is an address of RewardPoolContract contract
	RewardPoolContract = types.StringToAddress("0x105")
	// StateReceiverContract is an address of bridge contract
	StateReceiverContract = types.StringToAddress("0x1001")
	// NativeERC20TokenContract is an address of the native token
	NativeERC20TokenContract = types.StringToAddress("0x1010")
	// L2StateSenderContract is an address of arbitrary message bridge sender
	L2StateSenderContract = types.StringToAddress("0x1002")

	// ChildERC20Contract is an address of bridgable ERC20 token contract on the child chain
	ChildERC20Contract = types.StringToAddress("0x1003")
	// ChildERC20PredicateContract is an address of child ERC20 predicate contract
	ChildERC20PredicateContract = types.StringToAddress("0x1004")
	// ChildERC721Contract is an address of bridgable ERC721 token contract on the child chain
	ChildERC721Contract = types.StringToAddress("0x1005")
	// ChildERC721PredicateContract is an address of child ERC721 predicate contract
	ChildERC721PredicateContract = types.StringToAddress("0x1006")
	// ChildERC1155Contract is an address of bridgable ERC1155 token contract on the child chain
	ChildERC1155Contract = types.StringToAddress("0x1005")
	// ChildERC1155PredicateContract is an address of child ERC1155 predicate contract
	ChildERC1155PredicateContract = types.StringToAddress("0x1008")
	// RootMintableERC20PredicateContract is an address of mintable ERC20 predicate
	RootMintableERC20PredicateContract = types.StringToAddress("0x1009")
	// RootMintableERC721PredicateContract is an address of mintable ERC721 predicate
	RootMintableERC721PredicateContract = types.StringToAddress("0x100a")
	// RootMintableERC1155PredicateContract is an address of mintable ERC1155 predicate
	RootMintableERC1155PredicateContract = types.StringToAddress("0x100b")

	// SystemCaller is address of account, used for system calls to smart contracts
	SystemCaller = types.StringToAddress("0xffffFFFfFFffffffffffffffFfFFFfffFFFfFFfE")

	// NativeTransferPrecompile is an address of native transfer precompile
	NativeTransferPrecompile = types.StringToAddress("0x2020")
	// BLSAggSigsVerificationPrecompile is an address of BLS aggregated signatures verificatin precompile
	BLSAggSigsVerificationPrecompile = types.StringToAddress("0x2030")
	// ConsolePrecompile is and address of Hardhat console precompile
	ConsolePrecompile = types.StringToAddress("0x000000000000000000636F6e736F6c652e6c6f67")
	// AllowListContractsAddr is the address of the contract deployer allow list
	AllowListContractsAddr = types.StringToAddress("0x0200000000000000000000000000000000000000")
	// BlockListContractsAddr is the address of the contract deployer block list
	BlockListContractsAddr = types.StringToAddress("0x0300000000000000000000000000000000000000")
	// AllowListTransactionsAddr is the address of the transactions allow list
	AllowListTransactionsAddr = types.StringToAddress("0x0200000000000000000000000000000000000002")
	// BlockListTransactionsAddr is the address of the transactions block list
	BlockListTransactionsAddr = types.StringToAddress("0x0300000000000000000000000000000000000002")
	// AllowListBridgeAddr is the address of the bridge allow list
	AllowListBridgeAddr = types.StringToAddress("0x0200000000000000000000000000000000000004")
	// BlockListBridgeAddr is the address of the bridge block list
	BlockListBridgeAddr = types.StringToAddress("0x0300000000000000000000000000000000000004")
)
