package builder

import "github.com/LyonNee/etherscan-builder/enums/action"

type BuilderInf interface {
	Action(action action.AccountEnum) BuilderInf
	Build(apiKey string) string
}
