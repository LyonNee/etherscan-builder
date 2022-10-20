package builder

import (
	"strconv"
	"strings"

	"github.com/LyonNee/etherscan-builder/builder"
	"github.com/LyonNee/etherscan-builder/enums/action"
	"github.com/LyonNee/etherscan-builder/enums/sort"
)

type AccountBuilder struct {
	urlbuf *strings.Builder
}

func (builder *AccountBuilder) Build(apiKey string) string {
	builder.urlbuf.WriteString("&apikey=")
	builder.urlbuf.WriteString(apiKey)

	return builder.urlbuf.String()
}

func (builder *AccountBuilder) Action(action action.AccountEnum) builder.BuilderInf {
	builder.urlbuf.WriteString("&acction=")
	builder.urlbuf.WriteString(string(action))

	return builder
}

func (builder *AccountBuilder) Address(address string) builder.BuilderInf {
	builder.urlbuf.WriteString("&address=")
	builder.urlbuf.WriteString(address)

	return builder
}

func (builder *AccountBuilder) ContractAddress(address string) builder.BuilderInf {
	builder.urlbuf.WriteString("&contractaddress=")
	builder.urlbuf.WriteString(address)

	return builder
}

func (builder *AccountBuilder) Tag(tag string) builder.BuilderInf {
	builder.urlbuf.WriteString("&tag=")
	builder.urlbuf.WriteString(tag)

	return builder
}

func (builder *AccountBuilder) Startblock(startblock uint64) builder.BuilderInf {
	builder.urlbuf.WriteString("&startblock=")
	builder.urlbuf.WriteString(strconv.FormatUint(startblock, 64))

	return builder
}

func (builder *AccountBuilder) Endblock(endblock uint64) builder.BuilderInf {
	builder.urlbuf.WriteString("&endblock=")
	builder.urlbuf.WriteString(strconv.FormatUint(endblock, 64))

	return builder
}

func (builder *AccountBuilder) Page(page uint64) builder.BuilderInf {
	builder.urlbuf.WriteString("&page=")
	builder.urlbuf.WriteString(strconv.FormatUint(page, 64))

	return builder
}

func (builder *AccountBuilder) Offset(offset uint64) builder.BuilderInf {
	builder.urlbuf.WriteString("&offset=")
	builder.urlbuf.WriteString(strconv.FormatUint(offset, 64))

	return builder
}

//the sorting preference, use asc to sort by ascending and desc to sort by descending
func (builder *AccountBuilder) Sort(sort sort.SortEnum) builder.BuilderInf {
	builder.urlbuf.WriteString("&sort=")
	builder.urlbuf.WriteString(string(sort))

	return builder
}
