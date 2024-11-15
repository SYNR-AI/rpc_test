// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.28.2
// source: base.proto

package base

import (
	context "context"
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

type CommonErrorCode int32

const (
	CommonErrorCode_COMM_ERR_UNSPECIFIED        CommonErrorCode = 0
	CommonErrorCode_COMM_ERR_INVALID_PARAMS     CommonErrorCode = 1
	CommonErrorCode_COMM_ERR_NO_PERMISSION      CommonErrorCode = 2
	CommonErrorCode_COMM_ERR_NOT_FOUND          CommonErrorCode = 3
	CommonErrorCode_COMM_ERR_INVALID_STATE      CommonErrorCode = 4
	CommonErrorCode_COMM_ERR_NO_MODIFICATION    CommonErrorCode = 5
	CommonErrorCode_COMM_ERR_PORNOGRAPHIC_BLOCK CommonErrorCode = 6
	CommonErrorCode_COMM_ERR_TNS_BLOCK          CommonErrorCode = 7
	CommonErrorCode_COMM_ERR_INTERNAL_ERROR     CommonErrorCode = 99
)

// Enum value maps for CommonErrorCode.
var (
	CommonErrorCode_name = map[int32]string{
		0:  "COMM_ERR_UNSPECIFIED",
		1:  "COMM_ERR_INVALID_PARAMS",
		2:  "COMM_ERR_NO_PERMISSION",
		3:  "COMM_ERR_NOT_FOUND",
		4:  "COMM_ERR_INVALID_STATE",
		5:  "COMM_ERR_NO_MODIFICATION",
		6:  "COMM_ERR_PORNOGRAPHIC_BLOCK",
		7:  "COMM_ERR_TNS_BLOCK",
		99: "COMM_ERR_INTERNAL_ERROR",
	}
	CommonErrorCode_value = map[string]int32{
		"COMM_ERR_UNSPECIFIED":        0,
		"COMM_ERR_INVALID_PARAMS":     1,
		"COMM_ERR_NO_PERMISSION":      2,
		"COMM_ERR_NOT_FOUND":          3,
		"COMM_ERR_INVALID_STATE":      4,
		"COMM_ERR_NO_MODIFICATION":    5,
		"COMM_ERR_PORNOGRAPHIC_BLOCK": 6,
		"COMM_ERR_TNS_BLOCK":          7,
		"COMM_ERR_INTERNAL_ERROR":     99,
	}
)

func (x CommonErrorCode) Enum() *CommonErrorCode {
	p := new(CommonErrorCode)
	*p = x
	return p
}

func (x CommonErrorCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CommonErrorCode) Descriptor() protoreflect.EnumDescriptor {
	return file_base_proto_enumTypes[0].Descriptor()
}

func (CommonErrorCode) Type() protoreflect.EnumType {
	return &file_base_proto_enumTypes[0]
}

func (x CommonErrorCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CommonErrorCode.Descriptor instead.
func (CommonErrorCode) EnumDescriptor() ([]byte, []int) {
	return file_base_proto_rawDescGZIP(), []int{0}
}

type ScanOrderBy int32

const (
	ScanOrderBy_UNKNOWN ScanOrderBy = 0
	ScanOrderBy_ASC     ScanOrderBy = 1
	ScanOrderBy_DESC    ScanOrderBy = 2
)

// Enum value maps for ScanOrderBy.
var (
	ScanOrderBy_name = map[int32]string{
		0: "UNKNOWN",
		1: "ASC",
		2: "DESC",
	}
	ScanOrderBy_value = map[string]int32{
		"UNKNOWN": 0,
		"ASC":     1,
		"DESC":    2,
	}
)

func (x ScanOrderBy) Enum() *ScanOrderBy {
	p := new(ScanOrderBy)
	*p = x
	return p
}

func (x ScanOrderBy) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ScanOrderBy) Descriptor() protoreflect.EnumDescriptor {
	return file_base_proto_enumTypes[1].Descriptor()
}

func (ScanOrderBy) Type() protoreflect.EnumType {
	return &file_base_proto_enumTypes[1]
}

func (x ScanOrderBy) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ScanOrderBy.Descriptor instead.
func (ScanOrderBy) EnumDescriptor() ([]byte, []int) {
	return file_base_proto_rawDescGZIP(), []int{1}
}

type BaseResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrorCode    int32  `protobuf:"varint,1,opt,name=error_code,json=errorCode,proto3" json:"error_code,omitempty"`         // error code[1-99: common error code] 99+ custom error code
	ErrorMessage string `protobuf:"bytes,2,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"` // error message
}

func (x *BaseResp) Reset() {
	*x = BaseResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaseResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaseResp) ProtoMessage() {}

