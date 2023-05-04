// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.6
// source: Queries/Block/proto/queries.proto

package proto

import (
	context "context"
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

type BlockRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Lang    string `protobuf:"bytes,2,opt,name=lang,proto3" json:"lang,omitempty"`
	Version *int32 `protobuf:"varint,3,opt,name=version,proto3,oneof" json:"version,omitempty"`
}

func (x *BlockRequest) Reset() {
	*x = BlockRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Queries_Block_proto_queries_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockRequest) ProtoMessage() {}

func (x *BlockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_Queries_Block_proto_queries_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockRequest.ProtoReflect.Descriptor instead.
func (*BlockRequest) Descriptor() ([]byte, []int) {
	return file_Queries_Block_proto_queries_proto_rawDescGZIP(), []int{0}
}

func (x *BlockRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *BlockRequest) GetLang() string {
	if x != nil {
		return x.Lang
	}
	return ""
}

func (x *BlockRequest) GetVersion() int32 {
	if x != nil && x.Version != nil {
		return *x.Version
	}
	return 0
}

type VersionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Lang string `protobuf:"bytes,2,opt,name=lang,proto3" json:"lang,omitempty"`
}

func (x *VersionsRequest) Reset() {
	*x = VersionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Queries_Block_proto_queries_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VersionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VersionsRequest) ProtoMessage() {}

func (x *VersionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_Queries_Block_proto_queries_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VersionsRequest.ProtoReflect.Descriptor instead.
func (*VersionsRequest) Descriptor() ([]byte, []int) {
	return file_Queries_Block_proto_queries_proto_rawDescGZIP(), []int{1}
}

func (x *VersionsRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *VersionsRequest) GetLang() string {
	if x != nil {
		return x.Lang
	}
	return ""
}

type VersionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Versions []*Ver `protobuf:"bytes,1,rep,name=versions,proto3" json:"versions,omitempty"`
}

func (x *VersionResponse) Reset() {
	*x = VersionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Queries_Block_proto_queries_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VersionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VersionResponse) ProtoMessage() {}

func (x *VersionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_Queries_Block_proto_queries_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VersionResponse.ProtoReflect.Descriptor instead.
func (*VersionResponse) Descriptor() ([]byte, []int) {
	return file_Queries_Block_proto_queries_proto_rawDescGZIP(), []int{2}
}

func (x *VersionResponse) GetVersions() []*Ver {
	if x != nil {
		return x.Versions
	}
	return nil
}

type Ver struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version   int32   `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	ChangeLog *string `protobuf:"bytes,2,opt,name=changeLog,proto3,oneof" json:"changeLog,omitempty"`
}

func (x *Ver) Reset() {
	*x = Ver{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Queries_Block_proto_queries_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ver) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ver) ProtoMessage() {}

func (x *Ver) ProtoReflect() protoreflect.Message {
	mi := &file_Queries_Block_proto_queries_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ver.ProtoReflect.Descriptor instead.
func (*Ver) Descriptor() ([]byte, []int) {
	return file_Queries_Block_proto_queries_proto_rawDescGZIP(), []int{3}
}

func (x *Ver) GetVersion() int32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *Ver) GetChangeLog() string {
	if x != nil && x.ChangeLog != nil {
		return *x.ChangeLog
	}
	return ""
}

type LanguageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *LanguageRequest) Reset() {
	*x = LanguageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Queries_Block_proto_queries_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LanguageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LanguageRequest) ProtoMessage() {}

func (x *LanguageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_Queries_Block_proto_queries_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LanguageRequest.ProtoReflect.Descriptor instead.
func (*LanguageRequest) Descriptor() ([]byte, []int) {
	return file_Queries_Block_proto_queries_proto_rawDescGZIP(), []int{4}
}

func (x *LanguageRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type LanguageList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Langs []string `protobuf:"bytes,1,rep,name=langs,proto3" json:"langs,omitempty"`
}

func (x *LanguageList) Reset() {
	*x = LanguageList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Queries_Block_proto_queries_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LanguageList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LanguageList) ProtoMessage() {}

func (x *LanguageList) ProtoReflect() protoreflect.Message {
	mi := &file_Queries_Block_proto_queries_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LanguageList.ProtoReflect.Descriptor instead.
func (*LanguageList) Descriptor() ([]byte, []int) {
	return file_Queries_Block_proto_queries_proto_rawDescGZIP(), []int{5}
}

func (x *LanguageList) GetLangs() []string {
	if x != nil {
		return x.Langs
	}
	return nil
}

type Selectors struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Authors    []string `protobuf:"bytes,1,rep,name=authors,proto3" json:"authors,omitempty"`
	Categories []string `protobuf:"bytes,2,rep,name=categories,proto3" json:"categories,omitempty"`
	Tags       []string `protobuf:"bytes,3,rep,name=tags,proto3" json:"tags,omitempty"`
	Type       string   `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *Selectors) Reset() {
	*x = Selectors{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Queries_Block_proto_queries_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Selectors) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Selectors) ProtoMessage() {}

