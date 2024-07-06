package kafkax

import (
	"crypto/tls"

	"github.com/blackhorseya/godine/app/infra/configx"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/plain"
)

// NewWriterWithTopic will create a new kafka writer with topic
func NewWriterWithTopic(app *configx.Application, topic string) (*kafka.Writer, error) {
	return &kafka.Writer{
		Addr:     kafka.TCP(app.Kafka.Brokers...),
		Topic:    topic,
		Balancer: &kafka.Hash{},
		Transport: &kafka.Transport{
			TLS: &tls.Config{
				InsecureSkipVerify: true, //nolint:gosec // skip
			},
			SASL: plain.Mechanism{
				Username: app.Kafka.Username,
				Password: app.Kafka.Password,
			},
		},
	}, nil
}
