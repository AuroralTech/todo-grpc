package usecase

import (
	"github.com/AuroralTech/todo-grpc/pkg/domain/model"
	"github.com/AuroralTech/todo-grpc/pkg/domain/repository"
)

type TodoUsecase struct {
	repo repository.TodoRepository
}

func NewTodoUsecase(repo repository.TodoRepository) *TodoUsecase {
	return &TodoUsecase{repo}
}

func (u *TodoUsecase) AddTodo(todo *model.Todo) (*model.Todo, error) {
	return u.repo.AddTodo(todo)
}

func (u *TodoUsecase) UpdateTodoStatus(id string, isCompleted bool) (bool, error) {
	return u.repo.UpdateTodoStatus(id, isCompleted)
}

func (u *TodoUsecase) DeleteTodoById(id string) (bool, error) {
	return u.repo.DeleteTodoById(id)
}

func (u *TodoUsecase) GetTodoList() ([]*model.Todo, error) {
	return u.repo.GetTodoList()
}
