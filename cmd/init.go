package cmd

import (
	"errors"

	"github.com/blackhorseya/godine/app/infra/authz"
	"github.com/blackhorseya/godine/app/infra/configx"
	"github.com/blackhorseya/godine/pkg/logging"
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

		_, err := configx.NewConfiguration(viper.GetViper())
		cobra.CheckErr(err)

		app, err := configx.NewApplication(viper.GetViper(), "userRestful")
		cobra.CheckErr(err)

		err = logging.Init(app.Log)
		cobra.CheckErr(err)

		authz, err := authz.New(app)
		cobra.CheckErr(err)

		authz.ClearPolicy()

		// add policy
		_, err = authz.AddPolicy("owner", "restaurants", "manage")
		cobra.CheckErr(err)

		_, err = authz.AddPolicy("admin", "restaurants", "manage")
		cobra.CheckErr(err)

		_, err = authz.AddPolicy("editor", "restaurants", "edit")
		cobra.CheckErr(err)

		_, err = authz.AddPolicy("viewer", "restaurants", "view")
		cobra.CheckErr(err)

		// add role
		_, err = authz.AddGroupingPolicy(userID, "owner")
		cobra.CheckErr(err)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	initCmd.Flags().StringVar(&userID, "user-id", "", "root user id")
}
