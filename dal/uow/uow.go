package uow

import (
	"go-sharp/dal/repository"
	"go-sharp/ent"
)

type IUnitOfWork interface {
	GetUserRepository() repository.IUser
	CommitTransaction() error
	RollbackTransaction() error
}

type UnitOfWork struct {
	UserRepository repository.IUser
	Transaction *ent.Tx
}

func (uow *UnitOfWork) GetUserRepository() repository.IUser {
	return uow.UserRepository
}
func (uow *UnitOfWork) CommitTransaction() error {
	return uow.Transaction.Commit()
}
func (uow *UnitOfWork) RollbackTransaction() error {
	return uow.Transaction.Rollback()
}