func (x *Selectors) ProtoReflect() protoreflect.Message {
	mi := &file_Queries_Block_proto_queries_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Selectors.ProtoReflect.Descriptor instead.
func (*Selectors) Descriptor() ([]byte, []int) {
	return file_Queries_Block_proto_queries_proto_rawDescGZIP(), []int{6}
}

func (x *Selectors) GetAuthors() []string {
	if x != nil {
		return x.Authors
	}
	return nil
}

func (x *Selectors) GetCategories() []string {
	if x != nil {
		return x.Categories
	}
	return nil
}

func (x *Selectors) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *Selectors) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type Pagination struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageSize int32 `protobuf:"varint,1,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	PageNum  int32 `protobuf:"varint,2,opt,name=pageNum,proto3" json:"pageNum,omitempty"`
	Total    int32 `protobuf:"varint,3,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *Pagination) Reset() {
	*x = Pagination{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Queries_Block_proto_queries_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pagination) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pagination) ProtoMessage() {}

func (x *Pagination) ProtoReflect() protoreflect.Message {
	mi := &file_Queries_Block_proto_queries_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pagination.ProtoReflect.Descriptor instead.
func (*Pagination) Descriptor() ([]byte, []int) {
	return file_Queries_Block_proto_queries_proto_rawDescGZIP(), []int{7}
}

func (x *Pagination) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *Pagination) GetPageNum() int32 {
	if x != nil {
		return x.PageNum
	}
	return 0
}

func (x *Pagination) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_Queries_Block_proto_queries_proto protoreflect.FileDescriptor

var file_Queries_Block_proto_queries_proto_rawDesc = []byte{
	0x0a, 0x21, 0x51, 0x75, 0x65, 0x72, 0x69, 0x65, 0x73, 0x2f, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x53, 0x68, 0x61, 0x72,
	0x65, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5d, 0x0a, 0x0c, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x61, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6c, 0x61, 0x6e, 0x67, 0x12, 0x1d, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x22, 0x35, 0x0a, 0x0f, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x61, 0x6e, 0x67, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x61, 0x6e, 0x67, 0x22, 0x39, 0x0a, 0x0f, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a,
	0x08, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x65, 0x72, 0x52, 0x08, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x50, 0x0a, 0x03, 0x76, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x09, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x4c, 0x6f, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x09, 0x63, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x4c, 0x6f, 0x67, 0x88, 0x01, 0x01, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x63, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x4c, 0x6f, 0x67, 0x22, 0x21, 0x0a, 0x0f, 0x4c, 0x61, 0x6e, 0x67, 0x75,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x24, 0x0a, 0x0c, 0x4c, 0x61,
	0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61,
	0x6e, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x6e, 0x67, 0x73,
	0x22, 0x6d, 0x0a, 0x09, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x12, 0x18, 0x0a,
	0x07, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07,
	0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22,
	0x58, 0x0a, 0x0a, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x67,
	0x65, 0x4e, 0x75, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x70, 0x61, 0x67, 0x65,
	0x4e, 0x75, 0x6d, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x32, 0xed, 0x02, 0x0a, 0x07, 0x51, 0x75,
	0x65, 0x72, 0x69, 0x65, 0x73, 0x12, 0x2f, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f,
	0x63, 0x6b, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x4d, 0x65, 0x74, 0x61, 0x22, 0x00, 0x12, 0x39,
	0x0a, 0x0d, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x12,
	0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x6c, 0x6f,
	0x63, 0x6b, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x0b, 0x47, 0x65, 0x74,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x0c, 0x47, 0x65,
	0x74, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x73, 0x12, 0x16, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x61, 0x6e, 0x67, 0x75,
	0x61, 0x67, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x42, 0x0f, 0x5a, 0x0d, 0x2e, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_Queries_Block_proto_queries_proto_rawDescOnce sync.Once
	file_Queries_Block_proto_queries_proto_rawDescData = file_Queries_Block_proto_queries_proto_rawDesc
)

func file_Queries_Block_proto_queries_proto_rawDescGZIP() []byte {
	file_Queries_Block_proto_queries_proto_rawDescOnce.Do(func() {
		file_Queries_Block_proto_queries_proto_rawDescData = protoimpl.X.CompressGZIP(file_Queries_Block_proto_queries_proto_rawDescData)
	})
	return file_Queries_Block_proto_queries_proto_rawDescData
}

