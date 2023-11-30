# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [pkg/proto/todo.proto](#pkg_proto_todo-proto)
    - [DeleteTodoByIdRequest](#todo-DeleteTodoByIdRequest)
    - [DeleteTodoByIdResponse](#todo-DeleteTodoByIdResponse)
    - [TodoItem](#todo-TodoItem)
    - [TodoList](#todo-TodoList)
    - [UpdateTodoStatusRequest](#todo-UpdateTodoStatusRequest)
    - [UpdateTodoStatusResponse](#todo-UpdateTodoStatusResponse)
  
    - [TodoService](#todo-TodoService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="pkg_proto_todo-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## pkg/proto/todo.proto



<a name="todo-DeleteTodoByIdRequest"></a>

### DeleteTodoByIdRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | 削除するタスクのID |






<a name="todo-DeleteTodoByIdResponse"></a>

### DeleteTodoByIdResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| success | [bool](#bool) |  | 削除が成功したかどうか |






<a name="todo-TodoItem"></a>

### TodoItem
TodoItemは一つのタスクを表します。


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  | タスクの一意なID |
| task | [string](#string) |  | タスクの内容 |
| is_completed | [bool](#bool) |  | タスクが完了したかどうか |






<a name="todo-TodoList"></a>

### TodoList
TodoListは複数のタスクを含むリストを表します。


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| items | [TodoItem](#todo-TodoItem) | repeated | TodoItemのリスト |






<a name="todo-UpdateTodoStatusRequest"></a>

### UpdateTodoStatusRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | 更新するタスクのID |
| is_completed | [bool](#bool) |  | タスクの新しいステータス |






<a name="todo-UpdateTodoStatusResponse"></a>

### UpdateTodoStatusResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| success | [bool](#bool) |  | 更新が成功したかどうか |





 

 

 


<a name="todo-TodoService"></a>

### TodoService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| AddTodo | [TodoItem](#todo-TodoItem) | [TodoItem](#todo-TodoItem) |  |
| UpdateTodoStatus | [UpdateTodoStatusRequest](#todo-UpdateTodoStatusRequest) | [UpdateTodoStatusResponse](#todo-UpdateTodoStatusResponse) |  |
| DeleteTodoById | [DeleteTodoByIdRequest](#todo-DeleteTodoByIdRequest) | [DeleteTodoByIdResponse](#todo-DeleteTodoByIdResponse) |  |
| GetTodoList | [.google.protobuf.Empty](#google-protobuf-Empty) | [TodoList](#todo-TodoList) |  |

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

