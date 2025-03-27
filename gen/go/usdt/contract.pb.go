// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        v3.21.12
// source: usdt/contract.proto

package cryptoservicev1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ResponsePong struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Pong          string                 `protobuf:"bytes,1,opt,name=pong,proto3" json:"pong,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ResponsePong) Reset() {
	*x = ResponsePong{}
	mi := &file_usdt_contract_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ResponsePong) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponsePong) ProtoMessage() {}

func (x *ResponsePong) ProtoReflect() protoreflect.Message {
	mi := &file_usdt_contract_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponsePong.ProtoReflect.Descriptor instead.
func (*ResponsePong) Descriptor() ([]byte, []int) {
	return file_usdt_contract_proto_rawDescGZIP(), []int{0}
}

func (x *ResponsePong) GetPong() string {
	if x != nil {
		return x.Pong
	}
	return ""
}

type Ask struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Price         float64                `protobuf:"fixed64,1,opt,name=price,proto3" json:"price,omitempty"`
	Volume        float64                `protobuf:"fixed64,2,opt,name=volume,proto3" json:"volume,omitempty"`
	Timestamp     int64                  `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Ask) Reset() {
	*x = Ask{}
	mi := &file_usdt_contract_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Ask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ask) ProtoMessage() {}

func (x *Ask) ProtoReflect() protoreflect.Message {
	mi := &file_usdt_contract_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ask.ProtoReflect.Descriptor instead.
func (*Ask) Descriptor() ([]byte, []int) {
	return file_usdt_contract_proto_rawDescGZIP(), []int{1}
}

func (x *Ask) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Ask) GetVolume() float64 {
	if x != nil {
		return x.Volume
	}
	return 0
}

func (x *Ask) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type Bid struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Price         float64                `protobuf:"fixed64,1,opt,name=price,proto3" json:"price,omitempty"`
	Volume        float64                `protobuf:"fixed64,2,opt,name=volume,proto3" json:"volume,omitempty"`
	Timestamp     int64                  `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Bid) Reset() {
	*x = Bid{}
	mi := &file_usdt_contract_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Bid) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bid) ProtoMessage() {}

func (x *Bid) ProtoReflect() protoreflect.Message {
	mi := &file_usdt_contract_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bid.ProtoReflect.Descriptor instead.
func (*Bid) Descriptor() ([]byte, []int) {
	return file_usdt_contract_proto_rawDescGZIP(), []int{2}
}

func (x *Bid) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Bid) GetVolume() float64 {
	if x != nil {
		return x.Volume
	}
	return 0
}

func (x *Bid) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type Response struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Timestamp     int64                  `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Asks          []*Ask                 `protobuf:"bytes,2,rep,name=asks,proto3" json:"asks,omitempty"`
	Bids          []*Bid                 `protobuf:"bytes,3,rep,name=bids,proto3" json:"bids,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Response) Reset() {
	*x = Response{}
	mi := &file_usdt_contract_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_usdt_contract_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_usdt_contract_proto_rawDescGZIP(), []int{3}
}

func (x *Response) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *Response) GetAsks() []*Ask {
	if x != nil {
		return x.Asks
	}
	return nil
}

func (x *Response) GetBids() []*Bid {
	if x != nil {
		return x.Bids
	}
	return nil
}

var File_usdt_contract_proto protoreflect.FileDescriptor

var file_usdt_contract_proto_rawDesc = string([]byte{
	0x0a, 0x13, 0x75, 0x73, 0x64, 0x74, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x22, 0x0a, 0x0c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x50, 0x6f, 0x6e,
	0x67, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x70, 0x6f, 0x6e, 0x67, 0x22, 0x51, 0x0a, 0x03, 0x41, 0x73, 0x6b, 0x12, 0x14, 0x0a, 0x05,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x51, 0x0a, 0x03, 0x42, 0x69, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x78, 0x0a, 0x08, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x26, 0x0a, 0x04, 0x61, 0x73, 0x6b, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x41, 0x73, 0x6b, 0x52, 0x04, 0x61, 0x73, 0x6b, 0x73, 0x12, 0x26, 0x0a,
	0x04, 0x62, 0x69, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x72,
	0x79, 0x70, 0x74, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x42, 0x69, 0x64, 0x52,
	0x04, 0x62, 0x69, 0x64, 0x73, 0x32, 0x8a, 0x01, 0x0a, 0x0e, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f,
	0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x3b, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x52,
	0x61, 0x74, 0x65, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x17, 0x2e, 0x63,
	0x72, 0x79, 0x70, 0x74, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1b, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x50, 0x6f,
	0x6e, 0x67, 0x42, 0x22, 0x5a, 0x20, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x3b, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_usdt_contract_proto_rawDescOnce sync.Once
	file_usdt_contract_proto_rawDescData []byte
)

func file_usdt_contract_proto_rawDescGZIP() []byte {
	file_usdt_contract_proto_rawDescOnce.Do(func() {
		file_usdt_contract_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_usdt_contract_proto_rawDesc), len(file_usdt_contract_proto_rawDesc)))
	})
	return file_usdt_contract_proto_rawDescData
}

var file_usdt_contract_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_usdt_contract_proto_goTypes = []any{
	(*ResponsePong)(nil),  // 0: cryptoservice.ResponsePong
	(*Ask)(nil),           // 1: cryptoservice.Ask
	(*Bid)(nil),           // 2: cryptoservice.Bid
	(*Response)(nil),      // 3: cryptoservice.Response
	(*emptypb.Empty)(nil), // 4: google.protobuf.Empty
}
var file_usdt_contract_proto_depIdxs = []int32{
	1, // 0: cryptoservice.Response.asks:type_name -> cryptoservice.Ask
	2, // 1: cryptoservice.Response.bids:type_name -> cryptoservice.Bid
	4, // 2: cryptoservice.Cryptoprovider.GetRates:input_type -> google.protobuf.Empty
	4, // 3: cryptoservice.Cryptoprovider.Ping:input_type -> google.protobuf.Empty
	3, // 4: cryptoservice.Cryptoprovider.GetRates:output_type -> cryptoservice.Response
	0, // 5: cryptoservice.Cryptoprovider.Ping:output_type -> cryptoservice.ResponsePong
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_usdt_contract_proto_init() }
func file_usdt_contract_proto_init() {
	if File_usdt_contract_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_usdt_contract_proto_rawDesc), len(file_usdt_contract_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_usdt_contract_proto_goTypes,
		DependencyIndexes: file_usdt_contract_proto_depIdxs,
		MessageInfos:      file_usdt_contract_proto_msgTypes,
	}.Build()
	File_usdt_contract_proto = out.File
	file_usdt_contract_proto_goTypes = nil
	file_usdt_contract_proto_depIdxs = nil
}
