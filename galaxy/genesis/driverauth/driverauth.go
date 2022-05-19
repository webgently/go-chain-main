package driverauth

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// GetContractBin is NodeDriverAuth contract genesis implementation bin code
// Has to be compiled with flag bin-runtime
// Built from galaxy-sfc c1d33c81f74abf82c0e22807f16e609578e10ad8, solc 0.5.17+commit.d19bba13.Emscripten.clang, optimize-runs 10000
func GetContractBin() []byte {
	return hexutil.MustDecode("0x608060405234801561001057600080fd5b506004361061016c5760003560e01c80638da5cb5b116100cd578063d6a0c7af11610081578063ebdf104c11610066578063ebdf104c1461056c578063f2fde38b146106d2578063fd1b6ec1146106f85761016c565b8063d6a0c7af146104ce578063e08d7e66146104fc5761016c565b8063a4066fbe116100b2578063a4066fbe14610403578063b9cc6b1c14610426578063c0c53b8b146104965761016c565b80638da5cb5b146103c35780638f32d59b146103e75761016c565b80634ddaf8f21161012457806366e7ea0f1161010957806366e7ea0f14610363578063715018a61461038f57806379bead38146103975761016c565b80634ddaf8f21461029f5780634feb92f3146102c55761016c565b80631e702f83116101555780631e702f83146101e8578063242a6e3f1461020b578063267ab446146102825761016c565b80630aeeca001461017157806318f628d414610190575b600080fd5b61018e6004803603602081101561018757600080fd5b5035610726565b005b61018e60048036036101208110156101a757600080fd5b506001600160a01b038135169060208101359060408101359060608101359060808101359060a08101359060c08101359060e0810135906101000135610800565b61018e600480360360408110156101fe57600080fd5b508035906020013561090d565b61018e6004803603604081101561022157600080fd5b8135919081019060408101602082013564010000000081111561024357600080fd5b82018360208201111561025557600080fd5b8035906020019184600183028401116401000000008311171561027757600080fd5b5090925090506109df565b61018e6004803603602081101561029857600080fd5b5035610af4565b61018e600480360360208110156102b557600080fd5b50356001600160a01b0316610bb3565b61018e60048036036101008110156102dc57600080fd5b6001600160a01b038235169160208101359181019060608101604082013564010000000081111561030c57600080fd5b82018360208201111561031e57600080fd5b8035906020019184600183028401116401000000008311171561034057600080fd5b919350915080359060208101359060408101359060608101359060800135610c73565b61018e6004803603604081101561037957600080fd5b506001600160a01b038135169060200135610d8d565b61018e610eb1565b61018e600480360360408110156103ad57600080fd5b506001600160a01b038135169060200135610f6c565b6103cb611033565b604080516001600160a01b039092168252519081900360200190f35b6103ef611042565b604080519115158252519081900360200190f35b61018e6004803603604081101561041957600080fd5b5080359060200135611053565b61018e6004803603602081101561043c57600080fd5b81019060208101813564010000000081111561045757600080fd5b82018360208201111561046957600080fd5b8035906020019184600183028401116401000000008311171561048b57600080fd5b50909250905061111f565b61018e600480360360608110156104ac57600080fd5b506001600160a01b038135811691602081013582169160409091013516611208565b61018e600480360360408110156104e457600080fd5b506001600160a01b0381358116916020013516611356565b61018e6004803603602081101561051257600080fd5b81019060208101813564010000000081111561052d57600080fd5b82018360208201111561053f57600080fd5b8035906020019184602083028401116401000000008311171561056157600080fd5b50909250905061141e565b61018e6004803603608081101561058257600080fd5b81019060208101813564010000000081111561059d57600080fd5b8201836020820111156105af57600080fd5b803590602001918460208302840111640100000000831117156105d157600080fd5b9193909290916020810190356401000000008111156105ef57600080fd5b82018360208201111561060157600080fd5b8035906020019184602083028401116401000000008311171561062357600080fd5b91939092909160208101903564010000000081111561064157600080fd5b82018360208201111561065357600080fd5b8035906020019184602083028401116401000000008311171561067557600080fd5b91939092909160208101903564010000000081111561069357600080fd5b8201836020820111156106a557600080fd5b803590602001918460208302840111640100000000831117156106c757600080fd5b5090925090506114fd565b61018e600480360360208110156106e857600080fd5b50356001600160a01b03166116dc565b61018e6004803603604081101561070e57600080fd5b506001600160a01b0381358116916020013516611741565b61072e611042565b61077f576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b606754604080517f0aeeca000000000000000000000000000000000000000000000000000000000081526004810184905290516001600160a01b0390921691630aeeca009160248082019260009290919082900301818387803b1580156107e557600080fd5b505af11580156107f9573d6000803e3d6000fd5b5050505050565b6067546001600160a01b031633146108495760405162461bcd60e51b8152600401808060200182810382526025815260200180611b026025913960400191505060405180910390fd5b606654604080517f18f628d40000000000000000000000000000000000000000000000000000000081526001600160a01b038c81166004830152602482018c9052604482018b9052606482018a90526084820189905260a4820188905260c4820187905260e482018690526101048201859052915191909216916318f628d49161012480830192600092919082900301818387803b1580156108ea57600080fd5b505af11580156108fe573d6000803e3d6000fd5b50505050505050505050505050565b6067546001600160a01b031633146109565760405162461bcd60e51b8152600401808060200182810382526025815260200180611b026025913960400191505060405180910390fd5b606654604080517f1e702f83000000000000000000000000000000000000000000000000000000008152600481018590526024810184905290516001600160a01b0390921691631e702f839160448082019260009290919082900301818387803b1580156109c357600080fd5b505af11580156109d7573d6000803e3d6000fd5b505050505050565b6066546001600160a01b03163314610a3e576040805162461bcd60e51b815260206004820152601e60248201527f63616c6c6572206973206e6f74207468652053464320636f6e74726163740000604482015290519081900360640190fd5b606754604080517f242a6e3f0000000000000000000000000000000000000000000000000000000081526004810186815260248201928352604482018590526001600160a01b039093169263242a6e3f928792879287929091606401848480828437600081840152601f19601f820116905080830192505050945050505050600060405180830381600087803b158015610ad757600080fd5b505af1158015610aeb573d6000803e3d6000fd5b50505050505050565b610afc611042565b610b4d576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b606754604080517f267ab4460000000000000000000000000000000000000000000000000000000081526004810184905290516001600160a01b039092169163267ab4469160248082019260009290919082900301818387803b1580156107e557600080fd5b610bbb611042565b610c0c576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b606754604080517fda7fc24f0000000000000000000000000000000000000000000000000000000081526001600160a01b0384811660048301529151919092169163da7fc24f91602480830192600092919082900301818387803b1580156107e557600080fd5b6067546001600160a01b03163314610cbc5760405162461bcd60e51b8152600401808060200182810382526025815260200180611b026025913960400191505060405180910390fd5b606660009054906101000a90046001600160a01b03166001600160a01b0316634feb92f38a8a8a8a8a8a8a8a8a6040518a63ffffffff1660e01b8152600401808a6001600160a01b03166001600160a01b03168152602001898152602001806020018781526020018681526020018581526020018481526020018381526020018281038252898982818152602001925080828437600081840152601f19601f8201169050808301925050509a5050505050505050505050600060405180830381600087803b1580156108ea57600080fd5b6066546001600160a01b03163314610dec576040805162461bcd60e51b815260206004820152601e60248201527f63616c6c6572206973206e6f74207468652053464320636f6e74726163740000604482015290519081900360640190fd5b6066546001600160a01b03838116911614610e385760405162461bcd60e51b8152600401808060200182810382526021815260200180611ae16021913960400191505060405180910390fd5b6067546001600160a01b039081169063e30443bc908490610e62908216318563ffffffff61180416565b6040518363ffffffff1660e01b815260040180836001600160a01b03166001600160a01b0316815260200182815260200192505050600060405180830381600087803b1580156109c357600080fd5b610eb9611042565b610f0a576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b6033546040516000916001600160a01b0316907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3603380547fffffffffffffffffffffffff0000000000000000000000000000000000000000169055565b610f74611042565b610fc5576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b606754604080517f79bead380000000000000000000000000000000000000000000000000000000081526001600160a01b03858116600483015260248201859052915191909216916379bead3891604480830192600092919082900301818387803b1580156109c357600080fd5b6033546001600160a01b031690565b6033546001600160a01b0316331490565b6066546001600160a01b031633146110b2576040805162461bcd60e51b815260206004820152601e60248201527f63616c6c6572206973206e6f74207468652053464320636f6e74726163740000604482015290519081900360640190fd5b606754604080517fa4066fbe000000000000000000000000000000000000000000000000000000008152600481018590526024810184905290516001600160a01b039092169163a4066fbe9160448082019260009290919082900301818387803b1580156109c357600080fd5b611127611042565b611178576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b6067546040517fb9cc6b1c000000000000000000000000000000000000000000000000000000008152602060048201908152602482018490526001600160a01b039092169163b9cc6b1c91859185918190604401848480828437600081840152601f19601f8201169050808301925050509350505050600060405180830381600087803b1580156109c357600080fd5b600054610100900460ff16806112215750611221611865565b8061122f575060005460ff16155b61126a5760405162461bcd60e51b815260040180806020018281038252602e815260200180611ab3602e913960400191505060405180910390fd5b600054610100900460ff161580156112d057600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff909116610100171660011790555b6112d98261186b565b606780546001600160a01b038086167fffffffffffffffffffffffff0000000000000000000000000000000000000000928316179092556066805492871692909116919091179055801561135057600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff1690555b50505050565b61135e611042565b6113af576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b606754604080517fd6a0c7af0000000000000000000000000000000000000000000000000000000081526001600160a01b03858116600483015284811660248301529151919092169163d6a0c7af91604480830192600092919082900301818387803b1580156109c357600080fd5b6067546001600160a01b031633146114675760405162461bcd60e51b8152600401808060200182810382526025815260200180611b026025913960400191505060405180910390fd5b6066546040517fe08d7e66000000000000000000000000000000000000000000000000000000008152602060048201818152602483018590526001600160a01b039093169263e08d7e6692869286929182916044909101908590850280828437600081840152601f19601f8201169050808301925050509350505050600060405180830381600087803b1580156109c357600080fd5b6067546001600160a01b031633146115465760405162461bcd60e51b8152600401808060200182810382526025815260200180611b026025913960400191505060405180910390fd5b606660009054906101000a90046001600160a01b03166001600160a01b031663ebdf104c89898989898989896040518963ffffffff1660e01b8152600401808060200180602001806020018060200185810385528d8d82818152602001925060200280828437600083820152601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01690910186810385528b8152602090810191508c908c0280828437600083820152601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169091018681038452898152602090810191508a908a0280828437600083820152601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169091018681038352878152602090810191508890880280828437600081840152601f19601f8201169050808301925050509c50505050505050505050505050600060405180830381600087803b1580156116ba57600080fd5b505af11580156116ce573d6000803e3d6000fd5b505050505050505050505050565b6116e4611042565b611735576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b61173e816119cd565b50565b611749611042565b61179a576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b6117a382611a86565b80156117b357506117b381611a86565b6113af576040805162461bcd60e51b815260206004820152600e60248201527f6e6f74206120636f6e7472616374000000000000000000000000000000000000604482015290519081900360640190fd5b60008282018381101561185e576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b9392505050565b303b1590565b600054610100900460ff16806118845750611884611865565b80611892575060005460ff16155b6118cd5760405162461bcd60e51b815260040180806020018281038252602e815260200180611ab3602e913960400191505060405180910390fd5b600054610100900460ff1615801561193357600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff909116610100171660011790555b603380547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0384811691909117918290556040519116906000907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a380156119c957600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff1690555b5050565b6001600160a01b038116611a125760405162461bcd60e51b8152600401808060200182810382526026815260200180611a8d6026913960400191505060405180910390fd5b6033546040516001600160a01b038084169216907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a3603380547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b3b15159056fe4f776e61626c653a206e6577206f776e657220697320746865207a65726f2061646472657373436f6e747261637420696e7374616e63652068617320616c7265616479206265656e20696e697469616c697a6564726563697069656e74206973206e6f74207468652053464320636f6e747261637463616c6c6572206973206e6f7420746865204e6f646544726976657220636f6e7472616374a265627a7a72315820ed0d5f52fab9d38e7f0e18c7cc3ea5557f637a4770dc59ed65330f9c333c9b5d64736f6c63430005110032")
}

// ContractAddress is the NodeDriverAuth contract address
var ContractAddress = common.HexToAddress("0x1c1cb000000000000000000000000000000000ac")
