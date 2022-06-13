package registry

import (
	"github.com/GustavoGutierrez/goapi/interface/controllers"
	ip "github.com/GustavoGutierrez/goapi/interface/presenters"
	ir "github.com/GustavoGutierrez/goapi/interface/repository"
	"github.com/GustavoGutierrez/goapi/usecase/interactor"
	up "github.com/GustavoGutierrez/goapi/usecase/presenter"
	ur "github.com/GustavoGutierrez/goapi/usecase/repository"
)

func (r *registry) NewUserController() controllers.UserController {
	return controllers.NewUserController(r.NewUserInteractor())
}

func (r *registry) NewUserInteractor() interactor.UserInteractor {
	return interactor.NewUserInteractor(r.NewUserRepository(), r.NewUserPresenter())
}

func (r *registry) NewUserRepository() ur.UserRepository {
	return ir.NewUserRepository(r.db)
}

func (r *registry) NewUserPresenter() up.UserPresenter {
	return ip.NewUserPresenter()
}
