// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.22.3
// source: proto/room/room.proto

package room

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

type NewRoom struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *NewRoom) Reset() {
	*x = NewRoom{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_room_room_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewRoom) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewRoom) ProtoMessage() {}

func (x *NewRoom) ProtoReflect() protoreflect.Message {
	mi := &file_proto_room_room_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewRoom.ProtoReflect.Descriptor instead.
func (*NewRoom) Descriptor() ([]byte, []int) {
	return file_proto_room_room_proto_rawDescGZIP(), []int{0}
}

func (x *NewRoom) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Room struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Room) Reset() {
	*x = Room{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_room_room_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Room) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Room) ProtoMessage() {}

func (x *Room) ProtoReflect() protoreflect.Message {
	mi := &file_proto_room_room_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Room.ProtoReflect.Descriptor instead.
func (*Room) Descriptor() ([]byte, []int) {
	return file_proto_room_room_proto_rawDescGZIP(), []int{1}
}

func (x *Room) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Room) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type DummyListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DummyListRequest) Reset() {
	*x = DummyListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_room_room_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DummyListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DummyListRequest) ProtoMessage() {}

func (x *DummyListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_room_room_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DummyListRequest.ProtoReflect.Descriptor instead.
func (*DummyListRequest) Descriptor() ([]byte, []int) {
	return file_proto_room_room_proto_rawDescGZIP(), []int{2}
}

type DummyListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rooms []*Room `protobuf:"bytes,1,rep,name=rooms,proto3" json:"rooms,omitempty"`
}

func (x *DummyListResponse) Reset() {
	*x = DummyListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_room_room_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DummyListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DummyListResponse) ProtoMessage() {}

func (x *DummyListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_room_room_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DummyListResponse.ProtoReflect.Descriptor instead.
func (*DummyListResponse) Descriptor() ([]byte, []int) {
	return file_proto_room_room_proto_rawDescGZIP(), []int{3}
}

func (x *DummyListResponse) GetRooms() []*Room {
	if x != nil {
		return x.Rooms
	}
	return nil
}

type CreateReqeust struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Room *NewRoom `protobuf:"bytes,1,opt,name=room,proto3" json:"room,omitempty"`
}

func (x *CreateReqeust) Reset() {
	*x = CreateReqeust{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_room_room_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateReqeust) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateReqeust) ProtoMessage() {}

func (x *CreateReqeust) ProtoReflect() protoreflect.Message {
	mi := &file_proto_room_room_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateReqeust.ProtoReflect.Descriptor instead.
func (*CreateReqeust) Descriptor() ([]byte, []int) {
	return file_proto_room_room_proto_rawDescGZIP(), []int{4}
}

func (x *CreateReqeust) GetRoom() *NewRoom {
	if x != nil {
		return x.Room
	}
	return nil
}

type CreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Room *Room `protobuf:"bytes,1,opt,name=room,proto3" json:"room,omitempty"`
}

func (x *CreateResponse) Reset() {
	*x = CreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_room_room_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateResponse) ProtoMessage() {}

func (x *CreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_room_room_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateResponse.ProtoReflect.Descriptor instead.
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return file_proto_room_room_proto_rawDescGZIP(), []int{5}
}

func (x *CreateResponse) GetRoom() *Room {
	if x != nil {
		return x.Room
	}
	return nil
}

var File_proto_room_room_proto protoreflect.FileDescriptor

var file_proto_room_room_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x6f, 0x6f, 0x6d, 0x2f, 0x72, 0x6f, 0x6f,
	0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x22, 0x1d, 0x0a,
	0x07, 0x4e, 0x65, 0x77, 0x52, 0x6f, 0x6f, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x2a, 0x0a, 0x04,
	0x52, 0x6f, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x12, 0x0a, 0x10, 0x44, 0x75, 0x6d, 0x6d,
	0x79, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x35, 0x0a, 0x11,
	0x44, 0x75, 0x6d, 0x6d, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x20, 0x0a, 0x05, 0x72, 0x6f, 0x6f, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0a, 0x2e, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x05, 0x72, 0x6f,
	0x6f, 0x6d, 0x73, 0x22, 0x32, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x65, 0x75, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x4e, 0x65, 0x77, 0x52, 0x6f, 0x6f,
	0x6d, 0x52, 0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x22, 0x30, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x04, 0x72, 0x6f, 0x6f,
	0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x52,
	0x6f, 0x6f, 0x6d, 0x52, 0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x32, 0x84, 0x01, 0x0a, 0x0b, 0x52, 0x6f,
	0x6f, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3e, 0x0a, 0x09, 0x44, 0x75, 0x6d,
	0x6d, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x16, 0x2e, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x44, 0x75,
	0x6d, 0x6d, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17,
	0x2e, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x44, 0x75, 0x6d, 0x6d, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x06, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x12, 0x13, 0x2e, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x65, 0x75, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x72, 0x6f, 0x6f, 0x6d, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x30, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68,
	0x73, 0x6d, 0x74, 0x6b, 0x6b, 0x2f, 0x6c, 0x69, 0x74, 0x65, 0x72, 0x61, 0x74, 0x65, 0x2d, 0x62,
	0x61, 0x72, 0x6e, 0x61, 0x63, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x6f,
	0x6f, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_room_room_proto_rawDescOnce sync.Once
	file_proto_room_room_proto_rawDescData = file_proto_room_room_proto_rawDesc
)

func file_proto_room_room_proto_rawDescGZIP() []byte {
	file_proto_room_room_proto_rawDescOnce.Do(func() {
		file_proto_room_room_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_room_room_proto_rawDescData)
	})
	return file_proto_room_room_proto_rawDescData
}

var file_proto_room_room_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_room_room_proto_goTypes = []interface{}{
	(*NewRoom)(nil),           // 0: room.NewRoom
	(*Room)(nil),              // 1: room.Room
	(*DummyListRequest)(nil),  // 2: room.DummyListRequest
	(*DummyListResponse)(nil), // 3: room.DummyListResponse
	(*CreateReqeust)(nil),     // 4: room.CreateReqeust
	(*CreateResponse)(nil),    // 5: room.CreateResponse
}
var file_proto_room_room_proto_depIdxs = []int32{
	1, // 0: room.DummyListResponse.rooms:type_name -> room.Room
	0, // 1: room.CreateReqeust.room:type_name -> room.NewRoom
	1, // 2: room.CreateResponse.room:type_name -> room.Room
	2, // 3: room.RoomService.DummyList:input_type -> room.DummyListRequest
	4, // 4: room.RoomService.Create:input_type -> room.CreateReqeust
	3, // 5: room.RoomService.DummyList:output_type -> room.DummyListResponse
	5, // 6: room.RoomService.Create:output_type -> room.CreateResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_room_room_proto_init() }
func file_proto_room_room_proto_init() {
	if File_proto_room_room_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_room_room_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewRoom); i {
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
		file_proto_room_room_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Room); i {
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
		file_proto_room_room_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DummyListRequest); i {
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
		file_proto_room_room_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DummyListResponse); i {
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
		file_proto_room_room_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateReqeust); i {
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
		file_proto_room_room_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateResponse); i {
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
			RawDescriptor: file_proto_room_room_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_room_room_proto_goTypes,
		DependencyIndexes: file_proto_room_room_proto_depIdxs,
		MessageInfos:      file_proto_room_room_proto_msgTypes,
	}.Build()
	File_proto_room_room_proto = out.File
	file_proto_room_room_proto_rawDesc = nil
	file_proto_room_room_proto_goTypes = nil
	file_proto_room_room_proto_depIdxs = nil
}