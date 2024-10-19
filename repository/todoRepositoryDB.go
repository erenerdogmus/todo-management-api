package repository

import (
	"github.com/erenerdogmus/models"
	"context"
	"errors"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

//go:generate mockgen -destination=../mock/repository/mockTodoRepository.go -package=repository github.com/erenerdogmus/repository TodoRepository
type TodoRepositoryDB struct {
	TodoCollection *mongo.Collection
}

type TodoRepository interface {
	Insert(todo models.Todo) (bool, error)
	GetAll() ([]models.Todo, error)
	Delete(id primitive.ObjectID) (bool, error)
}

func (t TodoRepositoryDB) Insert(todo models.Todo) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	todo.ID = primitive.NewObjectID()
	result, err := t.TodoCollection.InsertOne(ctx, todo)
	if result.InsertedID == nil || err != nil {
		_ = errors.New("failed add")
		return false, err
	}

	return true, nil
}

func (t TodoRepositoryDB) GetAll() ([]models.Todo, error) {
	var todos []models.Todo

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := t.TodoCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	defer result.Close(ctx)

	for result.Next(ctx) {
		var todo models.Todo
		if err := result.Decode(&todo); err != nil {
			log.Fatalln(err)
			return nil, err
		}
		todos = append(todos, todo)
	}

	if err := result.Err(); err != nil {
		return nil, err
	}

	return todos, nil
}

func (t TodoRepositoryDB) Delete(id primitive.ObjectID) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := t.TodoCollection.DeleteOne(ctx, bson.M{"id": id})
	if err != nil || result.DeletedCount <= 0 {
		return false, err
	}
	return true, nil
}

func NewTodoRepositoryDb(dbClient *mongo.Collection) TodoRepositoryDB {
	return TodoRepositoryDB{TodoCollection: dbClient}
}
