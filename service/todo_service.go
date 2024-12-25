package service

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"errors"

	"four-layer-todo-app/model"
	"four-layer-todo-app/repository"
)

type TodoService interface {
	CreateTodo(ctx context.Context, todo *model.Todo) (*model.Todo, error)
	GetAllTodos(ctx context.Context) ([]model.Todo, error)
}

type todoService struct {
	repo repository.TodoRepository
}

func NewTodoService(repo repository.TodoRepository) TodoService {
	return &todoService{repo: repo}
}

func (s *todoService) CreateTodo(ctx context.Context, todo *model.Todo) (*model.Todo, error) {
	// Validate title
	if todo.Title == "" || len(todo.Title) < 3 {
		return nil, errors.New("title must be at least 3 characters long")
	}

	// Check for existing document with same title
	existingTodo, err := s.repo.FindByTitle(ctx, todo.Title)
	if err != nil {
		return nil, err
	}
	if existingTodo != nil {
		return nil, errors.New("todo with this title already exists")
	}

	// Generate ID
	todo.ID, err = generateID()
	if err != nil {
		return nil, err
	}

	return todo, s.repo.Create(ctx, todo)
}

func (s *todoService) GetAllTodos(ctx context.Context) ([]model.Todo, error) {
	return s.repo.FindAll(ctx)
}

func generateID() (string, error) {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}
