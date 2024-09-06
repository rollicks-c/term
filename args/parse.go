package args

import (
	"fmt"
	"github.com/rollicks-c/term/io"
	"os"
)

func withParsers[T any](parsers ...argParser[T]) ArgOption[T] {
	return func(ctx *ArgContext[T]) {
		ctx.parsers = append(ctx.parsers, parsers...)
	}
}

func parseString(exp string) (string, error) {
	return exp, nil
}

func itemSelector(provider ListProvider) argParser[any] {
	return func(exp string) (any, error) {

		// filter
		selection, err := provider.SearchItems(exp)
		if err != nil {
			return nil, err
		}
		if len(selection) == 0 {
			return nil, fmt.Errorf("no item found for expression [%s]", exp)
		}
		if len(selection) == 1 {
			return selection[0].Value, nil
		}

		// choose
		list := make(map[string]any)
		for _, item := range selection {
			list[item.Name] = item.Value
		}
		sel, err := io.New(os.Stdin, os.Stdout).Choose("choose exact item:", list)
		if err != nil {
			return nil, err
		}

		return sel, nil
	}
}
