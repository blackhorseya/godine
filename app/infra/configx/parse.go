package configx

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

var (
	A = new(Application)
)

// NewConfiguration creates a new configuration.
func NewConfiguration(v *viper.Viper) (*Configuration, error) {
	configFile := viper.GetString("config")
	if configFile == "" {
		home, _ := os.UserHomeDir()
		if home == "" {
			home = "/root"
		}
		configFile = home + "/.config/godine/.godine.yaml"
	}

	v.SetConfigFile(configFile)

	err := v.ReadInConfig()
	if err != nil && !errors.As(err, &viper.ConfigFileNotFoundError{}) {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	config := new(Configuration)
	err = v.Unmarshal(&config)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal configuration: %w", err)
	}

	return config, nil
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
