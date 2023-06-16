package service

import (
	"github.com/drakenchef/backproject"
	"github.com/drakenchef/backproject/pkg/repository"
)

type Authorizarion interface {
	CreateUser(user backproject.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(userId int, list backproject.TodoList) (int, error)
	GetAll(userdId int) ([]backproject.TodoList, error)
	GetById(userId, listId int) (backproject.TodoList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input backproject.UpdateListInput) error
}

type TodoItem interface {
	Create(userId, listId int, item backproject.TodoItem) (int, error)
	GetAll(userId, listId int) ([]backproject.TodoItem, error)
	GetById(userId, itemId int) (backproject.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input backproject.UpdateItemInput) error
}

type Service struct {
	Authorizarion
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorizarion: NewAuthService(repos.Authorizarion),
		TodoList:      NewTodoListService(repos.TodoList),
		TodoItem:      NewTodoItemService(repos.TodoItem, repos.TodoList),
	}
}
