// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.19.4
// source: models/model_user_traffic_bill.proto

package pb

import (
	proto "github.com/golang/protobuf/proto"
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

// 用户流量带宽子账单
type UserTrafficBill struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                  int64       `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	BillId              int64       `protobuf:"varint,2,opt,name=billId,proto3" json:"billId,omitempty"`
	NodeRegionId        int64       `protobuf:"varint,3,opt,name=nodeRegionId,proto3" json:"nodeRegionId,omitempty"`
	Amount              float32     `protobuf:"fixed32,4,opt,name=amount,proto3" json:"amount,omitempty"`
	BandwidthMB         float32     `protobuf:"fixed32,5,opt,name=bandwidthMB,proto3" json:"bandwidthMB,omitempty"`
	BandwidthPercentile int32       `protobuf:"varint,6,opt,name=bandwidthPercentile,proto3" json:"bandwidthPercentile,omitempty"`
	TrafficGB           float32     `protobuf:"fixed32,7,opt,name=trafficGB,proto3" json:"trafficGB,omitempty"`
	PricePerUnit        float32     `protobuf:"fixed32,8,opt,name=pricePerUnit,proto3" json:"pricePerUnit,omitempty"`
	PriceType           string      `protobuf:"bytes,9,opt,name=priceType,proto3" json:"priceType,omitempty"`
	NodeRegion          *NodeRegion `protobuf:"bytes,30,opt,name=nodeRegion,proto3" json:"nodeRegion,omitempty"`
}

func (x *UserTrafficBill) Reset() {
	*x = UserTrafficBill{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_model_user_traffic_bill_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserTrafficBill) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserTrafficBill) ProtoMessage() {}

func (x *UserTrafficBill) ProtoReflect() protoreflect.Message {
	mi := &file_models_model_user_traffic_bill_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserTrafficBill.ProtoReflect.Descriptor instead.
func (*UserTrafficBill) Descriptor() ([]byte, []int) {
	return file_models_model_user_traffic_bill_proto_rawDescGZIP(), []int{0}
}

func (x *UserTrafficBill) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserTrafficBill) GetBillId() int64 {
	if x != nil {
		return x.BillId
	}
	return 0
}

func (x *UserTrafficBill) GetNodeRegionId() int64 {
	if x != nil {
		return x.NodeRegionId
	}
	return 0
}

func (x *UserTrafficBill) GetAmount() float32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *UserTrafficBill) GetBandwidthMB() float32 {
	if x != nil {
		return x.BandwidthMB
	}
	return 0
}

func (x *UserTrafficBill) GetBandwidthPercentile() int32 {
	if x != nil {
		return x.BandwidthPercentile
	}
	return 0
}

func (x *UserTrafficBill) GetTrafficGB() float32 {
	if x != nil {
		return x.TrafficGB
	}
	return 0
}

func (x *UserTrafficBill) GetPricePerUnit() float32 {
	if x != nil {
		return x.PricePerUnit
	}
	return 0
}

func (x *UserTrafficBill) GetPriceType() string {
	if x != nil {
		return x.PriceType
	}
	return ""
}

func (x *UserTrafficBill) GetNodeRegion() *NodeRegion {
	if x != nil {
		return x.NodeRegion
	}
	return nil
}

var File_models_model_user_traffic_bill_proto protoreflect.FileDescriptor

var file_models_model_user_traffic_bill_proto_rawDesc = []byte{
	0x0a, 0x24, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x5f, 0x62, 0x69, 0x6c, 0x6c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x1e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x72, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd9, 0x02, 0x0a, 0x0f, 0x55,
	0x73, 0x65, 0x72, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x42, 0x69, 0x6c, 0x6c, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x62, 0x69, 0x6c, 0x6c, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x62, 0x69, 0x6c, 0x6c, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x6e, 0x6f, 0x64, 0x65, 0x52, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x6e, 0x6f,
	0x64, 0x65, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x62, 0x61, 0x6e, 0x64, 0x77, 0x69, 0x64, 0x74, 0x68, 0x4d,
	0x42, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x62, 0x61, 0x6e, 0x64, 0x77, 0x69, 0x64,
	0x74, 0x68, 0x4d, 0x42, 0x12, 0x30, 0x0a, 0x13, 0x62, 0x61, 0x6e, 0x64, 0x77, 0x69, 0x64, 0x74,
	0x68, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x69, 0x6c, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x13, 0x62, 0x61, 0x6e, 0x64, 0x77, 0x69, 0x64, 0x74, 0x68, 0x50, 0x65, 0x72, 0x63,
	0x65, 0x6e, 0x74, 0x69, 0x6c, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69,
	0x63, 0x47, 0x42, 0x18, 0x07, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x74, 0x72, 0x61, 0x66, 0x66,
	0x69, 0x63, 0x47, 0x42, 0x12, 0x22, 0x0a, 0x0c, 0x70, 0x72, 0x69, 0x63, 0x65, 0x50, 0x65, 0x72,
	0x55, 0x6e, 0x69, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0c, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x50, 0x65, 0x72, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2e, 0x0a, 0x0a, 0x6e, 0x6f, 0x64, 0x65, 0x52, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x62, 0x2e,
	0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x6e, 0x6f, 0x64, 0x65,
	0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_models_model_user_traffic_bill_proto_rawDescOnce sync.Once
	file_models_model_user_traffic_bill_proto_rawDescData = file_models_model_user_traffic_bill_proto_rawDesc
)

func file_models_model_user_traffic_bill_proto_rawDescGZIP() []byte {
	file_models_model_user_traffic_bill_proto_rawDescOnce.Do(func() {
		file_models_model_user_traffic_bill_proto_rawDescData = protoimpl.X.CompressGZIP(file_models_model_user_traffic_bill_proto_rawDescData)
	})
	return file_models_model_user_traffic_bill_proto_rawDescData
}

var file_models_model_user_traffic_bill_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_models_model_user_traffic_bill_proto_goTypes = []interface{}{
	(*UserTrafficBill)(nil), // 0: pb.UserTrafficBill
	(*NodeRegion)(nil),      // 1: pb.NodeRegion
}
var file_models_model_user_traffic_bill_proto_depIdxs = []int32{
	1, // 0: pb.UserTrafficBill.nodeRegion:type_name -> pb.NodeRegion
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_models_model_user_traffic_bill_proto_init() }
func file_models_model_user_traffic_bill_proto_init() {
	if File_models_model_user_traffic_bill_proto != nil {
		return
	}
	file_models_model_node_region_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_models_model_user_traffic_bill_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserTrafficBill); i {
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
			RawDescriptor: file_models_model_user_traffic_bill_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_models_model_user_traffic_bill_proto_goTypes,
		DependencyIndexes: file_models_model_user_traffic_bill_proto_depIdxs,
		MessageInfos:      file_models_model_user_traffic_bill_proto_msgTypes,
	}.Build()
	File_models_model_user_traffic_bill_proto = out.File
	file_models_model_user_traffic_bill_proto_rawDesc = nil
	file_models_model_user_traffic_bill_proto_goTypes = nil
	file_models_model_user_traffic_bill_proto_depIdxs = nil
}
