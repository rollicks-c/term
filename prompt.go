package term

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"strconv"
	"strings"
)

func PromptSecret(prompt string) (string, error) {

	dataPrompt := promptui.Prompt{
		Label:   prompt,
		Default: "",
		Mask:    '*',
		Validate: func(s string) error {
			if strings.TrimSpace(s) == "" {
				return fmt.Errorf("value required")
			}
			return nil
		},
	}
	value, err := dataPrompt.Run()
	if err != nil {
		return "", err
	}
	return value, nil

}

func PromptString(prompt, defaultValue string) (string, error) {

	dataPrompt := promptui.Prompt{
		Label:   fmt.Sprintf("%s (%s)", prompt, defaultValue),
		Default: defaultValue,
		Validate: func(s string) error {
			if strings.TrimSpace(s) == "" {
				return fmt.Errorf("value required")
			}
			return nil
		},
	}
	value, err := dataPrompt.Run()
	if err != nil {
		return "", err
	}
	return value, nil

}

func PromptStringOptional(prompt, defaultValue string) (string, error) {

	dataPrompt := promptui.Prompt{
		Label:   fmt.Sprintf("%s (%s)", prompt, defaultValue),
		Default: defaultValue,
	}
	value, err := dataPrompt.Run()
	if err != nil {
		return "", err
	}
	return value, nil

}

func PromptInt(prompt string, defaultValue int) (int, error) {

	dataPrompt := promptui.Prompt{
		Label:   fmt.Sprintf("%s (%d)", prompt, defaultValue),
		Default: fmt.Sprintf("%d", defaultValue),
		Validate: func(s string) error {
			if _, err := strconv.Atoi(s); err != nil {
				return err
			}
			return nil
		},
	}
	valueRaw, err := dataPrompt.Run()
	if err != nil {
		return defaultValue, err
	}
	value, err := strconv.Atoi(valueRaw)
	if err != nil {
		return defaultValue, err
	}
	return value, nil

}

func PromptFloat(prompt string, defaultValue float64) (float64, error) {

	dataPrompt := promptui.Prompt{
		Label:   fmt.Sprintf("%s (%f)", prompt, defaultValue),
		Default: fmt.Sprintf("%f", defaultValue),
		Validate: func(s string) error {
			if _, err := strconv.ParseFloat(s, 64); err != nil {
				return err
			}
			return nil
		},
	}
	valueRaw, err := dataPrompt.Run()
	if err != nil {
		return defaultValue, err
	}
	value, err := strconv.ParseFloat(valueRaw, 64)
	if err != nil {
		return defaultValue, err
	}
	return value, nil

}
