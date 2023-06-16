package repository

import (
	"github.com/drakenchef/backproject"
	"github.com/jmoiron/sqlx"
)

type Authorizarion interface {
	CreateUser(user backproject.User) (int, error)
	GetUser(username, password string) (backproject.User, error)
}

type TodoList interface {
	Create(userId int, list backproject.TodoList) (int, error)
	GetAll(userId int) ([]backproject.TodoList, error)
	GetById(userId, listId int) (backproject.TodoList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input backproject.UpdateListInput) error
}

type TodoItem interface {
	Create(listId int, item backproject.TodoItem) (int, error)
	GetAll(userId, listId int) ([]backproject.TodoItem, error)
	GetById(userId, itemId int) (backproject.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input backproject.UpdateItemInput) error
}

type Repository struct {
	Authorizarion
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorizarion: NewAuthPostgres(db),
		TodoList:      NewTodoListPostgres(db),
		TodoItem:      NewTodoItemPostgres(db),
	}
}
