package grpc

import (
	"context"

	"github.com/AuroralTech/todo-grpc/pkg/domain/model"
	pb "github.com/AuroralTech/todo-grpc/pkg/grpc/generated"
	"github.com/AuroralTech/todo-grpc/pkg/usecase"
	"google.golang.org/protobuf/types/known/emptypb"
)

type TodoHandler struct {
	usecase usecase.TodoUsecase
}

func NewTodoHandler(usecase usecase.TodoUsecase) pb.TodoServiceServer {
	return &TodoHandler{usecase: usecase}
}

func (h *TodoHandler) AddTodo(ctx context.Context, req *pb.TodoItem) (*pb.TodoItem, error) {
	if err := Authenticate(ctx); err != nil {
		return nil, err
	}
	todo := &model.Todo{Task: req.GetTask(), IsCompleted: req.GetIsCompleted()}
	result, err := h.usecase.AddTodo(todo)
	if err != nil {
		return nil, err
	}
	return &pb.TodoItem{Id: uint64(result.ID), Task: result.Task, IsCompleted: result.IsCompleted}, nil
}

func (h *TodoHandler) UpdateTodoStatus(ctx context.Context, req *pb.UpdateTodoStatusRequest) (*pb.UpdateTodoStatusResponse, error) {
	success, err := h.usecase.UpdateTodoStatus(req.GetId(), req.GetIsCompleted())
	if err != nil {
		return nil, err
	}
	return &pb.UpdateTodoStatusResponse{Success: success}, nil
}

func (h *TodoHandler) DeleteTodoById(ctx context.Context, req *pb.DeleteTodoByIdRequest) (*pb.DeleteTodoByIdResponse, error) {
	success, err := h.usecase.DeleteTodoById(req.GetId())
	if err != nil {
		return nil, err
	}
	return &pb.DeleteTodoByIdResponse{Success: success}, nil
}

func (h *TodoHandler) GetTodoList(ctx context.Context, req *emptypb.Empty) (*pb.TodoList, error) {
	todos, err := h.usecase.GetTodoList()
	if err != nil {
		return nil, err
	}
	var todoItems []*pb.TodoItem
	for _, todo := range todos {
		todoItems = append(todoItems, &pb.TodoItem{Id: uint64(todo.ID), Task: todo.Task, IsCompleted: todo.IsCompleted})
	}
	return &pb.TodoList{Items: todoItems}, nil
}
