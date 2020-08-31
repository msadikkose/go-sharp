package ioc

import (
	"context"
	"go-sharp/business/manager"
	"go-sharp/dal/db"
	"go-sharp/dal/repository"
	"go-sharp/dal/uow"
	_ "github.com/lib/pq"
	"sync"
)

const (
	hostName     = "localhost"
	hostPort     = 5432
	username     = "postgres"
	password     = "msk161893"
	databaseName = "postgres"
)

type IServiceContainer interface {
	GetBusinessManagerFactory(ctx context.Context) *manager.Factory
}

type kernel struct{}

func (k *kernel) GetBusinessManagerFactory(ctx context.Context) *manager.Factory {


	dbHandler:=&db.DbHandler{}
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
	k             *kernel
	containerOnce sync.Once
)

func ServiceContainer() IServiceContainer {
	if k == nil {
		containerOnce.Do(func() {
			k = &kernel{}
		})
	}
	return k
}
