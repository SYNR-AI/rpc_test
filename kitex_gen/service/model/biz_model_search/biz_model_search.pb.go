// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.28.2
// source: model/biz_model_search.proto

package biz_model_search

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

type ItemIndex struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ItemId      int64    `protobuf:"varint,1,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	Title       string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Hashtag     []string `protobuf:"bytes,4,rep,name=hashtag,proto3" json:"hashtag,omitempty"`
	CreateTime  int64    `protobuf:"varint,5,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
}

func (x *ItemIndex) Reset() {
	*x = ItemIndex{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_biz_model_search_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ItemIndex) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItemIndex) ProtoMessage() {}

func (x *ItemIndex) ProtoReflect() protoreflect.Message {
	mi := &file_model_biz_model_search_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItemIndex.ProtoReflect.Descriptor instead.
func (*ItemIndex) Descriptor() ([]byte, []int) {
	return file_model_biz_model_search_proto_rawDescGZIP(), []int{0}
}

func (x *ItemIndex) GetItemId() int64 {
	if x != nil {
		return x.ItemId
	}
	return 0
}

func (x *ItemIndex) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ItemIndex) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ItemIndex) GetHashtag() []string {
	if x != nil {
		return x.Hashtag
	}
	return nil
}

func (x *ItemIndex) GetCreateTime() int64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

type ItemIndexUg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ItemId     int64    `protobuf:"varint,1,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	Title      string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Content    []string `protobuf:"bytes,3,rep,name=content,proto3" json:"content,omitempty"`
	Hashtag    []string `protobuf:"bytes,4,rep,name=hashtag,proto3" json:"hashtag,omitempty"`
	CreateTime int64    `protobuf:"varint,5,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	MuserName  []string `protobuf:"bytes,6,rep,name=muser_name,json=muserName,proto3" json:"muser_name,omitempty"`
	ChatCount  int32    `protobuf:"varint,7,opt,name=chat_count,json=chatCount,proto3" json:"chat_count,omitempty"`
}

func (x *ItemIndexUg) Reset() {
	*x = ItemIndexUg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_biz_model_search_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ItemIndexUg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItemIndexUg) ProtoMessage() {}

func (x *ItemIndexUg) ProtoReflect() protoreflect.Message {
	mi := &file_model_biz_model_search_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItemIndexUg.ProtoReflect.Descriptor instead.
func (*ItemIndexUg) Descriptor() ([]byte, []int) {
	return file_model_biz_model_search_proto_rawDescGZIP(), []int{1}
}

func (x *ItemIndexUg) GetItemId() int64 {
	if x != nil {
		return x.ItemId
	}
	return 0
}

func (x *ItemIndexUg) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ItemIndexUg) GetContent() []string {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *ItemIndexUg) GetHashtag() []string {
	if x != nil {
		return x.Hashtag
	}
	return nil
}

func (x *ItemIndexUg) GetCreateTime() int64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *ItemIndexUg) GetMuserName() []string {
	if x != nil {
		return x.MuserName
	}
	return nil
}

func (x *ItemIndexUg) GetChatCount() int32 {
	if x != nil {
		return x.ChatCount
	}
	return 0
}

