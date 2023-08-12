/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	th "gitea.slauson.io/mslauson/ggit/thelper"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

func initialTagChoices() th.TeaChoices {
	return th.TeaChoices{
		// Our to-do list is a grocery list
		Choices: []string{"Patch", "Minor", "Major"},

		// A map which indicates which choices are selected. We're using
		// the  map like a mathematical set. The keys refer to the indexes
		// of the `choices` slice, above.
		Selected: make(map[int]struct{}),
	}
}

// tagCmd represents the tag command
// This is used for tagging the repo
var tagCmd = &cobra.Command{
	Use:   "tag",
	Short: "tag helper",
	Long:  `tag helper for tagging the repo`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("tag called")
	},
}

func init() {
	rootCmd.AddCommand(tagCmd)
	p := tea.NewProgram(initialTagChoices())
}
