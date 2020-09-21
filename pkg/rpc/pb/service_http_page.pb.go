// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: service_http_page.proto

package pb

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

// 创建Page
type CreateHTTPPageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusList []string `protobuf:"bytes,1,rep,name=statusList,proto3" json:"statusList,omitempty"`
	Url        string   `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	NewStatus  int32    `protobuf:"varint,3,opt,name=newStatus,proto3" json:"newStatus,omitempty"`
}

func (x *CreateHTTPPageRequest) Reset() {
	*x = CreateHTTPPageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_http_page_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateHTTPPageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateHTTPPageRequest) ProtoMessage() {}

func (x *CreateHTTPPageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_http_page_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateHTTPPageRequest.ProtoReflect.Descriptor instead.
func (*CreateHTTPPageRequest) Descriptor() ([]byte, []int) {
	return file_service_http_page_proto_rawDescGZIP(), []int{0}
}

func (x *CreateHTTPPageRequest) GetStatusList() []string {
	if x != nil {
		return x.StatusList
	}
	return nil
}

func (x *CreateHTTPPageRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *CreateHTTPPageRequest) GetNewStatus() int32 {
	if x != nil {
		return x.NewStatus
	}
	return 0
}

type CreateHTTPPageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageId int64 `protobuf:"varint,1,opt,name=pageId,proto3" json:"pageId,omitempty"`
}

func (x *CreateHTTPPageResponse) Reset() {
	*x = CreateHTTPPageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_http_page_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateHTTPPageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateHTTPPageResponse) ProtoMessage() {}

func (x *CreateHTTPPageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_http_page_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateHTTPPageResponse.ProtoReflect.Descriptor instead.
func (*CreateHTTPPageResponse) Descriptor() ([]byte, []int) {
	return file_service_http_page_proto_rawDescGZIP(), []int{1}
}

func (x *CreateHTTPPageResponse) GetPageId() int64 {
	if x != nil {
		return x.PageId
	}
	return 0
}

// 修改Page
type UpdateHTTPPageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageId     int64    `protobuf:"varint,1,opt,name=pageId,proto3" json:"pageId,omitempty"`
	StatusList []string `protobuf:"bytes,2,rep,name=statusList,proto3" json:"statusList,omitempty"`
	Url        string   `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
	NewStatus  int32    `protobuf:"varint,4,opt,name=newStatus,proto3" json:"newStatus,omitempty"`
}

func (x *UpdateHTTPPageRequest) Reset() {
	*x = UpdateHTTPPageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_http_page_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateHTTPPageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateHTTPPageRequest) ProtoMessage() {}

func (x *UpdateHTTPPageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_http_page_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateHTTPPageRequest.ProtoReflect.Descriptor instead.
func (*UpdateHTTPPageRequest) Descriptor() ([]byte, []int) {
	return file_service_http_page_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateHTTPPageRequest) GetPageId() int64 {
	if x != nil {
		return x.PageId
	}
	return 0
}

func (x *UpdateHTTPPageRequest) GetStatusList() []string {
	if x != nil {
		return x.StatusList
	}
	return nil
}

func (x *UpdateHTTPPageRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *UpdateHTTPPageRequest) GetNewStatus() int32 {
	if x != nil {
		return x.NewStatus
	}
	return 0
}

// 查找单个Page配置
type FindEnabledHTTPPageConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageId int64 `protobuf:"varint,1,opt,name=pageId,proto3" json:"pageId,omitempty"`
}

func (x *FindEnabledHTTPPageConfigRequest) Reset() {
	*x = FindEnabledHTTPPageConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_http_page_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindEnabledHTTPPageConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindEnabledHTTPPageConfigRequest) ProtoMessage() {}

func (x *FindEnabledHTTPPageConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_http_page_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindEnabledHTTPPageConfigRequest.ProtoReflect.Descriptor instead.
func (*FindEnabledHTTPPageConfigRequest) Descriptor() ([]byte, []int) {
	return file_service_http_page_proto_rawDescGZIP(), []int{3}
}

func (x *FindEnabledHTTPPageConfigRequest) GetPageId() int64 {
	if x != nil {
		return x.PageId
	}
	return 0
}

type FindEnabledHTTPPageConfigResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageJSON []byte `protobuf:"bytes,1,opt,name=pageJSON,proto3" json:"pageJSON,omitempty"`
}

func (x *FindEnabledHTTPPageConfigResponse) Reset() {
	*x = FindEnabledHTTPPageConfigResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_http_page_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindEnabledHTTPPageConfigResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindEnabledHTTPPageConfigResponse) ProtoMessage() {}

func (x *FindEnabledHTTPPageConfigResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_http_page_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindEnabledHTTPPageConfigResponse.ProtoReflect.Descriptor instead.
func (*FindEnabledHTTPPageConfigResponse) Descriptor() ([]byte, []int) {
	return file_service_http_page_proto_rawDescGZIP(), []int{4}
}

func (x *FindEnabledHTTPPageConfigResponse) GetPageJSON() []byte {
	if x != nil {
		return x.PageJSON
	}
	return nil
}

var File_service_http_page_proto protoreflect.FileDescriptor

var file_service_http_page_proto_rawDesc = []byte{
	0x0a, 0x17, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x70,
	0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x12, 0x72,
	0x70, 0x63, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x67, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x48, 0x54, 0x54, 0x50, 0x50,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x1c, 0x0a, 0x09,
	0x6e, 0x65, 0x77, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x09, 0x6e, 0x65, 0x77, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x30, 0x0a, 0x16, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x48, 0x54, 0x54, 0x50, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x67, 0x65, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x70, 0x61, 0x67, 0x65, 0x49, 0x64, 0x22, 0x7f, 0x0a, 0x15,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x48, 0x54, 0x54, 0x50, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x67, 0x65, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x70, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x1e, 0x0a,
	0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12,
	0x1c, 0x0a, 0x09, 0x6e, 0x65, 0x77, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x09, 0x6e, 0x65, 0x77, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x3a, 0x0a,
	0x20, 0x46, 0x69, 0x6e, 0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x48, 0x54, 0x54, 0x50,
	0x50, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x67, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x70, 0x61, 0x67, 0x65, 0x49, 0x64, 0x22, 0x3f, 0x0a, 0x21, 0x46, 0x69, 0x6e,
	0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x48, 0x54, 0x54, 0x50, 0x50, 0x61, 0x67, 0x65,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x4a, 0x53, 0x4f, 0x4e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x4a, 0x53, 0x4f, 0x4e, 0x32, 0x87, 0x02, 0x0a, 0x0f, 0x48,
	0x54, 0x54, 0x50, 0x50, 0x61, 0x67, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x47,
	0x0a, 0x0e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x48, 0x54, 0x54, 0x50, 0x50, 0x61, 0x67, 0x65,
	0x12, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x48, 0x54, 0x54, 0x50,
	0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x62,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x48, 0x54, 0x54, 0x50, 0x50, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x0e, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x48, 0x54, 0x54, 0x50, 0x50, 0x61, 0x67, 0x65, 0x12, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x48, 0x54, 0x54, 0x50, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x50, 0x43, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x68, 0x0a, 0x19, 0x66, 0x69,
	0x6e, 0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x48, 0x54, 0x54, 0x50, 0x50, 0x61, 0x67,
	0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x24, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e,
	0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x48, 0x54, 0x54, 0x50, 0x50, 0x61, 0x67, 0x65,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e,
	0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x48, 0x54,
	0x54, 0x50, 0x50, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_http_page_proto_rawDescOnce sync.Once
	file_service_http_page_proto_rawDescData = file_service_http_page_proto_rawDesc
)

func file_service_http_page_proto_rawDescGZIP() []byte {
	file_service_http_page_proto_rawDescOnce.Do(func() {
		file_service_http_page_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_http_page_proto_rawDescData)
	})
	return file_service_http_page_proto_rawDescData
}

var file_service_http_page_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_service_http_page_proto_goTypes = []interface{}{
	(*CreateHTTPPageRequest)(nil),             // 0: pb.CreateHTTPPageRequest
	(*CreateHTTPPageResponse)(nil),            // 1: pb.CreateHTTPPageResponse
	(*UpdateHTTPPageRequest)(nil),             // 2: pb.UpdateHTTPPageRequest
	(*FindEnabledHTTPPageConfigRequest)(nil),  // 3: pb.FindEnabledHTTPPageConfigRequest
	(*FindEnabledHTTPPageConfigResponse)(nil), // 4: pb.FindEnabledHTTPPageConfigResponse
	(*RPCUpdateSuccess)(nil),                  // 5: pb.RPCUpdateSuccess
}
var file_service_http_page_proto_depIdxs = []int32{
	0, // 0: pb.HTTPPageService.createHTTPPage:input_type -> pb.CreateHTTPPageRequest
	2, // 1: pb.HTTPPageService.updateHTTPPage:input_type -> pb.UpdateHTTPPageRequest
	3, // 2: pb.HTTPPageService.findEnabledHTTPPageConfig:input_type -> pb.FindEnabledHTTPPageConfigRequest
	1, // 3: pb.HTTPPageService.createHTTPPage:output_type -> pb.CreateHTTPPageResponse
	5, // 4: pb.HTTPPageService.updateHTTPPage:output_type -> pb.RPCUpdateSuccess
	4, // 5: pb.HTTPPageService.findEnabledHTTPPageConfig:output_type -> pb.FindEnabledHTTPPageConfigResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_service_http_page_proto_init() }
func file_service_http_page_proto_init() {
	if File_service_http_page_proto != nil {
		return
	}
	file_rpc_messages_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_service_http_page_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateHTTPPageRequest); i {
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
		file_service_http_page_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateHTTPPageResponse); i {
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
		file_service_http_page_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateHTTPPageRequest); i {
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
		file_service_http_page_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindEnabledHTTPPageConfigRequest); i {
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
		file_service_http_page_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindEnabledHTTPPageConfigResponse); i {
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
			RawDescriptor: file_service_http_page_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_http_page_proto_goTypes,
		DependencyIndexes: file_service_http_page_proto_depIdxs,
		MessageInfos:      file_service_http_page_proto_msgTypes,
	}.Build()
	File_service_http_page_proto = out.File
	file_service_http_page_proto_rawDesc = nil
	file_service_http_page_proto_goTypes = nil
	file_service_http_page_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// HTTPPageServiceClient is the client API for HTTPPageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HTTPPageServiceClient interface {
	// 创建Page
	CreateHTTPPage(ctx context.Context, in *CreateHTTPPageRequest, opts ...grpc.CallOption) (*CreateHTTPPageResponse, error)
	// 修改Page
	UpdateHTTPPage(ctx context.Context, in *UpdateHTTPPageRequest, opts ...grpc.CallOption) (*RPCUpdateSuccess, error)
	// 查找单个Page配置
	FindEnabledHTTPPageConfig(ctx context.Context, in *FindEnabledHTTPPageConfigRequest, opts ...grpc.CallOption) (*FindEnabledHTTPPageConfigResponse, error)
}

type hTTPPageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHTTPPageServiceClient(cc grpc.ClientConnInterface) HTTPPageServiceClient {
	return &hTTPPageServiceClient{cc}
}

func (c *hTTPPageServiceClient) CreateHTTPPage(ctx context.Context, in *CreateHTTPPageRequest, opts ...grpc.CallOption) (*CreateHTTPPageResponse, error) {
	out := new(CreateHTTPPageResponse)
	err := c.cc.Invoke(ctx, "/pb.HTTPPageService/createHTTPPage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hTTPPageServiceClient) UpdateHTTPPage(ctx context.Context, in *UpdateHTTPPageRequest, opts ...grpc.CallOption) (*RPCUpdateSuccess, error) {
	out := new(RPCUpdateSuccess)
	err := c.cc.Invoke(ctx, "/pb.HTTPPageService/updateHTTPPage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hTTPPageServiceClient) FindEnabledHTTPPageConfig(ctx context.Context, in *FindEnabledHTTPPageConfigRequest, opts ...grpc.CallOption) (*FindEnabledHTTPPageConfigResponse, error) {
	out := new(FindEnabledHTTPPageConfigResponse)
	err := c.cc.Invoke(ctx, "/pb.HTTPPageService/findEnabledHTTPPageConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HTTPPageServiceServer is the server API for HTTPPageService service.
type HTTPPageServiceServer interface {
	// 创建Page
	CreateHTTPPage(context.Context, *CreateHTTPPageRequest) (*CreateHTTPPageResponse, error)
	// 修改Page
	UpdateHTTPPage(context.Context, *UpdateHTTPPageRequest) (*RPCUpdateSuccess, error)
	// 查找单个Page配置
	FindEnabledHTTPPageConfig(context.Context, *FindEnabledHTTPPageConfigRequest) (*FindEnabledHTTPPageConfigResponse, error)
}

// UnimplementedHTTPPageServiceServer can be embedded to have forward compatible implementations.
type UnimplementedHTTPPageServiceServer struct {
}

func (*UnimplementedHTTPPageServiceServer) CreateHTTPPage(context.Context, *CreateHTTPPageRequest) (*CreateHTTPPageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateHTTPPage not implemented")
}
func (*UnimplementedHTTPPageServiceServer) UpdateHTTPPage(context.Context, *UpdateHTTPPageRequest) (*RPCUpdateSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateHTTPPage not implemented")
}
func (*UnimplementedHTTPPageServiceServer) FindEnabledHTTPPageConfig(context.Context, *FindEnabledHTTPPageConfigRequest) (*FindEnabledHTTPPageConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindEnabledHTTPPageConfig not implemented")
}

func RegisterHTTPPageServiceServer(s *grpc.Server, srv HTTPPageServiceServer) {
	s.RegisterService(&_HTTPPageService_serviceDesc, srv)
}

func _HTTPPageService_CreateHTTPPage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateHTTPPageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HTTPPageServiceServer).CreateHTTPPage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.HTTPPageService/CreateHTTPPage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HTTPPageServiceServer).CreateHTTPPage(ctx, req.(*CreateHTTPPageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HTTPPageService_UpdateHTTPPage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateHTTPPageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HTTPPageServiceServer).UpdateHTTPPage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.HTTPPageService/UpdateHTTPPage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HTTPPageServiceServer).UpdateHTTPPage(ctx, req.(*UpdateHTTPPageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HTTPPageService_FindEnabledHTTPPageConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindEnabledHTTPPageConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HTTPPageServiceServer).FindEnabledHTTPPageConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.HTTPPageService/FindEnabledHTTPPageConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HTTPPageServiceServer).FindEnabledHTTPPageConfig(ctx, req.(*FindEnabledHTTPPageConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _HTTPPageService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.HTTPPageService",
	HandlerType: (*HTTPPageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "createHTTPPage",
			Handler:    _HTTPPageService_CreateHTTPPage_Handler,
		},
		{
			MethodName: "updateHTTPPage",
			Handler:    _HTTPPageService_UpdateHTTPPage_Handler,
		},
		{
			MethodName: "findEnabledHTTPPageConfig",
			Handler:    _HTTPPageService_FindEnabledHTTPPageConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_http_page.proto",
}
