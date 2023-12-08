// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: pkg/proto/todo.proto

package generated

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// TodoItemは一つのタスクを表します。
type TodoItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// タスクの一意なID
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// タスクの内容
	Task string `protobuf:"bytes,2,opt,name=task,proto3" json:"task,omitempty"`
	// タスクが完了したかどうか
	IsCompleted bool `protobuf:"varint,3,opt,name=is_completed,json=isCompleted,proto3" json:"is_completed,omitempty"`
}

func (x *TodoItem) Reset() {
	*x = TodoItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_todo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TodoItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TodoItem) ProtoMessage() {}

func (x *TodoItem) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_todo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TodoItem.ProtoReflect.Descriptor instead.
func (*TodoItem) Descriptor() ([]byte, []int) {
	return file_pkg_proto_todo_proto_rawDescGZIP(), []int{0}
}

func (x *TodoItem) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TodoItem) GetTask() string {
	if x != nil {
		return x.Task
	}
	return ""
}

func (x *TodoItem) GetIsCompleted() bool {
	if x != nil {
		return x.IsCompleted
	}
	return false
}

// TodoListは複数のタスクを含むリストを表します。
type TodoList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// TodoItemのリスト
	Items []*TodoItem `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *TodoList) Reset() {
	*x = TodoList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_todo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TodoList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TodoList) ProtoMessage() {}

func (x *TodoList) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_todo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TodoList.ProtoReflect.Descriptor instead.
func (*TodoList) Descriptor() ([]byte, []int) {
	return file_pkg_proto_todo_proto_rawDescGZIP(), []int{1}
}

func (x *TodoList) GetItems() []*TodoItem {
	if x != nil {
		return x.Items
	}
	return nil
}

type UpdateTodoStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`                                       // 更新するタスクのID
	IsCompleted bool   `protobuf:"varint,2,opt,name=is_completed,json=isCompleted,proto3" json:"is_completed,omitempty"` // タスクの新しいステータス
}

func (x *UpdateTodoStatusRequest) Reset() {
	*x = UpdateTodoStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_todo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTodoStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTodoStatusRequest) ProtoMessage() {}

func (x *UpdateTodoStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_todo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTodoStatusRequest.ProtoReflect.Descriptor instead.
func (*UpdateTodoStatusRequest) Descriptor() ([]byte, []int) {
	return file_pkg_proto_todo_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateTodoStatusRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateTodoStatusRequest) GetIsCompleted() bool {
	if x != nil {
		return x.IsCompleted
	}
	return false
}

type UpdateTodoStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"` // 更新が成功したかどうか
}

func (x *UpdateTodoStatusResponse) Reset() {
	*x = UpdateTodoStatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_todo_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTodoStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTodoStatusResponse) ProtoMessage() {}

