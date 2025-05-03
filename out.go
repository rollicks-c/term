package term

import "fmt"

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
	Gray    = Color("\033[1;90m%s\033[0m")
)

type ColPrint = func(...interface{}) string

func Color(colorString string) func(...interface{}) string {
	sprint := func(args ...interface{}) string {
		return fmt.Sprintf(colorString,
			fmt.Sprint(args...))
	}
	return sprint
}

func Errorf(format string, v ...interface{}) error {
	msg := fmt.Sprintf(format, v...)
	return fmt.Errorf(Red(msg))
}

func Textf(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	fmt.Print(White(msg))
}

func Infof(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	fmt.Print(Info(msg))
}

func Printf(col ColPrint, format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	fmt.Print(col(msg))
}
func Print(col ColPrint, test string) {
	fmt.Print(col(test))
}

func Sprintf(col ColPrint, format string, v ...interface{}) string {
	msg := fmt.Sprintf(format, v...)
	return col(msg)
}

func Warnf(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	fmt.Print(Warn(msg))
}

func Fail(msg string) {
	fmt.Print(Fatal(msg))
}

func Failf(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	fmt.Print(Fatal(msg))
}
func Sfailf(format string, v ...interface{}) error {
	msg := fmt.Sprintf(format, v...)
	fmt.Print(Fatal(msg))
	return fmt.Errorf(msg)
}

func DebugF(format string, v ...interface{}) {
	if !DebugMode {
		return
	}
	msg := fmt.Sprintf(format, v...)
	fmt.Print(Red(msg))
}

func Successf(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	fmt.Print(Success(msg))
}

func ConditionalColor(value int) ColPrint {
	if value > 0 {
		return Success
	}
	if value == 0 {
		return White
	}
	return Warn
}

func PrintfConditional(value int, format string, v ...interface{}) {
	col := ConditionalColor(value)
	Printf(col, format, v...)
}

func StatusColor(status string) func(...interface{}) string {
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
