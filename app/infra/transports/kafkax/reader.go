package kafkax

import (
	"crypto/tls"
	"time"

	"github.com/blackhorseya/godine/app/infra/configx"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/plain"
)

// NewReaderWithTopic returns a new Reader instance with the specified topic.
func NewReaderWithTopic(topic string) (*kafka.Reader, error) {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers: configx.A.Kafka.Brokers,
		GroupID: configx.A.GetID(),
		Topic:   topic,
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
