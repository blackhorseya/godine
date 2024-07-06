package cmd

import (
	"log"

	logistics "github.com/blackhorseya/godine/adapter/logistics/restful"
	notify "github.com/blackhorseya/godine/adapter/notify/restful"
	order "github.com/blackhorseya/godine/adapter/order/restful"
	restaurant "github.com/blackhorseya/godine/adapter/restaurant/restful"
	user "github.com/blackhorseya/godine/adapter/user/restful"
	"github.com/blackhorseya/godine/pkg/adapterx"
	"github.com/blackhorseya/godine/pkg/cmdx"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/sync/errgroup"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the server",
	Run: func(cmd *cobra.Command, args []string) {
		services := []func(*viper.Viper) (adapterx.Restful, error){
			restaurant.New,
			order.New,
			user.New,
			logistics.New,
			notify.New,
		}

		var g errgroup.Group

		for _, getService := range services {
			getService := getService // capture range variable
			g.Go(func() error {
				v := viper.New()
				service, err := getService(v)
				if err != nil {
					log.Printf("Failed to initialize service: %v", err)
					return err
				}

				if err = service.Start(); err != nil {
					log.Printf("Failed to start service: %v", err)
					return err
				}

				if err = service.AwaitSignal(); err != nil {
					log.Printf("Service encountered an error: %v", err)
					return err
				}

				return nil
			})
		}

		if err := g.Wait(); err != nil {
			log.Fatalf("Failed to start all services: %v", err)
		}
	},
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
