package builder

import (
	"strconv"
	"strings"

	"github.com/LyonNee/etherscan-builder/enums/action"
	"github.com/LyonNee/etherscan-builder/enums/closest"
	"github.com/LyonNee/etherscan-builder/enums/module"
	"github.com/LyonNee/etherscan-builder/enums/network"
	"github.com/LyonNee/etherscan-builder/enums/sort"
	"github.com/LyonNee/etherscan-builder/enums/topicoperator"
)

type Builder struct {
	urlbuf *strings.Builder
}

func New(network network.NetworkEnum, module module.ModuleEnum) *Builder {
	urlbuf := strings.Builder{}
	urlbuf.WriteString(string(network))

	urlbuf.WriteString("api")
	urlbuf.WriteString("?module=")
	urlbuf.WriteString(string(module))

	return &Builder{&urlbuf}
}

func (builder *Builder) Build(apiKey string) string {
	builder.urlbuf.WriteString("&apikey=")
	builder.urlbuf.WriteString(apiKey)

	return builder.urlbuf.String()
}

func (builder *Builder) Action(action action.AccountEnum) *Builder {
	builder.urlbuf.WriteString("&acction=")
	builder.urlbuf.WriteString(string(action))

	return builder
}

// account
func (builder *Builder) Address(address string) *Builder {
	builder.urlbuf.WriteString("&address=")
	builder.urlbuf.WriteString(address)

	return builder
}

// account
func (builder *Builder) ContractAddress(address string) *Builder {
	builder.urlbuf.WriteString("&contractaddress=")
	builder.urlbuf.WriteString(address)

	return builder
}

// account
func (builder *Builder) Tag(tag string) *Builder {
	builder.urlbuf.WriteString("&tag=")
	builder.urlbuf.WriteString(tag)

	return builder
}

// account
func (builder *Builder) Startblock(startblock uint64) *Builder {
	builder.urlbuf.WriteString("&startblock=")
	builder.urlbuf.WriteString(strconv.FormatUint(startblock, 64))

	return builder
}

// account
func (builder *Builder) Endblock(endblock uint64) *Builder {
	builder.urlbuf.WriteString("&endblock=")
	builder.urlbuf.WriteString(strconv.FormatUint(endblock, 64))

	return builder
}

// account
func (builder *Builder) Page(page uint64) *Builder {
	builder.urlbuf.WriteString("&page=")
	builder.urlbuf.WriteString(strconv.FormatUint(page, 64))

	return builder
}

// account
func (builder *Builder) Offset(offset uint64) *Builder {
	builder.urlbuf.WriteString("&offset=")
	builder.urlbuf.WriteString(strconv.FormatUint(offset, 64))

	return builder
}

// account

//the sorting preference, use asc to sort by ascending and desc to sort by descending
func (builder *Builder) Sort(sort sort.SortEnum) *Builder {
	builder.urlbuf.WriteString("&sort=")
	builder.urlbuf.WriteString(string(sort))

	return builder
}

// Transactions
func (builder *Builder) Txhash(txhash string) *Builder {
	builder.urlbuf.WriteString("&txhash=")
	builder.urlbuf.WriteString(txhash)

	return builder
}

// Blocks
func (builder *Builder) Blockno(blockno uint64) *Builder {
	builder.urlbuf.WriteString("&blockno=")
	builder.urlbuf.WriteString(strconv.FormatUint(blockno, 64))

	return builder
}

func (builder *Builder) Timestamp(timestamp uint64) *Builder {
	builder.urlbuf.WriteString("&timestamp=")
	builder.urlbuf.WriteString(strconv.FormatUint(timestamp, 64))

	return builder
}

func (builder *Builder) Closest(closest closest.ClosestEnum) *Builder {
	builder.urlbuf.WriteString("&timestamp=")
	builder.urlbuf.WriteString(string(closest))

	return builder
}

// the starting date in yyyy-MM-dd format, eg. 2019-02-01
func (builder *Builder) Startdate(startdate string) *Builder {
	builder.urlbuf.WriteString("&startdate=")
	builder.urlbuf.WriteString(startdate)

	return builder
}

// the ending date in yyyy-MM-dd format, eg. 2019-02-28
func (builder *Builder) Enddate(enddate string) *Builder {
	builder.urlbuf.WriteString("&enddate=")
	builder.urlbuf.WriteString(enddate)

	return builder
}

// Logs

// the integer block number to start searching for logs eg. 12878196
func (builder *Builder) FromBlock(fromBlock uint64) *Builder {
	builder.urlbuf.WriteString("&fromBlock=")
	builder.urlbuf.WriteString(strconv.FormatUint(fromBlock, 64))

	return builder
}

// the integer block number to stop searching for logs eg. 12879196
func (builder *Builder) ToBlock(toBlock uint64) *Builder {
	builder.urlbuf.WriteString("&toBlock=")
	builder.urlbuf.WriteString(strconv.FormatUint(toBlock, 64))

	return builder
}

// the topic numbers to search for
// limited totopic0, topic1, topic2, topic3
func (builder *Builder) Topic0(topic0 string) *Builder {
	builder.urlbuf.WriteString("&topic0=")
	builder.urlbuf.WriteString(topic0)

	return builder
}

// the topic numbers to search for
// limited totopic0, topic1, topic2, topic3
func (builder *Builder) Topic1(topic1 string) *Builder {
	builder.urlbuf.WriteString("&topic1=")
	builder.urlbuf.WriteString(topic1)

	return builder
}

// the topic numbers to search for
// limited totopic0, topic1, topic2, topic3
func (builder *Builder) Topic2(topic2 string) *Builder {
	builder.urlbuf.WriteString("&topic2=")
	builder.urlbuf.WriteString(topic2)

	return builder
}

// the topic numbers to search for
// limited totopic0, topic1, topic2, topic3
func (builder *Builder) Topic3(topic3 string) *Builder {
	builder.urlbuf.WriteString("&topic3=")
	builder.urlbuf.WriteString(topic3)

	return builder
}

// the topic operator when multiple topic combinations are used
// limited to and or or
func (builder *Builder) TopicOperator(topicOperator topicoperator.TopicOperatorEnum) *Builder {
	builder.urlbuf.WriteString("&topicOperator=")
	builder.urlbuf.WriteString(string(topicOperator))

	return builder
}
