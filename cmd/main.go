package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().
	// 	StringVar(&cfgFile, "config", "", "config file (default is $HOME/.go-gopher-cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initConfig() {
	// Find home directory.
	// home, err := os.UserHomeDir()
	// cobra.CheckErr(err)

	// Search config in home directory with name ".go-gopher-cli" (without extension).
	// viper.AddConfigPath(home)
	// viper.SetConfigType("yaml")
	// viper.SetConfigName(".go-gopher-cli")

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
