// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.1
// source: service_grpc.proto

package pb

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

type CreatePersonRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Age  int64  `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
	Bio  string `protobuf:"bytes,3,opt,name=bio,proto3" json:"bio,omitempty"`
}

func (x *CreatePersonRequest) Reset() {
	*x = CreatePersonRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_grpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePersonRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePersonRequest) ProtoMessage() {}

func (x *CreatePersonRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_grpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePersonRequest.ProtoReflect.Descriptor instead.
func (*CreatePersonRequest) Descriptor() ([]byte, []int) {
	return file_service_grpc_proto_rawDescGZIP(), []int{0}
}

func (x *CreatePersonRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreatePersonRequest) GetAge() int64 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *CreatePersonRequest) GetBio() string {
	if x != nil {
		return x.Bio
	}
	return ""
}

type CreatePersonResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Age  int64  `protobuf:"varint,3,opt,name=age,proto3" json:"age,omitempty"`
	Bio  string `protobuf:"bytes,4,opt,name=bio,proto3" json:"bio,omitempty"`
}

func (x *CreatePersonResponse) Reset() {
	*x = CreatePersonResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_grpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePersonResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePersonResponse) ProtoMessage() {}

func (x *CreatePersonResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_grpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePersonResponse.ProtoReflect.Descriptor instead.
func (*CreatePersonResponse) Descriptor() ([]byte, []int) {
	return file_service_grpc_proto_rawDescGZIP(), []int{1}
}

func (x *CreatePersonResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreatePersonResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreatePersonResponse) GetAge() int64 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *CreatePersonResponse) GetBio() string {
	if x != nil {
		return x.Bio
	}
	return ""
}

type GetListPersonReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page    int64 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PerPage int64 `protobuf:"varint,2,opt,name=per_page,json=perPage,proto3" json:"per_page,omitempty"`
}

func (x *GetListPersonReq) Reset() {
	*x = GetListPersonReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_grpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListPersonReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListPersonReq) ProtoMessage() {}

func (x *GetListPersonReq) ProtoReflect() protoreflect.Message {
	mi := &file_service_grpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListPersonReq.ProtoReflect.Descriptor instead.
func (*GetListPersonReq) Descriptor() ([]byte, []int) {
	return file_service_grpc_proto_rawDescGZIP(), []int{2}
}

func (x *GetListPersonReq) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetListPersonReq) GetPerPage() int64 {
	if x != nil {
		return x.PerPage
	}
	return 0
}

type GetListPersonRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	People []*Person `protobuf:"bytes,1,rep,name=people,proto3" json:"people,omitempty"`
}

func (x *GetListPersonRes) Reset() {
	*x = GetListPersonRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_grpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListPersonRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListPersonRes) ProtoMessage() {}

func (x *GetListPersonRes) ProtoReflect() protoreflect.Message {
	mi := &file_service_grpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListPersonRes.ProtoReflect.Descriptor instead.
func (*GetListPersonRes) Descriptor() ([]byte, []int) {
	return file_service_grpc_proto_rawDescGZIP(), []int{3}
}

func (x *GetListPersonRes) GetPeople() []*Person {
	if x != nil {
		return x.People
	}
	return nil
}

type Person struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Age  int64  `protobuf:"varint,3,opt,name=age,proto3" json:"age,omitempty"`
	Bio  string `protobuf:"bytes,4,opt,name=bio,proto3" json:"bio,omitempty"`
}

func (x *Person) Reset() {
	*x = Person{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_grpc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Person) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Person) ProtoMessage() {}

func (x *Person) ProtoReflect() protoreflect.Message {
	mi := &file_service_grpc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Person.ProtoReflect.Descriptor instead.
func (*Person) Descriptor() ([]byte, []int) {
	return file_service_grpc_proto_rawDescGZIP(), []int{4}
}

func (x *Person) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Person) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Person) GetAge() int64 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *Person) GetBio() string {
	if x != nil {
		return x.Bio
	}
	return ""
}

var File_service_grpc_proto protoreflect.FileDescriptor

var file_service_grpc_proto_rawDesc = []byte{
	0x0a, 0x12, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x4d, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x03, 0x61, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x62, 0x69, 0x6f, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x62, 0x69, 0x6f, 0x22, 0x5e, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x03, 0x61, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x62, 0x69, 0x6f, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x62, 0x69, 0x6f, 0x22, 0x41, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x4c, 0x69,
	0x73, 0x74, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x70,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12,
	0x19, 0x0a, 0x08, 0x70, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x07, 0x70, 0x65, 0x72, 0x50, 0x61, 0x67, 0x65, 0x22, 0x36, 0x0a, 0x10, 0x47, 0x65,
	0x74, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x12, 0x22,
	0x0a, 0x06, 0x70, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a,
	0x2e, 0x70, 0x62, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x06, 0x70, 0x65, 0x6f, 0x70,
	0x6c, 0x65, 0x22, 0x50, 0x0a, 0x06, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x61,
	0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x62, 0x69, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x62, 0x69, 0x6f, 0x32, 0x8e, 0x01, 0x0a, 0x08, 0x54, 0x65, 0x73, 0x74, 0x47, 0x72, 0x70,
	0x63, 0x12, 0x43, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x73, 0x6f,
	0x6e, 0x12, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x62, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73,
	0x74, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74,
	0x4c, 0x69, 0x73, 0x74, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e,
	0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x22, 0x00, 0x42, 0x0e, 0x5a, 0x0c, 0x74, 0x65, 0x73, 0x74, 0x6d, 0x6f, 0x6e,
	0x67, 0x6f, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_grpc_proto_rawDescOnce sync.Once
	file_service_grpc_proto_rawDescData = file_service_grpc_proto_rawDesc
)

func file_service_grpc_proto_rawDescGZIP() []byte {
	file_service_grpc_proto_rawDescOnce.Do(func() {
		file_service_grpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_grpc_proto_rawDescData)
	})
	return file_service_grpc_proto_rawDescData
}

var file_service_grpc_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_service_grpc_proto_goTypes = []interface{}{
	(*CreatePersonRequest)(nil),  // 0: pb.CreatePersonRequest
	(*CreatePersonResponse)(nil), // 1: pb.CreatePersonResponse
	(*GetListPersonReq)(nil),     // 2: pb.GetListPersonReq
	(*GetListPersonRes)(nil),     // 3: pb.GetListPersonRes
	(*Person)(nil),               // 4: pb.Person
}
var file_service_grpc_proto_depIdxs = []int32{
	4, // 0: pb.GetListPersonRes.people:type_name -> pb.Person
	0, // 1: pb.TestGrpc.CreatePerson:input_type -> pb.CreatePersonRequest
	2, // 2: pb.TestGrpc.GetListPerson:input_type -> pb.GetListPersonReq
	1, // 3: pb.TestGrpc.CreatePerson:output_type -> pb.CreatePersonResponse
	3, // 4: pb.TestGrpc.GetListPerson:output_type -> pb.GetListPersonRes
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_service_grpc_proto_init() }
func file_service_grpc_proto_init() {
	if File_service_grpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_service_grpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePersonRequest); i {
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
		file_service_grpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePersonResponse); i {
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
		file_service_grpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListPersonReq); i {
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
		file_service_grpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListPersonRes); i {
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
		file_service_grpc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Person); i {
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
			RawDescriptor: file_service_grpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_grpc_proto_goTypes,
		DependencyIndexes: file_service_grpc_proto_depIdxs,
		MessageInfos:      file_service_grpc_proto_msgTypes,
	}.Build()
	File_service_grpc_proto = out.File
	file_service_grpc_proto_rawDesc = nil
	file_service_grpc_proto_goTypes = nil
	file_service_grpc_proto_depIdxs = nil
}
