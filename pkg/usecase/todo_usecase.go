package usecase

import (
	"github.com/AuroralTech/todo-grpc/pkg/domain/model"
	"github.com/AuroralTech/todo-grpc/pkg/domain/repository"
)

type TodoUsecase interface {
	AddTodo(todo *model.Todo) (*model.Todo, error)
	UpdateTodoStatus(id string, isCompleted bool) (bool, error)
	DeleteTodoById(id string) (bool, error)
	GetTodoList() ([]*model.Todo, error)
}
type todoUsecase struct {
	repo repository.TodoRepository
}

func NewTodoUsecase(repo repository.TodoRepository) TodoUsecase {
	return &todoUsecase{repo}
}

func (u *todoUsecase) AddTodo(todo *model.Todo) (*model.Todo, error) {
	return u.repo.AddTodo(todo)
}

func (u *todoUsecase) UpdateTodoStatus(id string, isCompleted bool) (bool, error) {
	return u.repo.UpdateTodoStatus(id, isCompleted)
}

func (u *todoUsecase) DeleteTodoById(id string) (bool, error) {
	return u.repo.DeleteTodoById(id)
}

func (u *todoUsecase) GetTodoList() ([]*model.Todo, error) {
	return u.repo.GetTodoList()
}
