package io

import (
	"bufio"
	"fmt"
	"github.com/manifoldco/promptui"
	"strconv"
	"strings"
)

type PromptOption func(*promptui.Prompt)

func WithDefault(defaultValue string) PromptOption {
	return func(p *promptui.Prompt) {
		p.Default = defaultValue
	}
}

func (m Module) Confirm(prompt string, v ...interface{}) bool {

	msg := fmt.Sprintf("%s (y/n): ", prompt)
	m.WarnF(msg, v...)
	var response string
	if _, err := fmt.Fscan(m.in, &response); err != nil {
		m.FailF("unable to read confirmation response: %v", err)
	}
	response = strings.TrimSpace(strings.ToLower(response))
	return response == "y"

}

func (m Module) ChooseManual(prompt string, min, max int, v ...interface{}) (int, bool) {

	// prompt
	msg := fmt.Sprintf("%s: ", prompt)
	m.WarnF(msg, v...)

	// get response
	buf := bufio.NewReader(m.in)
	response, err := buf.ReadString('\n')
	response = strings.ReplaceAll(response, "\n", "")

	// no response
	if strings.TrimSpace(response) == "" {
		return -1, false
	}

	// validate response
	choice, err := strconv.Atoi(response)
	if err != nil {
		m.FailF("invalid input: %s\n\n", response)
		return -1, false
	}
	if choice < min || choice > max {
		m.FailF("invalid input: %s\n\n", response)
		return -1, false
	}

	// valid response
	response = strings.TrimSpace(strings.ToLower(response))
	return choice, true

}

func (m Module) Choose(prompt string, list map[string]any) (any, error) {

	candidates := make([]string, 0, len(list))
	for k := range list {
		candidates = append(candidates, k)
	}

	selector := promptui.Select{
		Label:     prompt,
		Items:     candidates,
		CursorPos: 0,
		Templates: &promptui.SelectTemplates{Details: ""},
	}
	_, result, err := selector.Run()
	if err != nil {
		return nil, err
	}
	selectedTask, ok := list[result]
	if !ok {
		return nil, fmt.Errorf("invalid task: %s", result)
	}
	return selectedTask, nil
}

func (m Module) PromptString(text string, options ...PromptOption) (string, error) {
	prompt := promptui.Prompt{
		Label: text,
	}
	for _, opt := range options {
		opt(&prompt)
	}
	response, err := prompt.Run()
	if err != nil {
		return "", err
	}

	response = strings.TrimSpace(response)
	return response, nil
}

func (m Module) PromptInt(text string, options ...PromptOption) (int, error) {
	prompt := promptui.Prompt{
		Label: text,
	}
	for _, opt := range options {
		opt(&prompt)
	}
	response, err := prompt.Run()
	if err != nil {
		return 0, err
	}

	value, err := strconv.Atoi(response)
	if err != nil {
		return 0, err
	}
	return value, nil
}

func (m Module) PromptPassword(text string) (string, error) {
	prompt := promptui.Prompt{
		Label: text,
		Mask:  '*',
	}
	response, err := prompt.Run()
	if err != nil {
		return "", err
	}

	response = strings.TrimSpace(response)
	return response, nil

}
