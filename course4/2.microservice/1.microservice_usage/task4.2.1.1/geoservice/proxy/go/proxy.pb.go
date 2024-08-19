// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.12.4
// source: proxy.proto

package proxyservice

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

type RequestAddress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lat float64 `protobuf:"fixed64,1,opt,name=Lat,proto3" json:"Lat,omitempty"`
	Lng float64 `protobuf:"fixed64,2,opt,name=Lng,proto3" json:"Lng,omitempty"`
}

func (x *RequestAddress) Reset() {
	*x = RequestAddress{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proxy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestAddress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestAddress) ProtoMessage() {}

func (x *RequestAddress) ProtoReflect() protoreflect.Message {
	mi := &file_proxy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestAddress.ProtoReflect.Descriptor instead.
func (*RequestAddress) Descriptor() ([]byte, []int) {
	return file_proxy_proto_rawDescGZIP(), []int{0}
}

func (x *RequestAddress) GetLat() float64 {
	if x != nil {
		return x.Lat
	}
	return 0
}

func (x *RequestAddress) GetLng() float64 {
	if x != nil {
		return x.Lng
	}
	return 0
}

type Address struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HouseNumber string `protobuf:"bytes,1,opt,name=House_number,json=HouseNumber,proto3" json:"House_number,omitempty"`
	Road        string `protobuf:"bytes,2,opt,name=Road,proto3" json:"Road,omitempty"`
	Suburb      string `protobuf:"bytes,3,opt,name=Suburb,proto3" json:"Suburb,omitempty"`
	City        string `protobuf:"bytes,4,opt,name=City,proto3" json:"City,omitempty"`
	State       string `protobuf:"bytes,5,opt,name=State,proto3" json:"State,omitempty"`
	Country     string `protobuf:"bytes,6,opt,name=Country,proto3" json:"Country,omitempty"`
}

func (x *Address) Reset() {
	*x = Address{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proxy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Address) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Address) ProtoMessage() {}

func (x *Address) ProtoReflect() protoreflect.Message {
	mi := &file_proxy_proto_msgTypes[1]
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
	return file_proxy_proto_rawDescGZIP(), []int{1}
}

func (x *Address) GetHouseNumber() string {
	if x != nil {
		return x.HouseNumber
	}
	return ""
}

func (x *Address) GetRoad() string {
	if x != nil {
		return x.Road
	}
	return ""
}

func (x *Address) GetSuburb() string {
	if x != nil {
		return x.Suburb
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

func (x *Address) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

type ResponceAddress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address *Address `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *ResponceAddress) Reset() {
	*x = ResponceAddress{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proxy_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponceAddress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponceAddress) ProtoMessage() {}

func (x *ResponceAddress) ProtoReflect() protoreflect.Message {
	mi := &file_proxy_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponceAddress.ProtoReflect.Descriptor instead.
func (*ResponceAddress) Descriptor() ([]byte, []int) {
	return file_proxy_proto_rawDescGZIP(), []int{2}
}

func (x *ResponceAddress) GetAddress() *Address {
	if x != nil {
		return x.Address
	}
	return nil
}

type Coords struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lat string `protobuf:"bytes,1,opt,name=Lat,proto3" json:"Lat,omitempty"`
	Lon string `protobuf:"bytes,2,opt,name=Lon,proto3" json:"Lon,omitempty"`
}

func (x *Coords) Reset() {
	*x = Coords{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proxy_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Coords) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Coords) ProtoMessage() {}

func (x *Coords) ProtoReflect() protoreflect.Message {
	mi := &file_proxy_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Coords.ProtoReflect.Descriptor instead.
func (*Coords) Descriptor() ([]byte, []int) {
	return file_proxy_proto_rawDescGZIP(), []int{3}
}

func (x *Coords) GetLat() string {
	if x != nil {
		return x.Lat
	}
	return ""
}

func (x *Coords) GetLon() string {
	if x != nil {
		return x.Lon
	}
	return ""
}

type GetCoords struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Coords []*Coords `protobuf:"bytes,1,rep,name=coords,proto3" json:"coords,omitempty"`
}

func (x *GetCoords) Reset() {
	*x = GetCoords{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proxy_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCoords) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCoords) ProtoMessage() {}

func (x *GetCoords) ProtoReflect() protoreflect.Message {
	mi := &file_proxy_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCoords.ProtoReflect.Descriptor instead.
func (*GetCoords) Descriptor() ([]byte, []int) {
	return file_proxy_proto_rawDescGZIP(), []int{4}
}

func (x *GetCoords) GetCoords() []*Coords {
	if x != nil {
		return x.Coords
	}
	return nil
}

type ID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *ID) Reset() {
	*x = ID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proxy_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ID) ProtoMessage() {}

func (x *ID) ProtoReflect() protoreflect.Message {
	mi := &file_proxy_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ID.ProtoReflect.Descriptor instead.
func (*ID) Descriptor() ([]byte, []int) {
	return file_proxy_proto_rawDescGZIP(), []int{5}
}

func (x *ID) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type UserName struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=Username,proto3" json:"Username,omitempty"`
}

func (x *UserName) Reset() {
	*x = UserName{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proxy_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserName) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserName) ProtoMessage() {}

func (x *UserName) ProtoReflect() protoreflect.Message {
	mi := &file_proxy_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserName.ProtoReflect.Descriptor instead.
func (*UserName) Descriptor() ([]byte, []int) {
	return file_proxy_proto_rawDescGZIP(), []int{6}
}

func (x *UserName) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type IsValid struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsValid bool `protobuf:"varint,1,opt,name=isValid,proto3" json:"isValid,omitempty"`
}

func (x *IsValid) Reset() {
	*x = IsValid{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proxy_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsValid) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsValid) ProtoMessage() {}

func (x *IsValid) ProtoReflect() protoreflect.Message {
	mi := &file_proxy_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsValid.ProtoReflect.Descriptor instead.
func (*IsValid) Descriptor() ([]byte, []int) {
	return file_proxy_proto_rawDescGZIP(), []int{7}
}

func (x *IsValid) GetIsValid() bool {
	if x != nil {
		return x.IsValid
	}
	return false
}

type Token struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *Token) Reset() {
	*x = Token{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proxy_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Token) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Token) ProtoMessage() {}

func (x *Token) ProtoReflect() protoreflect.Message {
	mi := &file_proxy_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Token.ProtoReflect.Descriptor instead.
func (*Token) Descriptor() ([]byte, []int) {
	return file_proxy_proto_rawDescGZIP(), []int{8}
}

func (x *Token) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Login    string `protobuf:"bytes,1,opt,name=Login,proto3" json:"Login,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proxy_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_proxy_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_proxy_proto_rawDescGZIP(), []int{9}
}

func (x *User) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *User) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type IsRegistered struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsRegistered bool `protobuf:"varint,1,opt,name=isRegistered,proto3" json:"isRegistered,omitempty"`
}

func (x *IsRegistered) Reset() {
	*x = IsRegistered{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proxy_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsRegistered) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsRegistered) ProtoMessage() {}

func (x *IsRegistered) ProtoReflect() protoreflect.Message {
	mi := &file_proxy_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsRegistered.ProtoReflect.Descriptor instead.
func (*IsRegistered) Descriptor() ([]byte, []int) {
	return file_proxy_proto_rawDescGZIP(), []int{10}
}

func (x *IsRegistered) GetIsRegistered() bool {
	if x != nil {
		return x.IsRegistered
	}
	return false
}

type IsLogined struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsLogined bool   `protobuf:"varint,1,opt,name=isLogined,proto3" json:"isLogined,omitempty"`
	Token     *Token `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *IsLogined) Reset() {
	*x = IsLogined{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proxy_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsLogined) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsLogined) ProtoMessage() {}

func (x *IsLogined) ProtoReflect() protoreflect.Message {
	mi := &file_proxy_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsLogined.ProtoReflect.Descriptor instead.
func (*IsLogined) Descriptor() ([]byte, []int) {
	return file_proxy_proto_rawDescGZIP(), []int{11}
}

func (x *IsLogined) GetIsLogined() bool {
	if x != nil {
		return x.IsLogined
	}
	return false
}

func (x *IsLogined) GetToken() *Token {
	if x != nil {
		return x.Token
	}
	return nil
}

var File_proxy_proto protoreflect.FileDescriptor

var file_proxy_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x6d,
	0x61, 0x69, 0x6e, 0x22, 0x34, 0x0a, 0x0e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x4c, 0x61, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x03, 0x4c, 0x61, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x4c, 0x6e, 0x67, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x4c, 0x6e, 0x67, 0x22, 0x9c, 0x01, 0x0a, 0x07, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x48, 0x6f, 0x75, 0x73, 0x65, 0x5f, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x48, 0x6f, 0x75,
	0x73, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x52, 0x6f, 0x61, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x52, 0x6f, 0x61, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x53, 0x75, 0x62, 0x75, 0x72, 0x62, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x53, 0x75,
	0x62, 0x75, 0x72, 0x62, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x43, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x22, 0x3a, 0x0a, 0x0f, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x63, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x27, 0x0a, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6d,
	0x61, 0x69, 0x6e, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x22, 0x2c, 0x0a, 0x06, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x10,
	0x0a, 0x03, 0x4c, 0x61, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4c, 0x61, 0x74,
	0x12, 0x10, 0x0a, 0x03, 0x4c, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4c,
	0x6f, 0x6e, 0x22, 0x31, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x73, 0x12,
	0x24, 0x0a, 0x06, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x73, 0x52, 0x06, 0x63,
	0x6f, 0x6f, 0x72, 0x64, 0x73, 0x22, 0x14, 0x0a, 0x02, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x49, 0x64, 0x22, 0x26, 0x0a, 0x08, 0x55,
	0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0x23, 0x0a, 0x07, 0x49, 0x73, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x69, 0x73, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x69, 0x73, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x22, 0x1d, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x38, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12,
	0x14, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x22, 0x32, 0x0a, 0x0c, 0x49, 0x73, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x65,
	0x64, 0x12, 0x22, 0x0a, 0x0c, 0x69, 0x73, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x65,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x69, 0x73, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x65, 0x64, 0x22, 0x4c, 0x0a, 0x09, 0x49, 0x73, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x65, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x73, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x65, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x65, 0x64,
	0x12, 0x21, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0b, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x32, 0x7e, 0x0a, 0x0a, 0x47, 0x65, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x3d, 0x0a, 0x0c, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x41, 0x6e, 0x73, 0x77, 0x65,
	0x72, 0x12, 0x14, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x1a, 0x15, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x63, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x00,
	0x12, 0x31, 0x0a, 0x0d, 0x47, 0x65, 0x6f, 0x63, 0x6f, 0x64, 0x65, 0x41, 0x6e, 0x73, 0x77, 0x65,
	0x72, 0x12, 0x0d, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x1a, 0x0f, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6f, 0x72, 0x64,
	0x73, 0x22, 0x00, 0x32, 0x34, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x25, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x44, 0x12, 0x08, 0x2e,
	0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x49, 0x44, 0x1a, 0x0e, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x00, 0x32, 0x97, 0x01, 0x0a, 0x0b, 0x41, 0x75,
	0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x30, 0x0a, 0x0c, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0a, 0x2e, 0x6d, 0x61, 0x69, 0x6e,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x12, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x49, 0x73, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x65, 0x64, 0x22, 0x00, 0x12, 0x2a, 0x0a, 0x09, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0a, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x1a, 0x0f, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x49, 0x73, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x65, 0x64, 0x22, 0x00, 0x12, 0x2a, 0x0a, 0x0a, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x0b, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x1a, 0x0d, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x49, 0x73, 0x56, 0x61, 0x6c, 0x69,
	0x64, 0x22, 0x00, 0x42, 0x12, 0x5a, 0x10, 0x2f, 0x67, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proxy_proto_rawDescOnce sync.Once
	file_proxy_proto_rawDescData = file_proxy_proto_rawDesc
)

func file_proxy_proto_rawDescGZIP() []byte {
	file_proxy_proto_rawDescOnce.Do(func() {
		file_proxy_proto_rawDescData = protoimpl.X.CompressGZIP(file_proxy_proto_rawDescData)
	})
	return file_proxy_proto_rawDescData
}

var file_proxy_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_proxy_proto_goTypes = []any{
	(*RequestAddress)(nil),  // 0: main.RequestAddress
	(*Address)(nil),         // 1: main.Address
	(*ResponceAddress)(nil), // 2: main.ResponceAddress
	(*Coords)(nil),          // 3: main.Coords
	(*GetCoords)(nil),       // 4: main.GetCoords
	(*ID)(nil),              // 5: main.ID
	(*UserName)(nil),        // 6: main.UserName
	(*IsValid)(nil),         // 7: main.IsValid
	(*Token)(nil),           // 8: main.Token
	(*User)(nil),            // 9: main.User
	(*IsRegistered)(nil),    // 10: main.IsRegistered
	(*IsLogined)(nil),       // 11: main.IsLogined
}
var file_proxy_proto_depIdxs = []int32{
	1,  // 0: main.ResponceAddress.address:type_name -> main.Address
	3,  // 1: main.GetCoords.coords:type_name -> main.Coords
	8,  // 2: main.IsLogined.token:type_name -> main.Token
	0,  // 3: main.GeoService.SearchAnswer:input_type -> main.RequestAddress
	1,  // 4: main.GeoService.GeocodeAnswer:input_type -> main.Address
	5,  // 5: main.UserService.GetByID:input_type -> main.ID
	9,  // 6: main.AuthService.RegisterUser:input_type -> main.User
	9,  // 7: main.AuthService.LoginUser:input_type -> main.User
	8,  // 8: main.AuthService.CheckToken:input_type -> main.Token
	2,  // 9: main.GeoService.SearchAnswer:output_type -> main.ResponceAddress
	4,  // 10: main.GeoService.GeocodeAnswer:output_type -> main.GetCoords
	6,  // 11: main.UserService.GetByID:output_type -> main.UserName
	10, // 12: main.AuthService.RegisterUser:output_type -> main.IsRegistered
	11, // 13: main.AuthService.LoginUser:output_type -> main.IsLogined
	7,  // 14: main.AuthService.CheckToken:output_type -> main.IsValid
	9,  // [9:15] is the sub-list for method output_type
	3,  // [3:9] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_proxy_proto_init() }
func file_proxy_proto_init() {
	if File_proxy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proxy_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*RequestAddress); i {
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
		file_proxy_proto_msgTypes[1].Exporter = func(v any, i int) any {
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
		file_proxy_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*ResponceAddress); i {
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
		file_proxy_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*Coords); i {
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
		file_proxy_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*GetCoords); i {
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
		file_proxy_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*ID); i {
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
		file_proxy_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*UserName); i {
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
		file_proxy_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*IsValid); i {
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
		file_proxy_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*Token); i {
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
		file_proxy_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*User); i {
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
		file_proxy_proto_msgTypes[10].Exporter = func(v any, i int) any {
			switch v := v.(*IsRegistered); i {
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
		file_proxy_proto_msgTypes[11].Exporter = func(v any, i int) any {
			switch v := v.(*IsLogined); i {
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
			RawDescriptor: file_proxy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   3,
		},
		GoTypes:           file_proxy_proto_goTypes,
		DependencyIndexes: file_proxy_proto_depIdxs,
		MessageInfos:      file_proxy_proto_msgTypes,
	}.Build()
	File_proxy_proto = out.File
	file_proxy_proto_rawDesc = nil
	file_proxy_proto_goTypes = nil
	file_proxy_proto_depIdxs = nil
}
