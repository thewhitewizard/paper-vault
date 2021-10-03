package util

import "github.com/manifoldco/promptui"

func PasswordPrompt(Label string) (string, error) {
	prompt := promptui.Prompt{
		Label: Label,
		Mask:  '*',
	}
	result, err := prompt.Run()
	return result, err
}
