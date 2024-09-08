package cmd

import (
	"errors"

	"github.com/blackhorseya/godine/app/infra/authz"
	"github.com/blackhorseya/godine/app/infra/configx"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var userID string

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "init environment",
	Run: func(cmd *cobra.Command, args []string) {
		if userID == "" {
			cobra.CheckErr(errors.New("user id is required"))
		}

		config, err := configx.NewConfiguration(viper.GetViper())
		cobra.CheckErr(err)

		app, err := config.GetService("platform")
		cobra.CheckErr(err)

		authz, err := authz.New(app)
		cobra.CheckErr(err)

		err = authz.InitPolicy()
		cobra.CheckErr(err)

		_, err = authz.AddGroupingPolicy(userID, "owner")
		cobra.CheckErr(err)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	initCmd.Flags().StringVar(&userID, "user-id", "", "root user id")
}
