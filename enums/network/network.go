package network

type NetworkEnum string

const (
	MainNet NetworkEnum = "https://api.etherscan.io/"
	Goerli  NetworkEnum = "https://api-goerli.etherscan.io/"
	Kovan   NetworkEnum = "https://api-kovan.etherscan.io/"
	Rinkeby NetworkEnum = "https://api-rinkeby.etherscan.io/"
	Ropsten NetworkEnum = "https://api-ropsten.etherscan.io/"
	Sepolia NetworkEnum = "https://api-sepolia.etherscan.io/"
)
