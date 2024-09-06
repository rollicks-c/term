package term

import (
	"github.com/rollicks-c/term/args"
	"github.com/rollicks-c/term/io"
	"os"
)

func NewArgsCollector(argList []string, options ...args.CollectorOption) *args.Collector {
	return args.NewCollector(argList, options...)
}

func IO() *io.Module {
	return io.New(os.Stdin, os.Stdout)
}
