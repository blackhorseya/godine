// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: domain/logistics/model/logistics.proto

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

// DeliveryState is an enum that represents the state of a delivery.
type DeliveryStatus int32

const (
	DeliveryStatus_DELIVERY_STATUS_UNSPECIFIED DeliveryStatus = 0
	DeliveryStatus_DELIVERY_STATUS_PENDING     DeliveryStatus = 1
	DeliveryStatus_DELIVERY_STATUS_IN_TRANSIT  DeliveryStatus = 2
	DeliveryStatus_DELIVERY_STATUS_DELIVERED   DeliveryStatus = 3
)

// Enum value maps for DeliveryStatus.
var (
	DeliveryStatus_name = map[int32]string{
		0: "DELIVERY_STATUS_UNSPECIFIED",
		1: "DELIVERY_STATUS_PENDING",
		2: "DELIVERY_STATUS_IN_TRANSIT",
		3: "DELIVERY_STATUS_DELIVERED",
	}
	DeliveryStatus_value = map[string]int32{
		"DELIVERY_STATUS_UNSPECIFIED": 0,
		"DELIVERY_STATUS_PENDING":     1,
		"DELIVERY_STATUS_IN_TRANSIT":  2,
		"DELIVERY_STATUS_DELIVERED":   3,
	}
)

func (x DeliveryStatus) Enum() *DeliveryStatus {
	p := new(DeliveryStatus)
	*p = x
	return p
}

func (x DeliveryStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DeliveryStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_domain_logistics_model_logistics_proto_enumTypes[0].Descriptor()
}

func (DeliveryStatus) Type() protoreflect.EnumType {
	return &file_domain_logistics_model_logistics_proto_enumTypes[0]
}

func (x DeliveryStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DeliveryStatus.Descriptor instead.
func (DeliveryStatus) EnumDescriptor() ([]byte, []int) {
	return file_domain_logistics_model_logistics_proto_rawDescGZIP(), []int{0}
}

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
		mi := &file_domain_logistics_model_logistics_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Address) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Address) ProtoMessage() {}

func (x *Address) ProtoReflect() protoreflect.Message {
	mi := &file_domain_logistics_model_logistics_proto_msgTypes[0]
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
	return file_domain_logistics_model_logistics_proto_rawDescGZIP(), []int{0}
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

// Delivery is an entity that represents a delivery.
type Delivery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID is the unique identifier of the delivery.
	// @gotags: bson:"_id,omitempty"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" bson:"_id,omitempty"`
	// OrderID is the identifier of the order associated with the delivery.
	// @gotags: bson:"order_id"
	OrderId string `protobuf:"bytes,2,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty" bson:"order_id"`
	// UserID is the identifier of the user who placed the order.
	// @gotags: bson:"user_id"
	UserId string `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty" bson:"user_id"`
	// Address is the address where the delivery is to be made.
	// @gotags: bson:"address"
	Address *Address `protobuf:"bytes,10,opt,name=address,proto3" json:"address,omitempty" bson:"address"`
	// DriverID is the identifier of the driver assigned to the delivery.
	// @gotags: bson:"driver_id"
	DriverId string `protobuf:"bytes,4,opt,name=driver_id,json=driverId,proto3" json:"driver_id,omitempty" bson:"driver_id"`
	// Status is the current status of the delivery (e.g., pending, in transit, delivered).
	// @gotags: bson:"status"
	Status DeliveryStatus `protobuf:"varint,5,opt,name=status,proto3,enum=logistics.DeliveryStatus" json:"status,omitempty" bson:"status"`
	// PickupAt is the timestamp when the delivery was picked up.
	// @gotags: bson:"pickup_at"
	PickupAt *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=pickup_at,json=pickupAt,proto3" json:"pickup_at,omitempty" bson:"pickup_at"`
	// DeliveryAt is the timestamp when the delivery is expected to be delivered.
	// @gotags: bson:"delivery_at"
	DeliveryAt *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=delivery_at,json=deliveryAt,proto3" json:"delivery_at,omitempty" bson:"delivery_at"`
	// CreatedAt is the timestamp when the delivery was created.
	// @gotags: bson:"created_at"
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty" bson:"created_at"`
	// UpdatedAt is the timestamp when the delivery was last updated.
	// @gotags: bson:"updated_at"
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty" bson:"updated_at"`
}

func (x *Delivery) Reset() {
	*x = Delivery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_domain_logistics_model_logistics_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Delivery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Delivery) ProtoMessage() {}

func (x *Delivery) ProtoReflect() protoreflect.Message {
	mi := &file_domain_logistics_model_logistics_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Delivery.ProtoReflect.Descriptor instead.
func (*Delivery) Descriptor() ([]byte, []int) {
	return file_domain_logistics_model_logistics_proto_rawDescGZIP(), []int{1}
}

func (x *Delivery) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Delivery) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *Delivery) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Delivery) GetAddress() *Address {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *Delivery) GetDriverId() string {
	if x != nil {
		return x.DriverId
	}
	return ""
}

func (x *Delivery) GetStatus() DeliveryStatus {
	if x != nil {
		return x.Status
	}
	return DeliveryStatus_DELIVERY_STATUS_UNSPECIFIED
}

func (x *Delivery) GetPickupAt() *timestamppb.Timestamp {
	if x != nil {
		return x.PickupAt
	}
	return nil
}

func (x *Delivery) GetDeliveryAt() *timestamppb.Timestamp {
	if x != nil {
		return x.DeliveryAt
	}
	return nil
}

