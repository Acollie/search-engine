// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: protos/service/spider.proto

package spider

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

type SeenListRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The number of sites to return default is 10
	Limit         int32 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SeenListRequest) Reset() {
	*x = SeenListRequest{}
	mi := &file_protos_service_spider_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SeenListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SeenListRequest) ProtoMessage() {}

func (x *SeenListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_service_spider_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SeenListRequest.ProtoReflect.Descriptor instead.
func (*SeenListRequest) Descriptor() ([]byte, []int) {
	return file_protos_service_spider_proto_rawDescGZIP(), []int{0}
}

func (x *SeenListRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type SeenListResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	SeenSites     []*Page                `protobuf:"bytes,1,rep,name=seen_sites,json=seenSites,proto3" json:"seen_sites,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SeenListResponse) Reset() {
	*x = SeenListResponse{}
	mi := &file_protos_service_spider_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SeenListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SeenListResponse) ProtoMessage() {}

func (x *SeenListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_service_spider_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SeenListResponse.ProtoReflect.Descriptor instead.
func (*SeenListResponse) Descriptor() ([]byte, []int) {
	return file_protos_service_spider_proto_rawDescGZIP(), []int{1}
}

func (x *SeenListResponse) GetSeenSites() []*Page {
	if x != nil {
		return x.SeenSites
	}
	return nil
}

type HealthRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *HealthRequest) Reset() {
	*x = HealthRequest{}
	mi := &file_protos_service_spider_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HealthRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthRequest) ProtoMessage() {}

func (x *HealthRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_service_spider_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthRequest.ProtoReflect.Descriptor instead.
func (*HealthRequest) Descriptor() ([]byte, []int) {
	return file_protos_service_spider_proto_rawDescGZIP(), []int{2}
}

type HealthResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        *Status                `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *HealthResponse) Reset() {
	*x = HealthResponse{}
	mi := &file_protos_service_spider_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HealthResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthResponse) ProtoMessage() {}

func (x *HealthResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_service_spider_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthResponse.ProtoReflect.Descriptor instead.
func (*HealthResponse) Descriptor() ([]byte, []int) {
	return file_protos_service_spider_proto_rawDescGZIP(), []int{3}
}

func (x *HealthResponse) GetStatus() *Status {
	if x != nil {
		return x.Status
	}
	return nil
}

type Status struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Healthy       bool                   `protobuf:"varint,2,opt,name=healthy,proto3" json:"healthy,omitempty"`
	Tps           int32                  `protobuf:"varint,3,opt,name=tps,proto3" json:"tps,omitempty"`
	SeenSites     int64                  `protobuf:"varint,4,opt,name=seen_sites,json=seenSites,proto3" json:"seen_sites,omitempty"`
	QueuedSites   int64                  `protobuf:"varint,5,opt,name=queued_sites,json=queuedSites,proto3" json:"queued_sites,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Status) Reset() {
	*x = Status{}
	mi := &file_protos_service_spider_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_protos_service_spider_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_protos_service_spider_proto_rawDescGZIP(), []int{4}
}

func (x *Status) GetHealthy() bool {
	if x != nil {
		return x.Healthy
	}
	return false
}

func (x *Status) GetTps() int32 {
	if x != nil {
		return x.Tps
	}
	return 0
}

func (x *Status) GetSeenSites() int64 {
	if x != nil {
		return x.SeenSites
	}
	return 0
}

func (x *Status) GetQueuedSites() int64 {
	if x != nil {
		return x.QueuedSites
	}
	return 0
}

type Page struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Url           string                 `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Body          string                 `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	Links         []string               `protobuf:"bytes,3,rep,name=links,proto3" json:"links,omitempty"`
	Prominence    int32                  `protobuf:"varint,4,opt,name=prominence,proto3" json:"prominence,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Page) Reset() {
	*x = Page{}
	mi := &file_protos_service_spider_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Page) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Page) ProtoMessage() {}

func (x *Page) ProtoReflect() protoreflect.Message {
	mi := &file_protos_service_spider_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Page.ProtoReflect.Descriptor instead.
func (*Page) Descriptor() ([]byte, []int) {
	return file_protos_service_spider_proto_rawDescGZIP(), []int{5}
}

func (x *Page) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Page) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *Page) GetLinks() []string {
	if x != nil {
		return x.Links
	}
	return nil
}

