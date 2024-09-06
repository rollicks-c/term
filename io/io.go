package io

import "io"

type Module struct {
	in        io.Reader
	out       io.Writer
	debugMode bool
}

type Option func(*Module)

func New(in io.Reader, out io.Writer, options ...Option) *Module {
	m := &Module{
		in:        in,
		out:       out,
		debugMode: false,
	}
	for _, opt := range options {
		opt(m)
	}
	return m
}
