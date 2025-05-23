// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: domain/payment/model/payment.proto

package model

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// PaymentAmount is a value object that represents the amount of payment
type PaymentAmount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// value is the amount of payment
	// @gotags: bson:"value"
	Value float64 `protobuf:"fixed64,1,opt,name=value,proto3" json:"value,omitempty" bson:"value"`
	// currency is the currency of payment
	// @gotags: bson:"currency"
	Currency string `protobuf:"bytes,2,opt,name=currency,proto3" json:"currency,omitempty" bson:"currency"`
}

func (x *PaymentAmount) Reset() {
	*x = PaymentAmount{}
	mi := &file_domain_payment_model_payment_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PaymentAmount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentAmount) ProtoMessage() {}

func (x *PaymentAmount) ProtoReflect() protoreflect.Message {
	mi := &file_domain_payment_model_payment_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentAmount.ProtoReflect.Descriptor instead.
func (*PaymentAmount) Descriptor() ([]byte, []int) {
	return file_domain_payment_model_payment_proto_rawDescGZIP(), []int{0}
}

func (x *PaymentAmount) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *PaymentAmount) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

// Payment is an entity that represents a payment
type Payment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id is the identifier of payment
	// @gotags: bson:"_id,omitempty"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" bson:"_id,omitempty"`
	// amount is the amount of payment
	// @gotags: bson:"amount"
	Amount *PaymentAmount `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount,omitempty" bson:"amount"`
	// createdAt is the time when payment created
	// @gotags: bson:"created_at,omitempty"
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty" bson:"created_at,omitempty"`
	// updatedAt is the time when payment updated
	// @gotags: bson:"updated_at,omitempty"
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	// OrderID is the identifier of order
	// @gotags: bson:"order_id"
	OrderId string `protobuf:"bytes,5,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty" bson:"order_id"`
	// UserID is the identifier of user
	// @gotags: bson:"user_id"
	UserId string `protobuf:"bytes,6,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty" bson:"user_id"`
}

func (x *Payment) Reset() {
	*x = Payment{}
	mi := &file_domain_payment_model_payment_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Payment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Payment) ProtoMessage() {}

func (x *Payment) ProtoReflect() protoreflect.Message {
	mi := &file_domain_payment_model_payment_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Payment.ProtoReflect.Descriptor instead.
func (*Payment) Descriptor() ([]byte, []int) {
	return file_domain_payment_model_payment_proto_rawDescGZIP(), []int{1}
}

func (x *Payment) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Payment) GetAmount() *PaymentAmount {
	if x != nil {
		return x.Amount
	}
	return nil
}

func (x *Payment) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Payment) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Payment) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *Payment) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

var File_domain_payment_model_payment_proto protoreflect.FileDescriptor

var file_domain_payment_model_payment_proto_rawDesc = []byte{
	0x0a, 0x22, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x41,
	0x0a, 0x0d, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x79, 0x22, 0xf3, 0x01, 0x0a, 0x07, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2e, 0x0a,
	0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x41,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x39, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x42, 0x93, 0x01, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x2e,
	0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x0c, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x6c, 0x61, 0x63, 0x6b, 0x68, 0x6f, 0x72, 0x73, 0x65, 0x79, 0x61,
	0x2f, 0x67, 0x6f, 0x64, 0x69, 0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0xa2, 0x02, 0x03, 0x50, 0x58, 0x58, 0xaa, 0x02, 0x07, 0x50, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0xca, 0x02, 0x07, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0xe2, 0x02, 0x13,
	0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x07, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_domain_payment_model_payment_proto_rawDescOnce sync.Once
	file_domain_payment_model_payment_proto_rawDescData = file_domain_payment_model_payment_proto_rawDesc
)

func file_domain_payment_model_payment_proto_rawDescGZIP() []byte {
	file_domain_payment_model_payment_proto_rawDescOnce.Do(func() {
		file_domain_payment_model_payment_proto_rawDescData = protoimpl.X.CompressGZIP(file_domain_payment_model_payment_proto_rawDescData)
	})
	return file_domain_payment_model_payment_proto_rawDescData
}

var file_domain_payment_model_payment_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_domain_payment_model_payment_proto_goTypes = []any{
	(*PaymentAmount)(nil),         // 0: payment.PaymentAmount
	(*Payment)(nil),               // 1: payment.Payment
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_domain_payment_model_payment_proto_depIdxs = []int32{
	0, // 0: payment.Payment.amount:type_name -> payment.PaymentAmount
	2, // 1: payment.Payment.created_at:type_name -> google.protobuf.Timestamp
	2, // 2: payment.Payment.updated_at:type_name -> google.protobuf.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_domain_payment_model_payment_proto_init() }
func file_domain_payment_model_payment_proto_init() {
	if File_domain_payment_model_payment_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_domain_payment_model_payment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_domain_payment_model_payment_proto_goTypes,
		DependencyIndexes: file_domain_payment_model_payment_proto_depIdxs,
		MessageInfos:      file_domain_payment_model_payment_proto_msgTypes,
	}.Build()
	File_domain_payment_model_payment_proto = out.File
	file_domain_payment_model_payment_proto_rawDesc = nil
	file_domain_payment_model_payment_proto_goTypes = nil
	file_domain_payment_model_payment_proto_depIdxs = nil
}
