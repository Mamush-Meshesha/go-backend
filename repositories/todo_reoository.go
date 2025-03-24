package repositories

import (
	"todo/models"

	"gorm.io/gorm"
)

type TodoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) * TodoRepository {
	return & TodoRepository{db: db}
}

func (r *TodoRepository) Create(todo *models.Todo) error {
	return r.db.Create(todo).Error
}

func (r *TodoRepository) FindAll() ([]models.Todo, error) {
	var todos []models.Todo
	err := r.db.Find(&todos).Error
	return todos, err
}

func (r *TodoRepository) FindByID(id uint) (*models.Todo, error){
	var todo models.Todo
	err := r.db.First(&todo, id).Error
	return &todo,err
}

func (r *TodoRepository) Update(todo *models.Todo) error {
	return r.db.Save(todo).Error
}

func (r *TodoRepository) Delete(id uint) error {
	return r.db.Delete(&models.Todo{},id).Error
}