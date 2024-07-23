package configx

import (
	"encoding/json"
	"fmt"

	"github.com/blackhorseya/godine/app/infra/authx"
	"github.com/blackhorseya/godine/pkg/logging"
	"github.com/blackhorseya/godine/pkg/netx"
	"github.com/google/uuid"
	"github.com/spf13/viper"
)

// NewApplication is used to create a new application.
func NewApplication(v *viper.Viper, name string) (*Application, error) {
	app := new(Application)

	err := v.UnmarshalKey(name, app)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal application: %w", err)
	}

	return app, nil
}

// Application defines the application struct.
type Application struct {
	ID   string `json:"id" yaml:"id"`
	Name string `json:"name" yaml:"name"`

	Log  logging.Options `json:"log" yaml:"log"`
	HTTP HTTP            `json:"http" yaml:"http"`

	Auth0 authx.Options `json:"auth0" yaml:"auth0"`

	Casbin struct {
		// Enabled is the casbin enabled
		Enabled bool `json:"enabled" yaml:"enabled"`

		// ModelPath is the casbin model path
		ModelPath string `json:"model_path" yaml:"modelPath"`

		// PolicyDriver is the casbin policy driver
		PolicyDriver string `json:"policy_driver" yaml:"policyDriver"`
	} `json:"casbin" yaml:"casbin"`

	Storage struct {
		Mysql struct {
			DSN string `json:"dsn" yaml:"dsn"`
		} `json:"mysql" yaml:"mysql"`

		Postgresql struct {
			DSN string `json:"dsn" yaml:"dsn"`
		} `json:"postgresql" yaml:"postgresql"`

		Mongodb struct {
			DSN string `json:"dsn" yaml:"dsn"`
		} `json:"mongodb" yaml:"mongodb"`

		Redis struct {
			Addr string `json:"addr" yaml:"addr"`
		} `json:"redis" yaml:"redis"`
	} `json:"storage" yaml:"storage"`

	Kafka struct {
		Username string   `json:"username" yaml:"username"`
		Password string   `json:"password" yaml:"password"`
		Brokers  []string `json:"brokers" yaml:"brokers"`
	} `json:"kafka" yaml:"kafka"`

	OTel struct {
		Target string `json:"target" yaml:"target"`
	} `json:"otel" yaml:"otel"`
}

// GetID is used to get the application id.
func (x *Application) GetID() string {
	if x.ID == "" {
		x.ID = uuid.New().String()
	}

	return x.ID
}

func (x *Application) String() string {
	bytes, err := json.MarshalIndent(x, "", "  ")
	if err != nil {
		return err.Error()
	}

	return string(bytes)
}

// HTTP defines the http struct.
type HTTP struct {
	URL  string `json:"url" yaml:"url"`
	Host string `json:"host" yaml:"host"`
	Port int    `json:"port" yaml:"port"`
	Mode string `json:"mode" yaml:"mode"`
}

// GetAddr is used to get the http address.
func (http *HTTP) GetAddr() string {
	if http.Host == "" {
		http.Host = "0.0.0.0"
	}

	if http.Port == 0 {
		http.Port = netx.GetAvailablePort()
	}

	return fmt.Sprintf("%s:%d", http.Host, http.Port)
}
