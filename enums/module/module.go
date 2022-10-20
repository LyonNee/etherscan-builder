package module

type ModuleEnum string

const (
	Account     ModuleEnum = "account"
	Contract    ModuleEnum = "contract"
	Transaction ModuleEnum = "transaction"
	Block       ModuleEnum = "block"
	Logs        ModuleEnum = "logs"
	Proxy       ModuleEnum = "proxy"
	Stats       ModuleEnum = "stats"
	Gastracker  ModuleEnum = "gastracker"
)
