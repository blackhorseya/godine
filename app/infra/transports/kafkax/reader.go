package kafkax

import (
	"crypto/tls"
	"time"

	"github.com/blackhorseya/godine/app/infra/configx"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/plain"
)

// NewReader returns a new Reader instance.
func NewReader() (*kafka.Reader, error) {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers: configx.A.Kafka.Brokers,
		GroupID: configx.A.GetID(),
		Topic:   configx.A.Kafka.Topic,
		Dialer: &kafka.Dialer{
			Timeout:   10 * time.Second,
			DualStack: true,
			TLS: &tls.Config{
				InsecureSkipVerify: true, //nolint:gosec // skip
			},
			SASLMechanism: plain.Mechanism{
				Username: configx.A.Kafka.Username,
				Password: configx.A.Kafka.Password,
			},
		},
	}), nil
}