var file_Queries_Block_proto_queries_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_Queries_Block_proto_queries_proto_goTypes = []interface{}{
	(*BlockRequest)(nil),    // 0: proto.BlockRequest
	(*VersionsRequest)(nil), // 1: proto.VersionsRequest
	(*VersionResponse)(nil), // 2: proto.VersionResponse
	(*Ver)(nil),             // 3: proto.ver
	(*LanguageRequest)(nil), // 4: proto.LanguageRequest
	(*LanguageList)(nil),    // 5: proto.LanguageList
	(*Selectors)(nil),       // 6: proto.Selectors
	(*Pagination)(nil),      // 7: proto.Pagination
	(*Block)(nil),           // 8: proto.Block
	(*BlockContent)(nil),    // 9: proto.BlockContent
	(*BlockMeta)(nil),       // 10: proto.BlockMeta
	(*BlockRules)(nil),      // 11: proto.BlockRules
}
var file_Queries_Block_proto_queries_proto_depIdxs = []int32{
	3,  // 0: proto.VersionResponse.versions:type_name -> proto.ver
	0,  // 1: proto.Queries.GetBlock:input_type -> proto.BlockRequest
	0,  // 2: proto.Queries.GetBlockContent:input_type -> proto.BlockRequest
	0,  // 3: proto.Queries.GetBlockMeta:input_type -> proto.BlockRequest
	0,  // 4: proto.Queries.GetBlockRules:input_type -> proto.BlockRequest
	1,  // 5: proto.Queries.GetVersions:input_type -> proto.VersionsRequest
	4,  // 6: proto.Queries.GetLanguages:input_type -> proto.LanguageRequest
	8,  // 7: proto.Queries.GetBlock:output_type -> proto.Block
	9,  // 8: proto.Queries.GetBlockContent:output_type -> proto.BlockContent
	10, // 9: proto.Queries.GetBlockMeta:output_type -> proto.BlockMeta
	11, // 10: proto.Queries.GetBlockRules:output_type -> proto.BlockRules
	2,  // 11: proto.Queries.GetVersions:output_type -> proto.VersionResponse
	5,  // 12: proto.Queries.GetLanguages:output_type -> proto.LanguageList
	7,  // [7:13] is the sub-list for method output_type
	1,  // [1:7] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_Queries_Block_proto_queries_proto_init() }
func file_Queries_Block_proto_queries_proto_init() {
	if File_Queries_Block_proto_queries_proto != nil {
		return
	}
	file_Shared_proto_block_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_Queries_Block_proto_queries_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockRequest); i {
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
		file_Queries_Block_proto_queries_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VersionsRequest); i {
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
		file_Queries_Block_proto_queries_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VersionResponse); i {
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
		file_Queries_Block_proto_queries_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ver); i {
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
		file_Queries_Block_proto_queries_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LanguageRequest); i {
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
		file_Queries_Block_proto_queries_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LanguageList); i {
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
		file_Queries_Block_proto_queries_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Selectors); i {
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
		file_Queries_Block_proto_queries_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pagination); i {
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
	file_Queries_Block_proto_queries_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_Queries_Block_proto_queries_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_Queries_Block_proto_queries_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_Queries_Block_proto_queries_proto_goTypes,
		DependencyIndexes: file_Queries_Block_proto_queries_proto_depIdxs,
		MessageInfos:      file_Queries_Block_proto_queries_proto_msgTypes,
	}.Build()
	File_Queries_Block_proto_queries_proto = out.File
	file_Queries_Block_proto_queries_proto_rawDesc = nil
	file_Queries_Block_proto_queries_proto_goTypes = nil
	file_Queries_Block_proto_queries_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// QueriesClient is the client API for Queries service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueriesClient interface {
	GetBlock(ctx context.Context, in *BlockRequest, opts ...grpc.CallOption) (*Block, error)
	GetBlockContent(ctx context.Context, in *BlockRequest, opts ...grpc.CallOption) (*BlockContent, error)
	GetBlockMeta(ctx context.Context, in *BlockRequest, opts ...grpc.CallOption) (*BlockMeta, error)
	GetBlockRules(ctx context.Context, in *BlockRequest, opts ...grpc.CallOption) (*BlockRules, error)
	GetVersions(ctx context.Context, in *VersionsRequest, opts ...grpc.CallOption) (*VersionResponse, error)
	GetLanguages(ctx context.Context, in *LanguageRequest, opts ...grpc.CallOption) (*LanguageList, error)
}

type queriesClient struct {
	cc grpc.ClientConnInterface
}

func NewQueriesClient(cc grpc.ClientConnInterface) QueriesClient {
	return &queriesClient{cc}
}

