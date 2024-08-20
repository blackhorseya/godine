// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.20.3
// source: entity/domain/order/model/order.proto

package model

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Address is a value object that represents the address of a user.
type Address struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Street is the street address of the user.
	// @gotags: bson:"street"
	Street string `protobuf:"bytes,1,opt,name=street,proto3" json:"street,omitempty" bson:"street"`
	// City is the city where the user is located.
	// @gotags: bson:"city"
	City string `protobuf:"bytes,2,opt,name=city,proto3" json:"city,omitempty" bson:"city"`
	// State is the state where the user is located.
	// @gotags: bson:"state"
	State string `protobuf:"bytes,3,opt,name=state,proto3" json:"state,omitempty" bson:"state"`
	// ZipCode is the postal code of the user's location.
	// @gotags: bson:"zip_code"
	ZipCode string `protobuf:"bytes,4,opt,name=zip_code,json=zipCode,proto3" json:"zip_code,omitempty" bson:"zip_code"`
}

func (x *Address) Reset() {
	*x = Address{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entity_domain_order_model_order_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Address) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Address) ProtoMessage() {}

func (x *Address) ProtoReflect() protoreflect.Message {
	mi := &file_entity_domain_order_model_order_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Address.ProtoReflect.Descriptor instead.
func (*Address) Descriptor() ([]byte, []int) {
	return file_entity_domain_order_model_order_proto_rawDescGZIP(), []int{0}
}

func (x *Address) GetStreet() string {
	if x != nil {
		return x.Street
	}
	return ""
}

func (x *Address) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *Address) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *Address) GetZipCode() string {
	if x != nil {
		return x.ZipCode
	}
	return ""
}

var File_entity_domain_order_model_order_proto protoreflect.FileDescriptor

var file_entity_domain_order_model_order_proto_rawDesc = []byte{
	0x0a, 0x25, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x22, 0x66,
	0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x72,
	0x65, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x72, 0x65, 0x65,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x7a,
	0x69, 0x70, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x7a,
	0x69, 0x70, 0x43, 0x6f, 0x64, 0x65, 0x42, 0x3a, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x6c, 0x61, 0x63, 0x6b, 0x68, 0x6f, 0x72, 0x73, 0x65, 0x79,
	0x61, 0x2f, 0x67, 0x6f, 0x64, 0x69, 0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_entity_domain_order_model_order_proto_rawDescOnce sync.Once
	file_entity_domain_order_model_order_proto_rawDescData = file_entity_domain_order_model_order_proto_rawDesc
)

func file_entity_domain_order_model_order_proto_rawDescGZIP() []byte {
	file_entity_domain_order_model_order_proto_rawDescOnce.Do(func() {
		file_entity_domain_order_model_order_proto_rawDescData = protoimpl.X.CompressGZIP(file_entity_domain_order_model_order_proto_rawDescData)
	})
	return file_entity_domain_order_model_order_proto_rawDescData
}

var file_entity_domain_order_model_order_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_entity_domain_order_model_order_proto_goTypes = []any{
	(*Address)(nil), // 0: order.Address
}
var file_entity_domain_order_model_order_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_entity_domain_order_model_order_proto_init() }
func file_entity_domain_order_model_order_proto_init() {
	if File_entity_domain_order_model_order_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_entity_domain_order_model_order_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Address); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_entity_domain_order_model_order_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_entity_domain_order_model_order_proto_goTypes,
		DependencyIndexes: file_entity_domain_order_model_order_proto_depIdxs,
		MessageInfos:      file_entity_domain_order_model_order_proto_msgTypes,
	}.Build()
	File_entity_domain_order_model_order_proto = out.File
	file_entity_domain_order_model_order_proto_rawDesc = nil
	file_entity_domain_order_model_order_proto_goTypes = nil
	file_entity_domain_order_model_order_proto_depIdxs = nil
}