func (x *Delivery) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Delivery) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

var File_domain_logistics_model_logistics_proto protoreflect.FileDescriptor

var file_domain_logistics_model_logistics_proto_rawDesc = []byte{
	0x0a, 0x26, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69,
	0x63, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69,
	0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74,
	0x69, 0x63, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x66, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x12, 0x19, 0x0a, 0x08, 0x7a, 0x69, 0x70, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x7a, 0x69, 0x70, 0x43, 0x6f, 0x64, 0x65, 0x22, 0xb8, 0x03, 0x0a,
	0x08, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2c, 0x0a,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x64,
	0x72, 0x69, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x49, 0x64, 0x12, 0x31, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x73,
	0x74, 0x69, 0x63, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x37, 0x0a, 0x09, 0x70,
	0x69, 0x63, 0x6b, 0x75, 0x70, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x70, 0x69, 0x63, 0x6b,
	0x75, 0x70, 0x41, 0x74, 0x12, 0x3b, 0x0a, 0x0b, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79,
	0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x41,
	0x74, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x2a, 0x8d, 0x01, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x69,
	0x76, 0x65, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1f, 0x0a, 0x1b, 0x44, 0x45,
	0x4c, 0x49, 0x56, 0x45, 0x52, 0x59, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1b, 0x0a, 0x17, 0x44,
	0x45, 0x4c, 0x49, 0x56, 0x45, 0x52, 0x59, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x50,
	0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x1e, 0x0a, 0x1a, 0x44, 0x45, 0x4c, 0x49,
	0x56, 0x45, 0x52, 0x59, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x49, 0x4e, 0x5f, 0x54,
	0x52, 0x41, 0x4e, 0x53, 0x49, 0x54, 0x10, 0x02, 0x12, 0x1d, 0x0a, 0x19, 0x44, 0x45, 0x4c, 0x49,
	0x56, 0x45, 0x52, 0x59, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x44, 0x45, 0x4c, 0x49,
	0x56, 0x45, 0x52, 0x45, 0x44, 0x10, 0x03, 0x42, 0xa1, 0x01, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x2e,
	0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x42, 0x0e, 0x4c, 0x6f, 0x67, 0x69, 0x73,
	0x74, 0x69, 0x63, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3c, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x6c, 0x61, 0x63, 0x6b, 0x68, 0x6f, 0x72,
	0x73, 0x65, 0x79, 0x61, 0x2f, 0x67, 0x6f, 0x64, 0x69, 0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74,
	0x69, 0x63, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0xa2, 0x02, 0x03, 0x4c, 0x58, 0x58, 0xaa,
	0x02, 0x09, 0x4c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0xca, 0x02, 0x09, 0x4c, 0x6f,
	0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0xe2, 0x02, 0x15, 0x4c, 0x6f, 0x67, 0x69, 0x73, 0x74,
	0x69, 0x63, 0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x09, 0x4c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_domain_logistics_model_logistics_proto_rawDescOnce sync.Once
	file_domain_logistics_model_logistics_proto_rawDescData = file_domain_logistics_model_logistics_proto_rawDesc
)

func file_domain_logistics_model_logistics_proto_rawDescGZIP() []byte {
	file_domain_logistics_model_logistics_proto_rawDescOnce.Do(func() {
		file_domain_logistics_model_logistics_proto_rawDescData = protoimpl.X.CompressGZIP(file_domain_logistics_model_logistics_proto_rawDescData)
	})
	return file_domain_logistics_model_logistics_proto_rawDescData
}

var file_domain_logistics_model_logistics_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_domain_logistics_model_logistics_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_domain_logistics_model_logistics_proto_goTypes = []any{
	(DeliveryStatus)(0),           // 0: logistics.DeliveryStatus
	(*Address)(nil),               // 1: logistics.Address
	(*Delivery)(nil),              // 2: logistics.Delivery
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_domain_logistics_model_logistics_proto_depIdxs = []int32{
	1, // 0: logistics.Delivery.address:type_name -> logistics.Address
	0, // 1: logistics.Delivery.status:type_name -> logistics.DeliveryStatus
	3, // 2: logistics.Delivery.pickup_at:type_name -> google.protobuf.Timestamp
	3, // 3: logistics.Delivery.delivery_at:type_name -> google.protobuf.Timestamp
	3, // 4: logistics.Delivery.created_at:type_name -> google.protobuf.Timestamp
	3, // 5: logistics.Delivery.updated_at:type_name -> google.protobuf.Timestamp
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_domain_logistics_model_logistics_proto_init() }
func file_domain_logistics_model_logistics_proto_init() {
	if File_domain_logistics_model_logistics_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_domain_logistics_model_logistics_proto_msgTypes[0].Exporter = func(v any, i int) any {
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
		file_domain_logistics_model_logistics_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Delivery); i {
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
			RawDescriptor: file_domain_logistics_model_logistics_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_domain_logistics_model_logistics_proto_goTypes,
		DependencyIndexes: file_domain_logistics_model_logistics_proto_depIdxs,
		EnumInfos:         file_domain_logistics_model_logistics_proto_enumTypes,
		MessageInfos:      file_domain_logistics_model_logistics_proto_msgTypes,
	}.Build()
	File_domain_logistics_model_logistics_proto = out.File
	file_domain_logistics_model_logistics_proto_rawDesc = nil
	file_domain_logistics_model_logistics_proto_goTypes = nil
	file_domain_logistics_model_logistics_proto_depIdxs = nil
}
