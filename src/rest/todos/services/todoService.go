package services

import (
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"rest-api/src/models"
)

type TodoService struct{}

func (todoService TodoService) Create(todo *models.Todo) error {
	err := mgm.Coll(todo).Create(todo)
	if err != nil {
		return err
	}
	return nil
}

func (todoService TodoService) FindTodo(id string) (*models.Todo, error) {
	foundTodo := &models.Todo{}
	err := mgm.Coll(foundTodo).FindByID(id, foundTodo)
	if err != nil {
		return nil, err
	}

	return foundTodo, nil
}

func (todoService TodoService) ToggleCompleted(id string) (*models.Todo, error) {
	todo, err := todoService.FindTodo(id)
	if err != nil {
		return nil, err
	}

	todo.Completed = !todo.Completed
	mongoErr := mgm.Coll(todo).Update(todo)
	if mongoErr != nil {
		return nil, mongoErr
	}

	return todo, nil
}

func (todoService TodoService) GetTodos(user string) ([]models.Todo, error) {
	var results []models.Todo
	err := mgm.Coll(&models.Todo{}).SimpleFind(&results, bson.M{"user": user})
	if err != nil {
		return nil, err
	}
	return results, nil
}
