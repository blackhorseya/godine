package kafkax

import (
	"crypto/tls"

	"github.com/blackhorseya/godine/app/infra/configx"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/plain"
)

// NewWriter returns a new Writer instance.
func NewWriter() (*kafka.Writer, error) {
	return &kafka.Writer{
		Addr:     kafka.TCP(configx.A.Kafka.Brokers...),
		Balancer: &kafka.Hash{},
		Transport: &kafka.Transport{
			TLS: &tls.Config{
				InsecureSkipVerify: true, //nolint:gosec // skip
			},
			SASL: plain.Mechanism{
				Username: configx.A.Kafka.Username,
				Password: configx.A.Kafka.Password,
			},
		},
	}, nil
}