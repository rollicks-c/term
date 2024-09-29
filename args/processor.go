package args

import "fmt"

type argParser[T any] func(string) (T, error)

func retrieve[T any](args []string, index int, options ...ArgOption[T]) (T, error) {

	// init context
	ctx := &ArgContext[T]{}
	for _, opt := range options {
		opt(ctx)
	}
	if len(ctx.parsers) == 0 {
		return ctx.value, fmt.Errorf("no parsers provided")
	}

	// use default if arg not provided
	if index >= len(args) {
		if ctx.defaultValue == nil {
			return ctx.value, fmt.Errorf("missing arg at index %d", index)
		}
		return *ctx.defaultValue, nil
	}

	// try parsers
	var parsedValue T
	var parseErr error
	for _, p := range ctx.parsers {
		value, err := p(args[index])
		if err != nil {
			parseErr = err
			continue
		}
		parsedValue = value
		parseErr = nil
		break
	}

	return parsedValue, parseErr
}
