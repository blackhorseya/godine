package cmd

import (
	logistics "github.com/blackhorseya/godine/adapter/logistics/restful"
	notify "github.com/blackhorseya/godine/adapter/notify/restful"
	order "github.com/blackhorseya/godine/adapter/order/restful"
	restaurant "github.com/blackhorseya/godine/adapter/restaurant/restful"
	user "github.com/blackhorseya/godine/adapter/user/restful"
	"github.com/blackhorseya/godine/pkg/cmdx"
	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the server",
}

func init() {
	restaurantRestfulCmd := cmdx.NewServiceCmd(
		"restaurant-restful",
		"Start the restaurant restful server",
		restaurant.New,
	)

	orderRestfulCmd := cmdx.NewServiceCmd(
		"order-restful",
		"Start the order restful server",
		order.New,
	)

	userRestfulCmd := cmdx.NewServiceCmd(
		"user-restful",
		"Start the user restful server",
		user.New,
	)

	logisticsRestfulCmd := cmdx.NewServiceCmd(
		"logistics-restful",
		"Start the logistics restful server",
		logistics.New,
	)

	notifyRestfulCmd := cmdx.NewServiceCmd(
		"notify-restful",
		"Start the notify restful server",
		notify.New,
	)

	startCmd.AddCommand(restaurantRestfulCmd, orderRestfulCmd, userRestfulCmd, logisticsRestfulCmd, notifyRestfulCmd)

	rootCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