func (x *UpdateTodoStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_todo_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTodoStatusResponse.ProtoReflect.Descriptor instead.
func (*UpdateTodoStatusResponse) Descriptor() ([]byte, []int) {
	return file_pkg_proto_todo_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateTodoStatusResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type DeleteTodoByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"` // 削除するタスクのID
}

func (x *DeleteTodoByIdRequest) Reset() {
	*x = DeleteTodoByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_todo_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTodoByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTodoByIdRequest) ProtoMessage() {}

func (x *DeleteTodoByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_todo_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTodoByIdRequest.ProtoReflect.Descriptor instead.
func (*DeleteTodoByIdRequest) Descriptor() ([]byte, []int) {
	return file_pkg_proto_todo_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteTodoByIdRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteTodoByIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"` // 削除が成功したかどうか
}

func (x *DeleteTodoByIdResponse) Reset() {
	*x = DeleteTodoByIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_todo_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTodoByIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTodoByIdResponse) ProtoMessage() {}

func (x *DeleteTodoByIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_todo_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTodoByIdResponse.ProtoReflect.Descriptor instead.
func (*DeleteTodoByIdResponse) Descriptor() ([]byte, []int) {
	return file_pkg_proto_todo_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteTodoByIdResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_pkg_proto_todo_proto protoreflect.FileDescriptor

var file_pkg_proto_todo_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x6f, 0x64, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x74, 0x6f, 0x64, 0x6f, 0x1a, 0x1b, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d,
	0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x5a, 0x0a, 0x08, 0x54, 0x6f, 0x64, 0x6f, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b,
	0x0a, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42,
	0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x12, 0x21, 0x0a, 0x0c, 0x69,
	0x73, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0b, 0x69, 0x73, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x22, 0x30,
	0x0a, 0x08, 0x54, 0x6f, 0x64, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x05, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x74, 0x6f, 0x64, 0x6f,
	0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x22, 0x4c, 0x0a, 0x17, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x69,
	0x73, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0b, 0x69, 0x73, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x22, 0x34,
	0x0a, 0x18, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x22, 0x27, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x6f,
	0x64, 0x6f, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x32, 0x0a,
	0x16, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x42, 0x79, 0x49, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x32, 0x8f, 0x02, 0x0a, 0x0b, 0x54, 0x6f, 0x64, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x29, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x54, 0x6f, 0x64, 0x6f, 0x12, 0x0e, 0x2e, 0x74,
	0x6f, 0x64, 0x6f, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x49, 0x74, 0x65, 0x6d, 0x1a, 0x0e, 0x2e, 0x74,
	0x6f, 0x64, 0x6f, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x51, 0x0a, 0x10,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x1d, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f,
	0x64, 0x6f, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1e, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x64,
	0x6f, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x4b, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x42, 0x79, 0x49,
	0x64, 0x12, 0x1b, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54,
	0x6f, 0x64, 0x6f, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c,
	0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x6f, 0x64, 0x6f,
	0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x0b,
	0x47, 0x65, 0x74, 0x54, 0x6f, 0x64, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x0e, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x4c,
	0x69, 0x73, 0x74, 0x42, 0x14, 0x5a, 0x12, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f,
	0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_pkg_proto_todo_proto_rawDescOnce sync.Once
	file_pkg_proto_todo_proto_rawDescData = file_pkg_proto_todo_proto_rawDesc
)

func file_pkg_proto_todo_proto_rawDescGZIP() []byte {
	file_pkg_proto_todo_proto_rawDescOnce.Do(func() {
		file_pkg_proto_todo_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_proto_todo_proto_rawDescData)
	})
	return file_pkg_proto_todo_proto_rawDescData
}

var file_pkg_proto_todo_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_pkg_proto_todo_proto_goTypes = []interface{}{
	(*TodoItem)(nil),                 // 0: todo.TodoItem
	(*TodoList)(nil),                 // 1: todo.TodoList
	(*UpdateTodoStatusRequest)(nil),  // 2: todo.UpdateTodoStatusRequest
	(*UpdateTodoStatusResponse)(nil), // 3: todo.UpdateTodoStatusResponse
	(*DeleteTodoByIdRequest)(nil),    // 4: todo.DeleteTodoByIdRequest
	(*DeleteTodoByIdResponse)(nil),   // 5: todo.DeleteTodoByIdResponse
	(*emptypb.Empty)(nil),            // 6: google.protobuf.Empty
}
var file_pkg_proto_todo_proto_depIdxs = []int32{
	0, // 0: todo.TodoList.items:type_name -> todo.TodoItem
	0, // 1: todo.TodoService.AddTodo:input_type -> todo.TodoItem
	2, // 2: todo.TodoService.UpdateTodoStatus:input_type -> todo.UpdateTodoStatusRequest
	4, // 3: todo.TodoService.DeleteTodoById:input_type -> todo.DeleteTodoByIdRequest
	6, // 4: todo.TodoService.GetTodoList:input_type -> google.protobuf.Empty
	0, // 5: todo.TodoService.AddTodo:output_type -> todo.TodoItem
	3, // 6: todo.TodoService.UpdateTodoStatus:output_type -> todo.UpdateTodoStatusResponse
	5, // 7: todo.TodoService.DeleteTodoById:output_type -> todo.DeleteTodoByIdResponse
	1, // 8: todo.TodoService.GetTodoList:output_type -> todo.TodoList
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pkg_proto_todo_proto_init() }
func file_pkg_proto_todo_proto_init() {
	if File_pkg_proto_todo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_proto_todo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TodoItem); i {
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
		file_pkg_proto_todo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TodoList); i {
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
		file_pkg_proto_todo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTodoStatusRequest); i {
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
		file_pkg_proto_todo_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTodoStatusResponse); i {
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
		file_pkg_proto_todo_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteTodoByIdRequest); i {
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
		file_pkg_proto_todo_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteTodoByIdResponse); i {
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
			RawDescriptor: file_pkg_proto_todo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_proto_todo_proto_goTypes,
		DependencyIndexes: file_pkg_proto_todo_proto_depIdxs,
		MessageInfos:      file_pkg_proto_todo_proto_msgTypes,
	}.Build()
	File_pkg_proto_todo_proto = out.File
	file_pkg_proto_todo_proto_rawDesc = nil
	file_pkg_proto_todo_proto_goTypes = nil
	file_pkg_proto_todo_proto_depIdxs = nil
}
