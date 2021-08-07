package service

import (
	"gin-app/model"
)

type TodoService struct {}

func (TodoService) GetTodoList() []model.Todo {
	tests := make([]model.Todo, 0)
	return tests
}

func (TodoService) SetTodo(todo *model.Todo) error {
	DbConnection.Create(todo)
	// if err != nil{
	// 	return err
	// }
	return nil
}