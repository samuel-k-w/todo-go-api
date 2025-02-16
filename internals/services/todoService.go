package services

import (
	"todo-go/internals/models"
	"todo-go/internals/repository"
)

type TodoService struct {
	TodoRepo *repository.TodoRepository
}

func NewTodoService() *TodoService {
	return &TodoService{TodoRepo: repository.NewTodoRepository()}
}

func (t *TodoService) Get(id int) (*models.Todo, error) {
	todo, err := t.TodoRepo.GetOne(id)
	return todo, err
}

func (t *TodoService) GetAllTodos() ([]models.Todo, error) {
	return t.TodoRepo.GetAll()
}

func (t *TodoService) CreateTodo(todo *models.Todo) error {
	return t.TodoRepo.Create(todo)
}

func (t *TodoService) UpdateTodo(todo *models.Todo) error {
	return t.TodoRepo.Update(todo)
}

func (t *TodoService) Delete(id int) error {
	return t.TodoRepo.Delete(id)
}
