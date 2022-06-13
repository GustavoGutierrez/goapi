package repository

import "github.com/GustavoGutierrez/goapi/domain/model"

type UserRepository interface {
	FindAll(u []*model.User) ([]*model.User, error)
}
