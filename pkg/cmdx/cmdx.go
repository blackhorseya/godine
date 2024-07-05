package cmdx

import (
	"github.com/blackhorseya/godine/pkg/adapterx"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// ServiceCmd represents the service command.
type ServiceCmd struct {
	Use        string
	Short      string
	GetService func(v *viper.Viper) (adapterx.Restful, error)
}

// NewServiceCmd creates a new service command.
func NewServiceCmd(use string, short string, svc func(v *viper.Viper) (adapterx.Restful, error)) *cobra.Command {
	return (&ServiceCmd{Use: use, Short: short, GetService: svc}).NewCmd()
}

// NewCmd creates a new service command.
func (c *ServiceCmd) NewCmd() *cobra.Command {
	return &cobra.Command{
		Use:   c.Use,
		Short: c.Short,
		Run: func(cmd *cobra.Command, args []string) {
			v := viper.GetViper()

			service, err := c.GetService(v)
			cobra.CheckErr(err)

			err = service.Start()
			cobra.CheckErr(err)

			err = service.AwaitSignal()
			cobra.CheckErr(err)
		},
	}
}

// ServiceListCmd represents the service list command.
type ServiceListCmd struct {
	Use   string
	Short string
	Cmds  []*ServiceCmd
}

// NewServiceListCmd creates a new service list command.
func NewServiceListCmd(use string, short string, cmds ...*ServiceCmd) *cobra.Command {
	return (&ServiceListCmd{Use: use, Short: short, Cmds: cmds}).NewCmd()
}

// NewCmd creates a new service list command.
func (c *ServiceListCmd) NewCmd() *cobra.Command {
	return &cobra.Command{
		Use:   c.Use,
		Short: c.Short,
		Run: func(cmd *cobra.Command, args []string) {
			v := viper.GetViper()

			// todo: 2024/7/6|sean|implement me
		},
	}
}
