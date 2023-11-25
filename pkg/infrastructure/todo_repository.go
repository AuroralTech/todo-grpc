package infrastructure

import (
	"github.com/AuroralTech/todo-grpc/pkg/domain/model"
	"github.com/AuroralTech/todo-grpc/pkg/domain/repository"
	"gorm.io/gorm"
)

type todoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) repository.TodoRepository {
	return &todoRepository{db}
}

func (r *todoRepository) AddTodo(todo *model.Todo) (*model.Todo, error) {
	if err := r.db.Create(todo).Error; err != nil {
		return nil, err
	}
	return todo, nil
}

func (r *todoRepository) UpdateTodoStatus(id string, isCompleted bool) (bool, error) {
	todo := &model.Todo{}
	if err := r.db.First(todo, "id = ?", id).Error; err != nil {
		return false, err
	}
	todo.IsCompleted = isCompleted
	if err := r.db.Save(todo).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (r *todoRepository) DeleteTodoById(id string) (bool, error) {
	if err := r.db.Delete(&model.Todo{}, "id = ?", id).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (r *todoRepository) GetTodoList() ([]*model.Todo, error) {
	var todos []*model.Todo
	if err := r.db.Find(&todos).Error; err != nil {
		return nil, err
	}
	return todos, nil
}
