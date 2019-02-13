// Copyright © 2019 Krakakanok <krakakanok@protonmail.com>

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// gitWrapperCmd represents the gitWrapper command
var gitWrapperCmd = &cobra.Command{
	Use:   "gitWrapper",
	Short: "Wrapper to git commit command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("gitWrapper called")
	},
}

func init() {
	rootCmd.AddCommand(gitWrapperCmd)
}
