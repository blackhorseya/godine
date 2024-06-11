package cmd

import (
	"github.com/blackhorseya/godine/adapter/restaurant/restful"
	"github.com/blackhorseya/godine/pkg/cmdx"
	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the server",
}

func init() {
	startCmd.AddCommand(cmdx.NewServiceCmd(
		"restaurant-restful",
		"Start the restaurant restful server",
		restful.New,
	))

	rootCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}