func (x *BaseResp) ProtoReflect() protoreflect.Message {
	mi := &file_base_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaseResp.ProtoReflect.Descriptor instead.
func (*BaseResp) Descriptor() ([]byte, []int) {
	return file_base_proto_rawDescGZIP(), []int{0}
}

func (x *BaseResp) GetErrorCode() int32 {
	if x != nil {
		return x.ErrorCode
	}
	return 0
}

func (x *BaseResp) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

type TimeRange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Start int64 `protobuf:"varint,1,opt,name=start,proto3" json:"start,omitempty"` // >= millisecond
	End   int64 `protobuf:"varint,2,opt,name=end,proto3" json:"end,omitempty"`     // <= millisecond
}

func (x *TimeRange) Reset() {
	*x = TimeRange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TimeRange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimeRange) ProtoMessage() {}

func (x *TimeRange) ProtoReflect() protoreflect.Message {
	mi := &file_base_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimeRange.ProtoReflect.Descriptor instead.
func (*TimeRange) Descriptor() ([]byte, []int) {
	return file_base_proto_rawDescGZIP(), []int{1}
}

func (x *TimeRange) GetStart() int64 {
	if x != nil {
		return x.Start
	}
	return 0
}

func (x *TimeRange) GetEnd() int64 {
	if x != nil {
		return x.End
	}
	return 0
}

type ScanOption struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cursor    int64       `protobuf:"varint,1,opt,name=cursor,proto3" json:"cursor,omitempty"`
	Limit     int64       `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	OrderBy   ScanOrderBy `protobuf:"varint,3,opt,name=order_by,json=orderBy,proto3,enum=base.ScanOrderBy" json:"order_by,omitempty"`
	TimeRange *TimeRange  `protobuf:"bytes,4,opt,name=time_range,json=timeRange,proto3" json:"time_range,omitempty"`
}

func (x *ScanOption) Reset() {
	*x = ScanOption{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScanOption) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScanOption) ProtoMessage() {}

func (x *ScanOption) ProtoReflect() protoreflect.Message {
	mi := &file_base_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScanOption.ProtoReflect.Descriptor instead.
func (*ScanOption) Descriptor() ([]byte, []int) {
	return file_base_proto_rawDescGZIP(), []int{2}
}

func (x *ScanOption) GetCursor() int64 {
	if x != nil {
		return x.Cursor
	}
	return 0
}

func (x *ScanOption) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ScanOption) GetOrderBy() ScanOrderBy {
	if x != nil {
		return x.OrderBy
	}
	return ScanOrderBy_UNKNOWN
}

func (x *ScanOption) GetTimeRange() *TimeRange {
	if x != nil {
		return x.TimeRange
	}
	return nil
}

type ABExperiment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExperimentId string `protobuf:"bytes,1,opt,name=experiment_id,json=experimentId,proto3" json:"experiment_id,omitempty"`
	VariationId  int32  `protobuf:"varint,2,opt,name=variation_id,json=variationId,proto3" json:"variation_id,omitempty"`
}

func (x *ABExperiment) Reset() {
	*x = ABExperiment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ABExperiment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ABExperiment) ProtoMessage() {}

func (x *ABExperiment) ProtoReflect() protoreflect.Message {
	mi := &file_base_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ABExperiment.ProtoReflect.Descriptor instead.
func (*ABExperiment) Descriptor() ([]byte, []int) {
	return file_base_proto_rawDescGZIP(), []int{3}
}

func (x *ABExperiment) GetExperimentId() string {
	if x != nil {
		return x.ExperimentId
	}
	return ""
}

func (x *ABExperiment) GetVariationId() int32 {
	if x != nil {
		return x.VariationId
	}
	return 0
}

var File_base_proto protoreflect.FileDescriptor

var file_base_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x62, 0x61,
	0x73, 0x65, 0x22, 0x4e, 0x0a, 0x08, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1d,
	0x0a, 0x0a, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x23, 0x0a,
	0x0d, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x22, 0x33, 0x0a, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x03, 0x65, 0x6e, 0x64, 0x22, 0x98, 0x01, 0x0a, 0x0a, 0x53, 0x63, 0x61, 0x6e,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x12, 0x14,
	0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x12, 0x2c, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x62, 0x79,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x53, 0x63,
	0x61, 0x6e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x42, 0x79, 0x12, 0x2e, 0x0a, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e,
	0x67, 0x65, 0x22, 0x56, 0x0a, 0x0c, 0x41, 0x42, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65,
	0x6e, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x78, 0x70, 0x65, 0x72,
	0x69, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x76, 0x61, 0x72, 0x69, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x76,
	0x61, 0x72, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x2a, 0x8c, 0x02, 0x0a, 0x0f, 0x43,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x18,
	0x0a, 0x14, 0x43, 0x4f, 0x4d, 0x4d, 0x5f, 0x45, 0x52, 0x52, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1b, 0x0a, 0x17, 0x43, 0x4f, 0x4d, 0x4d,
	0x5f, 0x45, 0x52, 0x52, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x50, 0x41, 0x52,
	0x41, 0x4d, 0x53, 0x10, 0x01, 0x12, 0x1a, 0x0a, 0x16, 0x43, 0x4f, 0x4d, 0x4d, 0x5f, 0x45, 0x52,
	0x52, 0x5f, 0x4e, 0x4f, 0x5f, 0x50, 0x45, 0x52, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x10,
	0x02, 0x12, 0x16, 0x0a, 0x12, 0x43, 0x4f, 0x4d, 0x4d, 0x5f, 0x45, 0x52, 0x52, 0x5f, 0x4e, 0x4f,
	0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x03, 0x12, 0x1a, 0x0a, 0x16, 0x43, 0x4f, 0x4d,
	0x4d, 0x5f, 0x45, 0x52, 0x52, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x53, 0x54,
	0x41, 0x54, 0x45, 0x10, 0x04, 0x12, 0x1c, 0x0a, 0x18, 0x43, 0x4f, 0x4d, 0x4d, 0x5f, 0x45, 0x52,
	0x52, 0x5f, 0x4e, 0x4f, 0x5f, 0x4d, 0x4f, 0x44, 0x49, 0x46, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f,
	0x4e, 0x10, 0x05, 0x12, 0x1f, 0x0a, 0x1b, 0x43, 0x4f, 0x4d, 0x4d, 0x5f, 0x45, 0x52, 0x52, 0x5f,
	0x50, 0x4f, 0x52, 0x4e, 0x4f, 0x47, 0x52, 0x41, 0x50, 0x48, 0x49, 0x43, 0x5f, 0x42, 0x4c, 0x4f,
	0x43, 0x4b, 0x10, 0x06, 0x12, 0x16, 0x0a, 0x12, 0x43, 0x4f, 0x4d, 0x4d, 0x5f, 0x45, 0x52, 0x52,
	0x5f, 0x54, 0x4e, 0x53, 0x5f, 0x42, 0x4c, 0x4f, 0x43, 0x4b, 0x10, 0x07, 0x12, 0x1b, 0x0a, 0x17,
	0x43, 0x4f, 0x4d, 0x4d, 0x5f, 0x45, 0x52, 0x52, 0x5f, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x4e, 0x41,
	0x4c, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x63, 0x2a, 0x2d, 0x0a, 0x0b, 0x53, 0x63, 0x61,
	0x6e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e,
	0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x53, 0x43, 0x10, 0x01, 0x12, 0x08,
	0x0a, 0x04, 0x44, 0x45, 0x53, 0x43, 0x10, 0x02, 0x42, 0x19, 0x5a, 0x17, 0x72, 0x70, 0x63, 0x5f,
	0x74, 0x65, 0x73, 0x74, 0x2f, 0x6b, 0x69, 0x74, 0x65, 0x78, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x62,
	0x61, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_base_proto_rawDescOnce sync.Once
	file_base_proto_rawDescData = file_base_proto_rawDesc
)

func file_base_proto_rawDescGZIP() []byte {
	file_base_proto_rawDescOnce.Do(func() {
		file_base_proto_rawDescData = protoimpl.X.CompressGZIP(file_base_proto_rawDescData)
	})
	return file_base_proto_rawDescData
}

var file_base_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_base_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_base_proto_goTypes = []interface{}{
	(CommonErrorCode)(0), // 0: base.CommonErrorCode
	(ScanOrderBy)(0),     // 1: base.ScanOrderBy
	(*BaseResp)(nil),     // 2: base.BaseResp
	(*TimeRange)(nil),    // 3: base.TimeRange
	(*ScanOption)(nil),   // 4: base.ScanOption
	(*ABExperiment)(nil), // 5: base.ABExperiment
}
var file_base_proto_depIdxs = []int32{
	1, // 0: base.ScanOption.order_by:type_name -> base.ScanOrderBy
	3, // 1: base.ScanOption.time_range:type_name -> base.TimeRange
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_base_proto_init() }
func file_base_proto_init() {
	if File_base_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_base_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaseResp); i {
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
		file_base_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TimeRange); i {
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
		file_base_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScanOption); i {
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
		file_base_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ABExperiment); i {
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
			RawDescriptor: file_base_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_base_proto_goTypes,
		DependencyIndexes: file_base_proto_depIdxs,
		EnumInfos:         file_base_proto_enumTypes,
		MessageInfos:      file_base_proto_msgTypes,
	}.Build()
	File_base_proto = out.File
	file_base_proto_rawDesc = nil
	file_base_proto_goTypes = nil
	file_base_proto_depIdxs = nil
}

var _ context.Context
