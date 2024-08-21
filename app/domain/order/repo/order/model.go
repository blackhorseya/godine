package order

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm/schema"
)

// TimestampSerializer handles serialization for google.protobuf.Timestamp
type TimestampSerializer struct{}

// Scan converts database value to google.protobuf.Timestamp
func (ts TimestampSerializer) Scan(
	ctx context.Context,
	field *schema.Field,
	dst reflect.Value,
	dbValue interface{},
) error {
	if dbValue == nil {
		return nil
	}

	if value, ok := dbValue.(time.Time); ok {
		timestamp := timestamppb.New(value)
		if dst.CanSet() {
			dst.Set(reflect.ValueOf(timestamp))
		}
		return nil
	}
	return fmt.Errorf("unsupported data type: %T", dbValue)
}

// Value converts google.protobuf.Timestamp to time.Time for database
func (ts TimestampSerializer) Value(
	ctx context.Context,
	field *schema.Field,
	dst reflect.Value,
	fieldValue interface{},
) (interface{}, error) {
	switch value := fieldValue.(type) {
	case timestamppb.Timestamp:
		return value.AsTime(), nil
	case *timestamppb.Timestamp:
		if value != nil {
			return value.AsTime(), nil
		}
	}
	return nil, fmt.Errorf("invalid field value: %#v", fieldValue)
}
