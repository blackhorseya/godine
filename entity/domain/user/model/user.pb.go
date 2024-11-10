// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: domain/user/model/user.proto

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
	mi := &file_domain_user_model_user_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Address) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Address) ProtoMessage() {}

func (x *Address) ProtoReflect() protoreflect.Message {
	mi := &file_domain_user_model_user_proto_msgTypes[0]
	if x != nil {
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
	return file_domain_user_model_user_proto_rawDescGZIP(), []int{0}
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

// Role is an entity that represents the role of a user.
type Role struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID is the unique identifier of the role.
	// @gotags: bson:"_id"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" bson:"_id"`
	// Name is the name of the role.
	// @gotags: bson:"name"
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty" bson:"name"`
}

func (x *Role) Reset() {
	*x = Role{}
	mi := &file_domain_user_model_user_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Role) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Role) ProtoMessage() {}

func (x *Role) ProtoReflect() protoreflect.Message {
	mi := &file_domain_user_model_user_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Role.ProtoReflect.Descriptor instead.
func (*Role) Descriptor() ([]byte, []int) {
	return file_domain_user_model_user_proto_rawDescGZIP(), []int{1}
}

func (x *Role) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Role) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// Account is an entity that represents a user account.
type Account struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID is the unique identifier of the user account.
	// @gotags: bson:"_id"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" bson:"_id"`
	// Username is the username of the user.
	// @gotags: bson:"username"
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty" bson:"username"`
	// Email is the email address of the user.
	// @gotags: bson:"email"
	Email string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty" bson:"email"`
	// Password is the password of the user.
	// @gotags: json:"-" bson:"password"
	Password string `protobuf:"bytes,4,opt,name=password,proto3" json:"-" bson:"password"`
	// AccessToken is the access token of the user.
	// @gotags: bson:"-"
	AccessToken string `protobuf:"bytes,12,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty" bson:"-"`
	// Address is the address of the user.
	// @gotags: bson:"address"
	Address *Address `protobuf:"bytes,5,opt,name=address,proto3" json:"address,omitempty" bson:"address"`
	// IsActive is the status of the user account.
	// @gotags: bson:"is_active"
	IsActive bool `protobuf:"varint,6,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty" bson:"is_active"`
	// Level is the access level of the user account.
	// @gotags: bson:"level"
	Level uint32 `protobuf:"varint,7,opt,name=Level,proto3" json:"Level,omitempty" bson:"level"`
	// Roles is the list of roles assigned to the user.
	// @gotags: bson:"roles"
	Roles []*Role `protobuf:"bytes,8,rep,name=roles,proto3" json:"roles,omitempty" bson:"roles"`
	// SocialIDMap is the map of social IDs of the user.
	// @gotags: bson:"social_id_map"
	SocialIDMap map[string]string `protobuf:"bytes,9,rep,name=socialIDMap,proto3" json:"socialIDMap,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3" bson:"social_id_map"`
	// CreatedAt is the creation time of the user account.
	// @gotags: bson:"created_at"
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty" bson:"created_at"`
	// UpdatedAt is the update time of the user account.
	// @gotags: bson:"updated_at"
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,11,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty" bson:"updated_at"`
}

func (x *Account) Reset() {
	*x = Account{}
	mi := &file_domain_user_model_user_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Account) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Account) ProtoMessage() {}

func (x *Account) ProtoReflect() protoreflect.Message {
	mi := &file_domain_user_model_user_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Account.ProtoReflect.Descriptor instead.
func (*Account) Descriptor() ([]byte, []int) {
	return file_domain_user_model_user_proto_rawDescGZIP(), []int{2}
}

func (x *Account) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Account) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Account) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Account) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *Account) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *Account) GetAddress() *Address {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *Account) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

func (x *Account) GetLevel() uint32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *Account) GetRoles() []*Role {
	if x != nil {
		return x.Roles
	}
	return nil
}

func (x *Account) GetSocialIDMap() map[string]string {
	if x != nil {
		return x.SocialIDMap
	}
	return nil
}

func (x *Account) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Account) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

var File_domain_user_model_user_proto protoreflect.FileDescriptor

var file_domain_user_model_user_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x66, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x7a, 0x69, 0x70, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x7a, 0x69, 0x70, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x2a, 0x0a,
	0x04, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x80, 0x04, 0x0a, 0x07, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x27, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x4c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x4c, 0x65, 0x76,
	0x65, 0x6c, 0x12, 0x20, 0x0a, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x05, 0x72,
	0x6f, 0x6c, 0x65, 0x73, 0x12, 0x40, 0x0a, 0x0b, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x49, 0x44,
	0x4d, 0x61, 0x70, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x53, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x49,
	0x44, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x73, 0x6f, 0x63, 0x69, 0x61,
	0x6c, 0x49, 0x44, 0x4d, 0x61, 0x70, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x1a, 0x3e, 0x0a, 0x10,
	0x53, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x49, 0x44, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x7e, 0x0a, 0x08,
	0x63, 0x6f, 0x6d, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x42, 0x09, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x62, 0x6c, 0x61, 0x63, 0x6b, 0x68, 0x6f, 0x72, 0x73, 0x65, 0x79, 0x61, 0x2f, 0x67,
	0x6f, 0x64, 0x69, 0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x64, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0xa2, 0x02,
	0x03, 0x55, 0x58, 0x58, 0xaa, 0x02, 0x04, 0x55, 0x73, 0x65, 0x72, 0xca, 0x02, 0x04, 0x55, 0x73,
	0x65, 0x72, 0xe2, 0x02, 0x10, 0x55, 0x73, 0x65, 0x72, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x04, 0x55, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_domain_user_model_user_proto_rawDescOnce sync.Once
	file_domain_user_model_user_proto_rawDescData = file_domain_user_model_user_proto_rawDesc
)

func file_domain_user_model_user_proto_rawDescGZIP() []byte {
	file_domain_user_model_user_proto_rawDescOnce.Do(func() {
		file_domain_user_model_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_domain_user_model_user_proto_rawDescData)
	})
	return file_domain_user_model_user_proto_rawDescData
}

var file_domain_user_model_user_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_domain_user_model_user_proto_goTypes = []any{
	(*Address)(nil),               // 0: user.Address
	(*Role)(nil),                  // 1: user.Role
	(*Account)(nil),               // 2: user.Account
	nil,                           // 3: user.Account.SocialIDMapEntry
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_domain_user_model_user_proto_depIdxs = []int32{
	0, // 0: user.Account.address:type_name -> user.Address
	1, // 1: user.Account.roles:type_name -> user.Role
	3, // 2: user.Account.socialIDMap:type_name -> user.Account.SocialIDMapEntry
	4, // 3: user.Account.created_at:type_name -> google.protobuf.Timestamp
	4, // 4: user.Account.updated_at:type_name -> google.protobuf.Timestamp
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_domain_user_model_user_proto_init() }
func file_domain_user_model_user_proto_init() {
	if File_domain_user_model_user_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_domain_user_model_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_domain_user_model_user_proto_goTypes,
		DependencyIndexes: file_domain_user_model_user_proto_depIdxs,
		MessageInfos:      file_domain_user_model_user_proto_msgTypes,
	}.Build()
	File_domain_user_model_user_proto = out.File
	file_domain_user_model_user_proto_rawDesc = nil
	file_domain_user_model_user_proto_goTypes = nil
	file_domain_user_model_user_proto_depIdxs = nil
}
