package action

type AccountEnum string

const (
	// account
	Balance        AccountEnum = "balance"
	Balancemulti   AccountEnum = "balancemulti"
	Txlist         AccountEnum = "txlist"
	Txlistinternal AccountEnum = "txlistinternal"
	Tokentx        AccountEnum = "tokentx"
	Tokennfttx     AccountEnum = "tokennfttx"
	Token1155tx    AccountEnum = "token1155tx"
	Getminedblocks AccountEnum = "getminedblocks"
	Balancehistory AccountEnum = "balancehistory"

	// contract
	Getabi              AccountEnum = "getabi"
	Getsourcecode       AccountEnum = "getsourcecode"
	Getcontractcreation AccountEnum = "getcontractcreation"

	// transactions
	Getstatus          AccountEnum = "getstatus"
	Gettxreceiptstatus AccountEnum = "gettxreceiptstatus"

	// blocks
	Getblockreward     AccountEnum = "getblockreward"
	Getblockcountdown  AccountEnum = "getblockcountdown"
	Getblocknobytime   AccountEnum = "getblocknobytime"
	Dailyavgblocksize  AccountEnum = "dailyavgblocksize"
	Dailyblkcount      AccountEnum = "dailyblkcount"
	Dailyblockrewards  AccountEnum = "dailyblockrewards"
	Dailyavgblocktime  AccountEnum = "dailyavgblocktime"
	Dailyuncleblkcount AccountEnum = "dailyuncleblkcount"

	// logs
	GetLogs AccountEnum = "getLogs"

	// proxy
	EthblockNumber                         AccountEnum = "eth_blockNumber"
	EthgetBlockByNumber                    AccountEnum = "eth_getBlockByNumber"
	EthgetUncleByBlockNumberAndIndex       AccountEnum = "eth_getUncleByBlockNumberAndIndex"
	EthgetBlockTransactionCountByNumber    AccountEnum = "eth_getBlockTransactionCountByNumber"
	EthgetTransactionByHash                AccountEnum = "eth_getTransactionByHash"
	EthgetTransactionByBlockNumberAndIndex AccountEnum = "eth_getTransactionByBlockNumberAndIndex"
	EthgetTransactionCount                 AccountEnum = "eth_getTransactionCount"
	EthsendRawTransaction                  AccountEnum = "eth_sendRawTransaction"
	EthgetTransactionReceipt               AccountEnum = "eth_getTransactionReceipt"
	Ethcall                                AccountEnum = "eth_call"
	EthgetCode                             AccountEnum = "eth_getCode"
	EthgetStorageAt                        AccountEnum = "eth_getStorageAt"
	EthgasPrice                            AccountEnum = "eth_gasPrice"
	EthestimateGas                         AccountEnum = "eth_estimateGas"

	// token
	Tokensupply              AccountEnum = "tokensupply"
	Tokenbalance             AccountEnum = "tokenbalance"
	Tokensupplyhistory       AccountEnum = "tokensupplyhistory"
	Tokenholderlist          AccountEnum = "tokenholderlist"
	Tokeninfo                AccountEnum = "tokeninfo"
	Addresstokenbalance      AccountEnum = "addresstokenbalance"
	Addresstokennftbalance   AccountEnum = "addresstokennftbalance"
	Addresstokennftinventory AccountEnum = "addresstokennftinventory"

	// gastracker
	Gasestimate      AccountEnum = "gasestimate"
	Gasoracle        AccountEnum = "gasoracle"
	Dailyavggaslimit AccountEnum = "dailyavggaslimit"
	Dailygasused     AccountEnum = "dailygasused"
	Dailyavggasprice AccountEnum = "dailyavggasprice"
)
