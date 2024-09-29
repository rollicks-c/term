package term

import (
	"github.com/rollicks-c/term/io/table"
)

func WithHideHeaders(state bool) table.Option {
	return func(config *table.Config) {
		config.HideHeaders = state
	}
}

func WithIndention(chars string) table.Option {
	return func(config *table.Config) {
		config.Indention = chars
	}
}

func TableEx[T any](options ...table.Option) *table.Builder[T] {
	return table.NewBuilder[T](options...)
}
