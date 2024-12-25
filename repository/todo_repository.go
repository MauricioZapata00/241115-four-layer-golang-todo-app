package repository

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"four-layer-todo-app/model"
)

type TodoRepository interface {
	Create(ctx context.Context, todo *model.Todo) error
	FindAll(ctx context.Context) ([]model.Todo, error)
	FindByTitle(ctx context.Context, title string) (*model.Todo, error)
}

type mongoTodoRepository struct {
	db *mongo.Database
}

func NewTodoRepository(db *mongo.Database) TodoRepository {
	return &mongoTodoRepository{db: db}
}

func (r *mongoTodoRepository) Create(ctx context.Context, todo *model.Todo) error {
	_, err := r.db.Collection("todos").InsertOne(ctx, todo)
	return err
}

func (r *mongoTodoRepository) FindAll(ctx context.Context) ([]model.Todo, error) {
	var todos []model.Todo
	cursor, err := r.db.Collection("todos").Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			log.Panic(err)
		}
	}(cursor, ctx)

	if err = cursor.All(ctx, &todos); err != nil {
		return nil, err
	}
	return todos, nil
}

func (r *mongoTodoRepository) FindByTitle(ctx context.Context, title string) (*model.Todo, error) {
	todo := new(model.Todo)
	err := r.db.Collection("todos").FindOne(ctx, bson.M{"title": title}).Decode(todo)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return todo, nil
}
