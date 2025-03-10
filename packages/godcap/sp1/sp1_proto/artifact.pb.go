// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v4.25.6
// source: artifact.proto

package sp1_proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type ArtifactType int32

const (
	ArtifactType_UnspecifiedArtifactType ArtifactType = 0
	// / A program artifact.
	ArtifactType_Program ArtifactType = 1
	// / A stdin artifact.
	ArtifactType_Stdin ArtifactType = 2
	// / A proof artifact.
	ArtifactType_Proof ArtifactType = 3
)

// Enum value maps for ArtifactType.
var (
	ArtifactType_name = map[int32]string{
		0: "UnspecifiedArtifactType",
		1: "Program",
		2: "Stdin",
		3: "Proof",
	}
	ArtifactType_value = map[string]int32{
		"UnspecifiedArtifactType": 0,
		"Program":                 1,
		"Stdin":                   2,
		"Proof":                   3,
	}
)

func (x ArtifactType) Enum() *ArtifactType {
	p := new(ArtifactType)
	*p = x
	return p
}

func (x ArtifactType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ArtifactType) Descriptor() protoreflect.EnumDescriptor {
	return file_artifact_proto_enumTypes[0].Descriptor()
}

func (ArtifactType) Type() protoreflect.EnumType {
	return &file_artifact_proto_enumTypes[0]
}

func (x ArtifactType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ArtifactType.Descriptor instead.
func (ArtifactType) EnumDescriptor() ([]byte, []int) {
	return file_artifact_proto_rawDescGZIP(), []int{0}
}

type CreateArtifactRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Signature     []byte                 `protobuf:"bytes,1,opt,name=signature,proto3" json:"signature,omitempty"`
	ArtifactType  ArtifactType           `protobuf:"varint,2,opt,name=artifact_type,json=artifactType,proto3,enum=artifact.ArtifactType" json:"artifact_type,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateArtifactRequest) Reset() {
	*x = CreateArtifactRequest{}
	mi := &file_artifact_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateArtifactRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateArtifactRequest) ProtoMessage() {}

func (x *CreateArtifactRequest) ProtoReflect() protoreflect.Message {
	mi := &file_artifact_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateArtifactRequest.ProtoReflect.Descriptor instead.
func (*CreateArtifactRequest) Descriptor() ([]byte, []int) {
	return file_artifact_proto_rawDescGZIP(), []int{0}
}

func (x *CreateArtifactRequest) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

func (x *CreateArtifactRequest) GetArtifactType() ArtifactType {
	if x != nil {
		return x.ArtifactType
	}
	return ArtifactType_UnspecifiedArtifactType
}

type CreateArtifactResponse struct {
	state                protoimpl.MessageState `protogen:"open.v1"`
	ArtifactUri          string                 `protobuf:"bytes,1,opt,name=artifact_uri,json=artifactUri,proto3" json:"artifact_uri,omitempty"`
	ArtifactPresignedUrl string                 `protobuf:"bytes,2,opt,name=artifact_presigned_url,json=artifactPresignedUrl,proto3" json:"artifact_presigned_url,omitempty"`
	unknownFields        protoimpl.UnknownFields
	sizeCache            protoimpl.SizeCache
}

func (x *CreateArtifactResponse) Reset() {
	*x = CreateArtifactResponse{}
	mi := &file_artifact_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateArtifactResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateArtifactResponse) ProtoMessage() {}

func (x *CreateArtifactResponse) ProtoReflect() protoreflect.Message {
	mi := &file_artifact_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateArtifactResponse.ProtoReflect.Descriptor instead.
func (*CreateArtifactResponse) Descriptor() ([]byte, []int) {
	return file_artifact_proto_rawDescGZIP(), []int{1}
}

func (x *CreateArtifactResponse) GetArtifactUri() string {
	if x != nil {
		return x.ArtifactUri
	}
	return ""
}

func (x *CreateArtifactResponse) GetArtifactPresignedUrl() string {
	if x != nil {
		return x.ArtifactPresignedUrl
	}
	return ""
}

var File_artifact_proto protoreflect.FileDescriptor

var file_artifact_proto_rawDesc = string([]byte{
	0x0a, 0x0e, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x22, 0x72, 0x0a, 0x15, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x12, 0x3b, 0x0a, 0x0d, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x66,
	0x61, 0x63, 0x74, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x0c, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x22, 0x71,
	0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x72, 0x74, 0x69,
	0x66, 0x61, 0x63, 0x74, 0x5f, 0x75, 0x72, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x55, 0x72, 0x69, 0x12, 0x34, 0x0a, 0x16, 0x61,
	0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x5f, 0x70, 0x72, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x65,
	0x64, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x61, 0x72, 0x74,
	0x69, 0x66, 0x61, 0x63, 0x74, 0x50, 0x72, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x55, 0x72,
	0x6c, 0x2a, 0x4e, 0x0a, 0x0c, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x1b, 0x0a, 0x17, 0x55, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x65, 0x64,
	0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x10, 0x00, 0x12, 0x0b,
	0x0a, 0x07, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x53,
	0x74, 0x64, 0x69, 0x6e, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x10,
	0x03, 0x32, 0x64, 0x0a, 0x0d, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x53, 0x74, 0x6f,
	0x72, 0x65, 0x12, 0x53, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69,
	0x66, 0x61, 0x63, 0x74, 0x12, 0x1f, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x2f, 0x73, 0x70, 0x31,
	0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_artifact_proto_rawDescOnce sync.Once
	file_artifact_proto_rawDescData []byte
)

func file_artifact_proto_rawDescGZIP() []byte {
	file_artifact_proto_rawDescOnce.Do(func() {
		file_artifact_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_artifact_proto_rawDesc), len(file_artifact_proto_rawDesc)))
	})
	return file_artifact_proto_rawDescData
}

var file_artifact_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_artifact_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_artifact_proto_goTypes = []any{
	(ArtifactType)(0),              // 0: artifact.ArtifactType
	(*CreateArtifactRequest)(nil),  // 1: artifact.CreateArtifactRequest
	(*CreateArtifactResponse)(nil), // 2: artifact.CreateArtifactResponse
}
var file_artifact_proto_depIdxs = []int32{
	0, // 0: artifact.CreateArtifactRequest.artifact_type:type_name -> artifact.ArtifactType
	1, // 1: artifact.ArtifactStore.CreateArtifact:input_type -> artifact.CreateArtifactRequest
	2, // 2: artifact.ArtifactStore.CreateArtifact:output_type -> artifact.CreateArtifactResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_artifact_proto_init() }
func file_artifact_proto_init() {
	if File_artifact_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_artifact_proto_rawDesc), len(file_artifact_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_artifact_proto_goTypes,
		DependencyIndexes: file_artifact_proto_depIdxs,
		EnumInfos:         file_artifact_proto_enumTypes,
		MessageInfos:      file_artifact_proto_msgTypes,
	}.Build()
	File_artifact_proto = out.File
	file_artifact_proto_goTypes = nil
	file_artifact_proto_depIdxs = nil
}