func (x *Page) GetProminence() int32 {
	if x != nil {
		return x.Prominence
	}
	return 0
}

var File_protos_service_spider_proto protoreflect.FileDescriptor

var file_protos_service_spider_proto_rawDesc = string([]byte{
	0x0a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2f, 0x73, 0x70, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x27, 0x0a, 0x0f, 0x53, 0x65, 0x65, 0x6e, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22,
	0x40, 0x0a, 0x10, 0x53, 0x65, 0x65, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x0a, 0x73, 0x65, 0x65, 0x6e, 0x5f, 0x73, 0x69, 0x74, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x09, 0x73, 0x65, 0x65, 0x6e, 0x53, 0x69, 0x74, 0x65,
	0x73, 0x22, 0x0f, 0x0a, 0x0d, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x39, 0x0a, 0x0e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x76, 0x0a,
	0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x68, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x70, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03,
	0x74, 0x70, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x65, 0x6e, 0x5f, 0x73, 0x69, 0x74, 0x65,
	0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x65, 0x65, 0x6e, 0x53, 0x69, 0x74,
	0x65, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x71, 0x75, 0x65, 0x75, 0x65, 0x64, 0x5f, 0x73, 0x69, 0x74,
	0x65, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x71, 0x75, 0x65, 0x75, 0x65, 0x64,
	0x53, 0x69, 0x74, 0x65, 0x73, 0x22, 0x62, 0x0a, 0x04, 0x50, 0x61, 0x67, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12,
	0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62,
	0x6f, 0x64, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x05, 0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x6f,
	0x6d, 0x69, 0x6e, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x70,
	0x72, 0x6f, 0x6d, 0x69, 0x6e, 0x65, 0x6e, 0x63, 0x65, 0x32, 0x8a, 0x01, 0x0a, 0x06, 0x53, 0x70,
	0x69, 0x64, 0x65, 0x72, 0x12, 0x42, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x53, 0x65, 0x65, 0x6e, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x65,
	0x65, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x65, 0x65, 0x6e, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x48,
	0x65, 0x61, 0x6c, 0x74, 0x68, 0x12, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x10, 0x5a, 0x0e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2f, 0x73, 0x70, 0x69, 0x64, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_protos_service_spider_proto_rawDescOnce sync.Once
	file_protos_service_spider_proto_rawDescData []byte
)

func file_protos_service_spider_proto_rawDescGZIP() []byte {
	file_protos_service_spider_proto_rawDescOnce.Do(func() {
		file_protos_service_spider_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_protos_service_spider_proto_rawDesc), len(file_protos_service_spider_proto_rawDesc)))
	})
	return file_protos_service_spider_proto_rawDescData
}

var file_protos_service_spider_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_protos_service_spider_proto_goTypes = []any{
	(*SeenListRequest)(nil),  // 0: service.SeenListRequest
	(*SeenListResponse)(nil), // 1: service.SeenListResponse
	(*HealthRequest)(nil),    // 2: service.HealthRequest
	(*HealthResponse)(nil),   // 3: service.HealthResponse
	(*Status)(nil),           // 4: service.Status
	(*Page)(nil),             // 5: service.Page
}
var file_protos_service_spider_proto_depIdxs = []int32{
	5, // 0: service.SeenListResponse.seen_sites:type_name -> service.Page
	4, // 1: service.HealthResponse.status:type_name -> service.Status
	0, // 2: service.Spider.GetSeenList:input_type -> service.SeenListRequest
	2, // 3: service.Spider.GetHealth:input_type -> service.HealthRequest
	1, // 4: service.Spider.GetSeenList:output_type -> service.SeenListResponse
	3, // 5: service.Spider.GetHealth:output_type -> service.HealthResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_protos_service_spider_proto_init() }
func file_protos_service_spider_proto_init() {
	if File_protos_service_spider_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_protos_service_spider_proto_rawDesc), len(file_protos_service_spider_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_service_spider_proto_goTypes,
		DependencyIndexes: file_protos_service_spider_proto_depIdxs,
		MessageInfos:      file_protos_service_spider_proto_msgTypes,
	}.Build()
	File_protos_service_spider_proto = out.File
	file_protos_service_spider_proto_goTypes = nil
	file_protos_service_spider_proto_depIdxs = nil
}
