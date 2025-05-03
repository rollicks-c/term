package io

import (
	"fmt"
	"os"
)

var DebugMode = false

var (
	Info    = Purple
	Warn    = Yellow
	Fatal   = Red
	Success = Green
)

var (
	Default = Color("%s")
	Black   = Color("\033[1;30m%s\033[0m")
	Red     = Color("\033[1;31m%s\033[0m")
	Green   = Color("\033[1;32m%s\033[0m")
	Yellow  = Color("\033[1;33m%s\033[0m")
	Purple  = Color("\033[1;34m%s\033[0m")
	Magenta = Color("\033[1;35m%s\033[0m")
	Teal    = Color("\033[1;36m%s\033[0m")
	White   = Color("\033[1;37m%s\033[0m")
)

type ColPrint = func(...interface{}) string

func Color(colorString string) func(...interface{}) string {
	sprint := func(args ...interface{}) string {
		return fmt.Sprintf(colorString,
			fmt.Sprint(args...))
	}
	return sprint
}

func (m Module) Errorf(format string, v ...interface{}) error {
	msg := fmt.Sprintf(format, v...)
	return fmt.Errorf(Red(msg))
}

func (m Module) TextF(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	if _, err := fmt.Fprint(m.out, White(msg)); err != nil {
		fmt.Print(err)
	}
}

func (m Module) InfoF(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	if _, err := fmt.Fprint(m.out, Info(msg)); err != nil {
		fmt.Print(err)
	}
}

func (m Module) PrintF(col ColPrint, format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	if _, err := fmt.Fprint(m.out, col(msg)); err != nil {
		fmt.Print(err)
	}
}

func (m Module) SPrintF(col ColPrint, format string, v ...interface{}) string {
	msg := fmt.Sprintf(format, v...)
	return col(msg)
}

func (m Module) WarnF(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	if _, err := fmt.Fprint(m.out, Warn(msg)); err != nil {
		fmt.Print(err)
	}
}

func (m Module) Fail(msg string) {
	if _, err := fmt.Fprint(m.out, Fatal(msg)); err != nil {
		fmt.Print(err)
	}
}

func (m Module) FailF(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	if _, err := fmt.Fprint(m.out, Fatal(msg)); err != nil {
		fmt.Print(err)
	}
}
func (m Module) SFailF(format string, v ...interface{}) error {
	msg := fmt.Sprintf(format, v...)
	return fmt.Errorf(msg)
}
func (m Module) FatalF(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	if _, err := fmt.Fprint(m.out, Fatal(msg)); err != nil {
		fmt.Print(err)
	}
	os.Exit(1)
}

func (m Module) DebugF(format string, v ...interface{}) {
	if !m.debugMode {
		return
	}
	msg := fmt.Sprintf(format, v...)
	if _, err := fmt.Fprint(m.out, Red(msg)); err != nil {
		fmt.Print(err)
	}
}

func (m Module) SuccessF(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	if _, err := fmt.Fprint(m.out, Success(msg)); err != nil {
		fmt.Print(err)
	}
}

func (m Module) ConditionalColor(value int) ColPrint {
	if value > 0 {
		return Success
	}
	if value == 0 {
		return White
	}
	return Warn
}

func (m Module) PrintFConditional(value int, format string, v ...interface{}) {
	col := m.ConditionalColor(value)
	m.PrintF(col, format, v...)
}

func (m Module) StatusColor(status string) func(...interface{}) string {
	switch status {
	case "success":
		return Success
	case "failed":
		return Fatal
	case "none":
		return White
	default:
		return Warn
	}

}
