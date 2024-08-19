package term

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Confirm(prompt string, v ...interface{}) bool {

	msg := fmt.Sprintf("%s (y/n): ", prompt)
	Warnf(msg, v...)
	var response string
	if _, err := fmt.Scan(&response); err != nil {
		Failf("Unable to read confirmation response: %v", err)
	}

	response = strings.TrimSpace(strings.ToLower(response))
	return response == "y"

}

func Choose(prompt string, min, max int, v ...interface{}) (int, bool) {

	// prompt
	msg := fmt.Sprintf("%s: ", prompt)
	Warnf(msg, v...)

	// get response
	buf := bufio.NewReader(os.Stdin)
	response, err := buf.ReadString('\n')
	response = strings.ReplaceAll(response, "\n", "")

	// no response
	if strings.TrimSpace(response) == "" {
		return -1, false
	}

	// validate response
	choice, err := strconv.Atoi(response)
	if err != nil {
		Failf("invalid input: %s\n\n", response)
		return -1, false
	}
	if choice < min || choice > max {
		Failf("invalid input: %s\n\n", response)
		return -1, false
	}

	// valid response
	response = strings.TrimSpace(strings.ToLower(response))
	return choice, true

}
