// Copyright © 2019 Krakakanok <krakakanok@protonmail.com>

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// gitWrapperCmd represents the gitWrapper command
var gitWrapperCmd = &cobra.Command{
	Use:   "gitWrapper",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("gitWrapper called")
	},
}

func init() {
	rootCmd.AddCommand(gitWrapperCmd)
}
