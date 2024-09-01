package cmdx

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/blackhorseya/godine/pkg/adapterx"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// ServiceCmd represents the service command.
type ServiceCmd struct {
	Use        string
	Short      string
	GetService func(v *viper.Viper) (adapterx.RestfulLegacy, error)
}

// NewServiceCmd creates a new service command.
func NewServiceCmd(use string, short string, svc func(v *viper.Viper) (adapterx.RestfulLegacy, error)) *cobra.Command {
	return (&ServiceCmd{Use: use, Short: short, GetService: svc}).NewCmd()
}

// NewCmd creates a new service command.
func (c *ServiceCmd) NewCmd() *cobra.Command {
	return &cobra.Command{
		Use:   c.Use,
		Short: c.Short,
		Run: func(cmd *cobra.Command, args []string) {
			v := viper.New()

			service, err := c.GetService(v)
			cobra.CheckErr(err)

			err = service.Start()
			cobra.CheckErr(err)

			signalChan := make(chan os.Signal, 1)
			signal.Notify(signalChan, syscall.SIGTERM, syscall.SIGINT)

			<-signalChan

			err = service.AwaitSignal()
			cobra.CheckErr(err)
		},
	}
}
