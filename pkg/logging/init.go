package logging

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Options is the logging options.
type Options struct {
	// Level is the log level. options: debug, info, warn, error, dpanic, panic, fatal (default: info)
	Level string `json:"level" yaml:"level"`

	// Format is the log format. options: json, console (default: console)
	Format string `json:"format" yaml:"format"`
}

// Init initializes the logging instance.
func Init(options Options) error {
	level := zap.NewAtomicLevel()
	err := level.UnmarshalText([]byte(options.Level))
	if err != nil {
		return err
	}

	cw := zapcore.Lock(os.Stdout)
	zapConfig := zap.NewDevelopmentEncoderConfig()
	zapConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	enc := zapcore.NewConsoleEncoder(zapConfig)
	if options.Format == "json" {
		zapConfig = zap.NewProductionEncoderConfig()
		enc = zapcore.NewJSONEncoder(zapConfig)
	}

	cores := make([]zapcore.Core, 0)
	cores = append(cores, zapcore.NewCore(enc, cw, level))

	core := zapcore.NewTee(cores...)
	logger := zap.New(core)

	zap.ReplaceGlobals(logger)

	return nil
}
