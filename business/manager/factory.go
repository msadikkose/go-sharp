package manager

import "go-sharp/dal/uow"

type Factory struct {
    Uow uow.IUnitOfWork
    UserManager IUser
}
