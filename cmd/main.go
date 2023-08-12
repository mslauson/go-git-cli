package cmd

import (
	"github.com/spf13/cobra"
)

// rootCmd ggit provides a CLI for various git operations
var rootCmd = &cobra.Command{
	Use:   "ggit",
	Short: "Go git helper",
	Long:  `Go git helper is a CLI for various git operations`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}
