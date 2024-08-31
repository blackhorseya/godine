package postgresqlx

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
	var t *timestamppb.Timestamp

	if dbValue != nil {
		switch v := dbValue.(type) {
		case time.Time:
			t = timestamppb.New(v)
		default:
			return fmt.Errorf("unsupported data type: %T", dbValue)
		}

		field.ReflectValueOf(ctx, dst).Set(reflect.ValueOf(t))
	}

	return nil
}

// Value converts google.protobuf.Timestamp to time.Time for database
func (ts TimestampSerializer) Value(
	c context.Context,
	field *schema.Field,
	dst reflect.Value,
	fieldValue interface{},
) (interface{}, error) {
	var (
		t  *timestamppb.Timestamp
		ok bool
	)

	if fieldValue == nil || reflect.ValueOf(fieldValue).IsNil() {
		return nil, nil //nolint:nilnil // return nil for nil value
	}

	if t, ok = fieldValue.(*timestamppb.Timestamp); !ok {
		return nil, fmt.Errorf("unsupported data type: %T", fieldValue)
	}

	return t.AsTime(), nil
}
