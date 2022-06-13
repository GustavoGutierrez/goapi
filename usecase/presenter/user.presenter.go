package presenter

import "github.com/GustavoGutierrez/goapi/domain/model"

type UserPresenter interface {
	ResponseUsers(u []*model.User) []*model.User
}
