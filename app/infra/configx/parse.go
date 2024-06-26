package configx

import (
	"errors"
	"os"
	"strings"

	"github.com/spf13/viper"
)

var (
	C = new(Configuration)
	A = new(Application)
)

// LoadConfig loads the configuration.
func LoadConfig(path string) (err error) {
	v := viper.GetViper()

	if path != "" {
		v.SetConfigFile(path)
	} else {
		home, _ := os.UserHomeDir()
		if home == "" {
			home = "/root"
		}
		v.AddConfigPath(home + "/.config/godine")
		v.SetConfigType("yaml")
		v.SetConfigName(".godine")
	}

	err = v.ReadInConfig()
	if err != nil && !errors.As(err, &viper.ConfigFileNotFoundError{}) {
		return err
	}

	err = v.Unmarshal(&C)
	if err != nil {
		return err
	}

	return nil
}

// LoadApplication loads the application configuration.
func LoadApplication(app *Application) (*Application, error) {
	v := viper.GetViper()

	err := bindEnv(v)
	if err != nil {
		return nil, err
	}

	err = v.Unmarshal(&app)
	if err != nil {
		return nil, err
	}

	A = app

	return A, nil
}

func bindEnv(v *viper.Viper) (err error) {
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err = v.BindEnv("log.level", "LOG_LEVEL")
	if err != nil {
		return err
	}

	err = v.BindEnv("log.format", "LOG_FORMAT")
	if err != nil {
		return err
	}

	err = v.BindEnv("http.mode", "HTTP_MODE")
	if err != nil {
		return err
	}

	err = v.BindEnv("storage.mongodb.dsn", "STORAGE_MONGODB_DSN")
	if err != nil {
		return err
	}

	return nil
}