type CharacterItemIndex struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ItemId       int64 `protobuf:"varint,1,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	MuserId      int64 `protobuf:"varint,2,opt,name=muser_id,json=muserId,proto3" json:"muser_id,omitempty"`
	CreatorId    int64 `protobuf:"varint,3,opt,name=creator_id,json=creatorId,proto3" json:"creator_id,omitempty"`
	CollectionId int64 `protobuf:"varint,4,opt,name=collection_id,json=collectionId,proto3" json:"collection_id,omitempty"`
	CharacterId  int64 `protobuf:"varint,5,opt,name=character_id,json=characterId,proto3" json:"character_id,omitempty"`
}

func (x *CharacterItemIndex) Reset() {
	*x = CharacterItemIndex{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_biz_model_search_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CharacterItemIndex) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CharacterItemIndex) ProtoMessage() {}

func (x *CharacterItemIndex) ProtoReflect() protoreflect.Message {
	mi := &file_model_biz_model_search_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CharacterItemIndex.ProtoReflect.Descriptor instead.
func (*CharacterItemIndex) Descriptor() ([]byte, []int) {
	return file_model_biz_model_search_proto_rawDescGZIP(), []int{2}
}

func (x *CharacterItemIndex) GetItemId() int64 {
	if x != nil {
		return x.ItemId
	}
	return 0
}

func (x *CharacterItemIndex) GetMuserId() int64 {
	if x != nil {
		return x.MuserId
	}
	return 0
}

func (x *CharacterItemIndex) GetCreatorId() int64 {
	if x != nil {
		return x.CreatorId
	}
	return 0
}

func (x *CharacterItemIndex) GetCollectionId() int64 {
	if x != nil {
		return x.CollectionId
	}
	return 0
}

func (x *CharacterItemIndex) GetCharacterId() int64 {
	if x != nil {
		return x.CharacterId
	}
	return 0
}

var File_model_biz_model_search_proto protoreflect.FileDescriptor

var file_model_biz_model_search_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x62, 0x69, 0x7a, 0x5f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x5f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10,
	0x62, 0x69, 0x7a, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x22, 0x97, 0x01, 0x0a, 0x09, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x17,
	0x0a, 0x07, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x69, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x18, 0x0a, 0x07, 0x68, 0x61, 0x73, 0x68, 0x74, 0x61, 0x67, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x07, 0x68, 0x61, 0x73, 0x68, 0x74, 0x61, 0x67, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0xcf, 0x01, 0x0a, 0x0b, 0x49,
	0x74, 0x65, 0x6d, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x55, 0x67, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x74,
	0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x69, 0x74, 0x65,
	0x6d, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x68, 0x61, 0x73, 0x68, 0x74, 0x61, 0x67, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x68, 0x61, 0x73, 0x68, 0x74, 0x61, 0x67, 0x12, 0x1f, 0x0a,
	0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x6d, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x09, 0x6d, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x09, 0x63, 0x68, 0x61, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0xaf, 0x01, 0x0a,
	0x12, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08,
	0x6d, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07,
	0x6d, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x63,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x63,
	0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0b, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x49, 0x64, 0x42, 0x33,
	0x5a, 0x31, 0x72, 0x70, 0x63, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x6b, 0x69, 0x74, 0x65, 0x78,
	0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2f, 0x62, 0x69, 0x7a, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x73, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_model_biz_model_search_proto_rawDescOnce sync.Once
	file_model_biz_model_search_proto_rawDescData = file_model_biz_model_search_proto_rawDesc
)

func file_model_biz_model_search_proto_rawDescGZIP() []byte {
	file_model_biz_model_search_proto_rawDescOnce.Do(func() {
		file_model_biz_model_search_proto_rawDescData = protoimpl.X.CompressGZIP(file_model_biz_model_search_proto_rawDescData)
	})
	return file_model_biz_model_search_proto_rawDescData
}

var file_model_biz_model_search_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_model_biz_model_search_proto_goTypes = []interface{}{
	(*ItemIndex)(nil),          // 0: biz_model_search.ItemIndex
	(*ItemIndexUg)(nil),        // 1: biz_model_search.ItemIndexUg
	(*CharacterItemIndex)(nil), // 2: biz_model_search.CharacterItemIndex
}
var file_model_biz_model_search_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_model_biz_model_search_proto_init() }
func file_model_biz_model_search_proto_init() {
	if File_model_biz_model_search_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_model_biz_model_search_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ItemIndex); i {
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
		file_model_biz_model_search_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ItemIndexUg); i {
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
		file_model_biz_model_search_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CharacterItemIndex); i {
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
			RawDescriptor: file_model_biz_model_search_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_model_biz_model_search_proto_goTypes,
		DependencyIndexes: file_model_biz_model_search_proto_depIdxs,
		MessageInfos:      file_model_biz_model_search_proto_msgTypes,
	}.Build()
	File_model_biz_model_search_proto = out.File
	file_model_biz_model_search_proto_rawDesc = nil
	file_model_biz_model_search_proto_goTypes = nil
	file_model_biz_model_search_proto_depIdxs = nil
}

var _ context.Context
