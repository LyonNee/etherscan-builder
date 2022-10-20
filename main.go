package etherscanbuilder

import (
	"fmt"

	"github.com/LyonNee/etherscan-builder/builder"
	"github.com/LyonNee/etherscan-builder/enums/module"
	"github.com/LyonNee/etherscan-builder/enums/network"
)

func main() {
	b := builder.New(network.MainNet, module.Account)
	b.Address("0x1a33b")

	fmt.Println(b.Build(""))
}
