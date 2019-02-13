// Copyright © 2019 Krakakanok <krakakanok@protonmail.com>

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"gopkg.in/AlecAivazis/survey.v1"
)

var qs = []*survey.Question{
	{
		Name: "type",
		Prompt: &survey.Select{
			Message: "Choose a type:",
			Options: []string{"feat", "fix", "docs", "style", "refactor", "test", "chrore"},
			Default: "feat",
		},
	},
	{
		Name:      "scope",
		Prompt:    &survey.Input{Message: "What is the scope?"},
		Transform: survey.ToLower,
	},
	{
		Name:     "subject",
		Prompt:   &survey.Input{Message: "What is the subject?"},
		Validate: survey.Required,
	},
	{
		Name:   "body",
		Prompt: &survey.Multiline{Message: "What is the body?"},
	},
	{
		Name:   "footer",
		Prompt: &survey.Multiline{Message: "What is the footer?"},
	},
}

// gitWrapperCmd represents the gitWrapper command
var gitWrapperCmd = &cobra.Command{
	Use:   "gitWrapper",
	Short: "Wrapper to git commit command",
	Run: func(cmd *cobra.Command, args []string) {

		message_file := ".git/COMMIT_EDITMSG"

		//message := ""
		var message [5]string

		answers := struct {
			Type    string
			Scope   string
			Subject string
			Body    string
			Footer  string
		}{}

		// perform the questions
		err := survey.Ask(qs, &answers)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		if answers.Scope == "" {
			message[0] = fmt.Sprintf("%s: %s\n", answers.Type, answers.Subject)
		} else {
			message[0] = fmt.Sprintf("%s(%s): %s\n", answers.Type, answers.Scope, answers.Subject)
		}

		if answers.Body != "" {
			message[1] = " "
			message[2] = fmt.Sprintf("\n%s\n", answers.Body)
			message[3] = " "
			message[4] = fmt.Sprintf("\n%s\n", answers.Footer)
		}

		f, err := os.OpenFile(message_file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return
		}

		for _, line := range message {
			f.WriteString(line)
			if err != nil {
				f.Close()
				return
			}
		}

		err = f.Close()
		if err != nil {
			return
		}

	},
}

func init() {
	rootCmd.AddCommand(gitWrapperCmd)
}
