package repository

import (
	"todo-go/database"
	"todo-go/internals/models"

	"gorm.io/gorm"
)

type TodoRepository struct {
	DB *gorm.DB
}

func NewTodoRepository() *TodoRepository {
	return &TodoRepository{DB: database.DB}
}

func (r *TodoRepository) GetAll() ([]models.Todo, error) {
	var todos []models.Todo
	err := r.DB.Find(&todos).Error
	return todos, err
}

func (r *TodoRepository) GetOne(id int) (*models.Todo, error) {
	var todo models.Todo
	err := r.DB.First(&todo, id).Error
	return &todo, err
}

func (r *TodoRepository) Create(todo *models.Todo) error {
	return r.DB.Create(todo).Error
}

func (r *TodoRepository) Update(todo *models.Todo) error {
	return r.DB.Save(todo).Error
}

func (r *TodoRepository) Delete(id int) error {
	return r.DB.Delete(&models.Todo{}, id).Error
}
