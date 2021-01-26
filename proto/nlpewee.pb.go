// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.14.0
// source: nlpewee.proto

package proto

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Language int32

const (
	Language_ENGLISH Language = 0
)

// Enum value maps for Language.
var (
	Language_name = map[int32]string{
		0: "ENGLISH",
	}
	Language_value = map[string]int32{
		"ENGLISH": 0,
	}
)

func (x Language) Enum() *Language {
	p := new(Language)
	*p = x
	return p
}

func (x Language) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Language) Descriptor() protoreflect.EnumDescriptor {
	return file_nlpewee_proto_enumTypes[0].Descriptor()
}

func (Language) Type() protoreflect.EnumType {
	return &file_nlpewee_proto_enumTypes[0]
}

func (x Language) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Language.Descriptor instead.
func (Language) EnumDescriptor() ([]byte, []int) {
	return file_nlpewee_proto_rawDescGZIP(), []int{0}
}

type TokenizeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text     string   `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	Language Language `protobuf:"varint,2,opt,name=language,proto3,enum=nlpewee.Language" json:"language,omitempty"`
}

func (x *TokenizeRequest) Reset() {
	*x = TokenizeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nlpewee_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenizeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenizeRequest) ProtoMessage() {}

func (x *TokenizeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nlpewee_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenizeRequest.ProtoReflect.Descriptor instead.
func (*TokenizeRequest) Descriptor() ([]byte, []int) {
	return file_nlpewee_proto_rawDescGZIP(), []int{0}
}

func (x *TokenizeRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *TokenizeRequest) GetLanguage() Language {
	if x != nil {
		return x.Language
	}
	return Language_ENGLISH
}

type TokenizeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sentences []*Sentence `protobuf:"bytes,1,rep,name=sentences,proto3" json:"sentences,omitempty"`
}

func (x *TokenizeResponse) Reset() {
	*x = TokenizeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nlpewee_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenizeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenizeResponse) ProtoMessage() {}

func (x *TokenizeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_nlpewee_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenizeResponse.ProtoReflect.Descriptor instead.
func (*TokenizeResponse) Descriptor() ([]byte, []int) {
	return file_nlpewee_proto_rawDescGZIP(), []int{1}
}

func (x *TokenizeResponse) GetSentences() []*Sentence {
	if x != nil {
		return x.Sentences
	}
	return nil
}

type Sentence struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tokens   []*Token  `protobuf:"bytes,1,rep,name=tokens,proto3" json:"tokens,omitempty"`
	Entities []*Entity `protobuf:"bytes,2,rep,name=entities,proto3" json:"entities,omitempty"`
}

func (x *Sentence) Reset() {
	*x = Sentence{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nlpewee_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sentence) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sentence) ProtoMessage() {}

func (x *Sentence) ProtoReflect() protoreflect.Message {
	mi := &file_nlpewee_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sentence.ProtoReflect.Descriptor instead.
func (*Sentence) Descriptor() ([]byte, []int) {
	return file_nlpewee_proto_rawDescGZIP(), []int{2}
}

func (x *Sentence) GetTokens() []*Token {
	if x != nil {
		return x.Tokens
	}
	return nil
}

func (x *Sentence) GetEntities() []*Entity {
	if x != nil {
		return x.Entities
	}
	return nil
}

type Token struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text  string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	Stem  string `protobuf:"bytes,2,opt,name=stem,proto3" json:"stem,omitempty"`
	Tag   string `protobuf:"bytes,3,opt,name=tag,proto3" json:"tag,omitempty"`
	Label string `protobuf:"bytes,4,opt,name=label,proto3" json:"label,omitempty"`
}

func (x *Token) Reset() {
	*x = Token{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nlpewee_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Token) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Token) ProtoMessage() {}

func (x *Token) ProtoReflect() protoreflect.Message {
	mi := &file_nlpewee_proto_msgTypes[3]
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
	return file_nlpewee_proto_rawDescGZIP(), []int{3}
}

func (x *Token) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Token) GetStem() string {
	if x != nil {
		return x.Stem
	}
	return ""
}

func (x *Token) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

func (x *Token) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

type Entity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text  string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	Label string `protobuf:"bytes,2,opt,name=label,proto3" json:"label,omitempty"`
}

func (x *Entity) Reset() {
	*x = Entity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nlpewee_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Entity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Entity) ProtoMessage() {}

func (x *Entity) ProtoReflect() protoreflect.Message {
	mi := &file_nlpewee_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Entity.ProtoReflect.Descriptor instead.
func (*Entity) Descriptor() ([]byte, []int) {
	return file_nlpewee_proto_rawDescGZIP(), []int{4}
}

func (x *Entity) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Entity) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

var File_nlpewee_proto protoreflect.FileDescriptor

var file_nlpewee_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6e, 0x6c, 0x70, 0x65, 0x77, 0x65, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x6e, 0x6c, 0x70, 0x65, 0x77, 0x65, 0x65, 0x22, 0x54, 0x0a, 0x0f, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12,
	0x2d, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x11, 0x2e, 0x6e, 0x6c, 0x70, 0x65, 0x77, 0x65, 0x65, 0x2e, 0x4c, 0x61, 0x6e, 0x67,
	0x75, 0x61, 0x67, 0x65, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x22, 0x43,
	0x0a, 0x10, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x2f, 0x0a, 0x09, 0x73, 0x65, 0x6e, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6e, 0x6c, 0x70, 0x65, 0x77, 0x65, 0x65, 0x2e,
	0x53, 0x65, 0x6e, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x09, 0x73, 0x65, 0x6e, 0x74, 0x65, 0x6e,
	0x63, 0x65, 0x73, 0x22, 0x5f, 0x0a, 0x08, 0x53, 0x65, 0x6e, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x12,
	0x26, 0x0a, 0x06, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0e, 0x2e, 0x6e, 0x6c, 0x70, 0x65, 0x77, 0x65, 0x65, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52,
	0x06, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x12, 0x2b, 0x0a, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6e, 0x6c, 0x70, 0x65,
	0x77, 0x65, 0x65, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x08, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x69, 0x65, 0x73, 0x22, 0x57, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x74, 0x65, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x73, 0x74, 0x65, 0x6d, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x74, 0x61, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x22, 0x32, 0x0a,
	0x06, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65,
	0x6c, 0x2a, 0x17, 0x0a, 0x08, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x0b, 0x0a,
	0x07, 0x45, 0x4e, 0x47, 0x4c, 0x49, 0x53, 0x48, 0x10, 0x00, 0x32, 0x90, 0x01, 0x0a, 0x07, 0x4e,
	0x4c, 0x50, 0x65, 0x77, 0x65, 0x65, 0x12, 0x3f, 0x0a, 0x08, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x69,
	0x7a, 0x65, 0x12, 0x18, 0x2e, 0x6e, 0x6c, 0x70, 0x65, 0x77, 0x65, 0x65, 0x2e, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x6e,
	0x6c, 0x70, 0x65, 0x77, 0x65, 0x65, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x69, 0x7a, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x0d, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x69, 0x7a, 0x65, 0x43, 0x6c, 0x65, 0x61, 0x6e, 0x12, 0x18, 0x2e, 0x6e, 0x6c, 0x70, 0x65, 0x77,
	0x65, 0x65, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x19, 0x2e, 0x6e, 0x6c, 0x70, 0x65, 0x77, 0x65, 0x65, 0x2e, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x26, 0x5a,
	0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x73, 0x68, 0x69, 0x72, 0x65, 0x2f, 0x6e, 0x6c, 0x70, 0x65, 0x77, 0x65, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_nlpewee_proto_rawDescOnce sync.Once
	file_nlpewee_proto_rawDescData = file_nlpewee_proto_rawDesc
)

func file_nlpewee_proto_rawDescGZIP() []byte {
	file_nlpewee_proto_rawDescOnce.Do(func() {
		file_nlpewee_proto_rawDescData = protoimpl.X.CompressGZIP(file_nlpewee_proto_rawDescData)
	})
	return file_nlpewee_proto_rawDescData
}

var file_nlpewee_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_nlpewee_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_nlpewee_proto_goTypes = []interface{}{
	(Language)(0),            // 0: nlpewee.Language
	(*TokenizeRequest)(nil),  // 1: nlpewee.TokenizeRequest
	(*TokenizeResponse)(nil), // 2: nlpewee.TokenizeResponse
	(*Sentence)(nil),         // 3: nlpewee.Sentence
	(*Token)(nil),            // 4: nlpewee.Token
	(*Entity)(nil),           // 5: nlpewee.Entity
}
var file_nlpewee_proto_depIdxs = []int32{
	0, // 0: nlpewee.TokenizeRequest.language:type_name -> nlpewee.Language
	3, // 1: nlpewee.TokenizeResponse.sentences:type_name -> nlpewee.Sentence
	4, // 2: nlpewee.Sentence.tokens:type_name -> nlpewee.Token
	5, // 3: nlpewee.Sentence.entities:type_name -> nlpewee.Entity
	1, // 4: nlpewee.NLPewee.Tokenize:input_type -> nlpewee.TokenizeRequest
	1, // 5: nlpewee.NLPewee.TokenizeClean:input_type -> nlpewee.TokenizeRequest
	2, // 6: nlpewee.NLPewee.Tokenize:output_type -> nlpewee.TokenizeResponse
	2, // 7: nlpewee.NLPewee.TokenizeClean:output_type -> nlpewee.TokenizeResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_nlpewee_proto_init() }
func file_nlpewee_proto_init() {
	if File_nlpewee_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_nlpewee_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenizeRequest); i {
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
		file_nlpewee_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenizeResponse); i {
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
		file_nlpewee_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sentence); i {
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
		file_nlpewee_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_nlpewee_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Entity); i {
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
			RawDescriptor: file_nlpewee_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_nlpewee_proto_goTypes,
		DependencyIndexes: file_nlpewee_proto_depIdxs,
		EnumInfos:         file_nlpewee_proto_enumTypes,
		MessageInfos:      file_nlpewee_proto_msgTypes,
	}.Build()
	File_nlpewee_proto = out.File
	file_nlpewee_proto_rawDesc = nil
	file_nlpewee_proto_goTypes = nil
	file_nlpewee_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// NLPeweeClient is the client API for NLPewee service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NLPeweeClient interface {
	Tokenize(ctx context.Context, in *TokenizeRequest, opts ...grpc.CallOption) (*TokenizeResponse, error)
	TokenizeClean(ctx context.Context, in *TokenizeRequest, opts ...grpc.CallOption) (*TokenizeResponse, error)
}

type nLPeweeClient struct {
	cc grpc.ClientConnInterface
}

func NewNLPeweeClient(cc grpc.ClientConnInterface) NLPeweeClient {
	return &nLPeweeClient{cc}
}

func (c *nLPeweeClient) Tokenize(ctx context.Context, in *TokenizeRequest, opts ...grpc.CallOption) (*TokenizeResponse, error) {
	out := new(TokenizeResponse)
	err := c.cc.Invoke(ctx, "/nlpewee.NLPewee/Tokenize", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nLPeweeClient) TokenizeClean(ctx context.Context, in *TokenizeRequest, opts ...grpc.CallOption) (*TokenizeResponse, error) {
	out := new(TokenizeResponse)
	err := c.cc.Invoke(ctx, "/nlpewee.NLPewee/TokenizeClean", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NLPeweeServer is the server API for NLPewee service.
type NLPeweeServer interface {
	Tokenize(context.Context, *TokenizeRequest) (*TokenizeResponse, error)
	TokenizeClean(context.Context, *TokenizeRequest) (*TokenizeResponse, error)
}

// UnimplementedNLPeweeServer can be embedded to have forward compatible implementations.
type UnimplementedNLPeweeServer struct {
}

func (*UnimplementedNLPeweeServer) Tokenize(context.Context, *TokenizeRequest) (*TokenizeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Tokenize not implemented")
}
func (*UnimplementedNLPeweeServer) TokenizeClean(context.Context, *TokenizeRequest) (*TokenizeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TokenizeClean not implemented")
}

func RegisterNLPeweeServer(s *grpc.Server, srv NLPeweeServer) {
	s.RegisterService(&_NLPewee_serviceDesc, srv)
}

func _NLPewee_Tokenize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TokenizeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NLPeweeServer).Tokenize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nlpewee.NLPewee/Tokenize",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NLPeweeServer).Tokenize(ctx, req.(*TokenizeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NLPewee_TokenizeClean_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TokenizeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NLPeweeServer).TokenizeClean(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nlpewee.NLPewee/TokenizeClean",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NLPeweeServer).TokenizeClean(ctx, req.(*TokenizeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NLPewee_serviceDesc = grpc.ServiceDesc{
	ServiceName: "nlpewee.NLPewee",
	HandlerType: (*NLPeweeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Tokenize",
			Handler:    _NLPewee_Tokenize_Handler,
		},
		{
			MethodName: "TokenizeClean",
			Handler:    _NLPewee_TokenizeClean_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "nlpewee.proto",
}