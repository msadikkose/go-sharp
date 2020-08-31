package tests

import (
	"context"
	"go-sharp/business/manager"
	"go-sharp/dal/repository"
	"go-sharp/dal/uow"
	"go-sharp/ioc"
	"sync"
)


type testKernel struct{}

func (k *testKernel) GetBusinessManagerFactory(ctx context.Context) *manager.Factory {


	dbHandler:=&TestDbHandler{}
	client := dbHandler.GetDbClient()

	tx, err := client.Tx(ctx)
	if err != nil{

	}
	userRepository := &repository.User{repository.Base{Transaction: tx, Ctx: ctx}}

	unitOfWork := &uow.UnitOfWork{UserRepository: userRepository, Transaction: tx}

	userManager := &manager.User{Uow: unitOfWork}
	businessManagerFactory := &manager.Factory{Uow: unitOfWork, UserManager: userManager}

	return businessManagerFactory
}

var (
	k             *testKernel
	containerOnce sync.Once
)



func TestServiceContainer() ioc.IServiceContainer {
	if k == nil {
		containerOnce.Do(func() {
			k = &testKernel{}
		})
	}
	return k
}