func (c *queriesClient) GetBlock(ctx context.Context, in *BlockRequest, opts ...grpc.CallOption) (*Block, error) {
	out := new(Block)
	err := c.cc.Invoke(ctx, "/proto.Queries/GetBlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queriesClient) GetBlockContent(ctx context.Context, in *BlockRequest, opts ...grpc.CallOption) (*BlockContent, error) {
	out := new(BlockContent)
	err := c.cc.Invoke(ctx, "/proto.Queries/GetBlockContent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queriesClient) GetBlockMeta(ctx context.Context, in *BlockRequest, opts ...grpc.CallOption) (*BlockMeta, error) {
	out := new(BlockMeta)
	err := c.cc.Invoke(ctx, "/proto.Queries/GetBlockMeta", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queriesClient) GetBlockRules(ctx context.Context, in *BlockRequest, opts ...grpc.CallOption) (*BlockRules, error) {
	out := new(BlockRules)
	err := c.cc.Invoke(ctx, "/proto.Queries/GetBlockRules", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queriesClient) GetVersions(ctx context.Context, in *VersionsRequest, opts ...grpc.CallOption) (*VersionResponse, error) {
	out := new(VersionResponse)
	err := c.cc.Invoke(ctx, "/proto.Queries/GetVersions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queriesClient) GetLanguages(ctx context.Context, in *LanguageRequest, opts ...grpc.CallOption) (*LanguageList, error) {
	out := new(LanguageList)
	err := c.cc.Invoke(ctx, "/proto.Queries/GetLanguages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueriesServer is the server API for Queries service.
type QueriesServer interface {
	GetBlock(context.Context, *BlockRequest) (*Block, error)
	GetBlockContent(context.Context, *BlockRequest) (*BlockContent, error)
	GetBlockMeta(context.Context, *BlockRequest) (*BlockMeta, error)
	GetBlockRules(context.Context, *BlockRequest) (*BlockRules, error)
	GetVersions(context.Context, *VersionsRequest) (*VersionResponse, error)
	GetLanguages(context.Context, *LanguageRequest) (*LanguageList, error)
}

// UnimplementedQueriesServer can be embedded to have forward compatible implementations.
type UnimplementedQueriesServer struct {
}

func (*UnimplementedQueriesServer) GetBlock(context.Context, *BlockRequest) (*Block, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlock not implemented")
}
func (*UnimplementedQueriesServer) GetBlockContent(context.Context, *BlockRequest) (*BlockContent, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlockContent not implemented")
}
func (*UnimplementedQueriesServer) GetBlockMeta(context.Context, *BlockRequest) (*BlockMeta, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlockMeta not implemented")
}
func (*UnimplementedQueriesServer) GetBlockRules(context.Context, *BlockRequest) (*BlockRules, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlockRules not implemented")
}
func (*UnimplementedQueriesServer) GetVersions(context.Context, *VersionsRequest) (*VersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVersions not implemented")
}
func (*UnimplementedQueriesServer) GetLanguages(context.Context, *LanguageRequest) (*LanguageList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLanguages not implemented")
}

func RegisterQueriesServer(s *grpc.Server, srv QueriesServer) {
	s.RegisterService(&_Queries_serviceDesc, srv)
}

func _Queries_GetBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueriesServer).GetBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Queries/GetBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueriesServer).GetBlock(ctx, req.(*BlockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Queries_GetBlockContent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueriesServer).GetBlockContent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Queries/GetBlockContent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueriesServer).GetBlockContent(ctx, req.(*BlockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Queries_GetBlockMeta_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueriesServer).GetBlockMeta(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Queries/GetBlockMeta",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueriesServer).GetBlockMeta(ctx, req.(*BlockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Queries_GetBlockRules_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueriesServer).GetBlockRules(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Queries/GetBlockRules",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueriesServer).GetBlockRules(ctx, req.(*BlockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Queries_GetVersions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VersionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueriesServer).GetVersions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Queries/GetVersions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueriesServer).GetVersions(ctx, req.(*VersionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Queries_GetLanguages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LanguageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueriesServer).GetLanguages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Queries/GetLanguages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueriesServer).GetLanguages(ctx, req.(*LanguageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Queries_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Queries",
	HandlerType: (*QueriesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBlock",
			Handler:    _Queries_GetBlock_Handler,
		},
		{
			MethodName: "GetBlockContent",
			Handler:    _Queries_GetBlockContent_Handler,
		},
		{
			MethodName: "GetBlockMeta",
			Handler:    _Queries_GetBlockMeta_Handler,
		},
		{
			MethodName: "GetBlockRules",
			Handler:    _Queries_GetBlockRules_Handler,
		},
		{
			MethodName: "GetVersions",
			Handler:    _Queries_GetVersions_Handler,
		},
		{
			MethodName: "GetLanguages",
			Handler:    _Queries_GetLanguages_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Queries/Block/proto/queries.proto",
}