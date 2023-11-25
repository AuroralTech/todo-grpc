package repository

import "github.com/AuroralTech/todo-grpc/pkg/domain/model"

type TodoRepository interface {
	AddTodo(todo *model.Todo) (*model.Todo, error)
	UpdateTodoStatus(id string, isCompleted bool) (bool, error)
	DeleteTodoById(id string) (bool, error)
	GetTodoList() ([]*model.Todo, error)
}
