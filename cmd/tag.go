/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"gitea.slauson.io/mslauson/ggit/hlpr"
	"github.com/spf13/cobra"
)

func initialTagChoices() hlpr.TeaChoices {
	return hlpr.TeaChoices{
		// Our to-do list is a grocery list
		choices: []string{"Buy carrots", "Buy celery", "Buy kohlrabi"},

		// A map which indicates which choices are selected. We're using
		// the  map like a mathematical set. The keys refer to the indexes
		// of the `choices` slice, above.
		selected: make(map[int]struct{}),
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// tagCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// tagCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
