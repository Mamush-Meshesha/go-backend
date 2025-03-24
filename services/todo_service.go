package services

import (
	"todo/models"
	"todo/repositories"
)

type TodoService struct {
	repo *repositories.TodoRepository
}

func NewTodoService(repo *repositories.TodoRepository) *TodoService {
	return &TodoService{repo: repo}
}

func (s *TodoService) CreateTodo(todo *models.Todo) error {
	return s.repo.Create(todo)
}

func (s *TodoService) GetAllTodos() ([]models.Todo,error) {
	return s.repo.FindAll()
}

func (s *TodoService) GetTodoByID(id uint) (*models.Todo, error) {
	return s.repo.FindByID(id)
}

func (s *TodoService) UpdateTodo(todo *models.Todo) error {{
	return s.repo.Update(todo)
}}

func (s *TodoService) DeleteTodo(id uint) error {
	return s.repo.Delete(id)
}