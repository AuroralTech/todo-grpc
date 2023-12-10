package usecase

import (
	"fmt"

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
	todo, err := u.repo.AddTodo(todo)
	if err != nil {
		return nil, fmt.Errorf("TodoUsecase: failed to AddTodo: %v", err)
	}
	return todo, nil
}

func (u *todoUsecase) UpdateTodoStatus(id string, isCompleted bool) (bool, error) {
	res, err := u.repo.UpdateTodoStatus(id, isCompleted)
	if err != nil {
		return false, fmt.Errorf("TodoUsecase: failed to UpdateTodoStatus: %v", err)
	}
	return res, nil
}

func (u *todoUsecase) DeleteTodoById(id string) (bool, error) {
	res, err := u.repo.DeleteTodoById(id)
	if err != nil {
		return false, fmt.Errorf("TodoUsecase: failed to DeleteTodoById: %v", err)
	}
	return res, nil
}

func (u *todoUsecase) GetTodoList() ([]*model.Todo, error) {
	todoList, err := u.repo.GetTodoList()
	if err != nil {
		return nil, fmt.Errorf("TodoUsecase: failed to GetTodoList: %v", err)
	}
	return todoList, nil
}
