package manager

import (
	"go-sharp/dal/uow"
	"go-sharp/model"
)


type IUser interface {
	GetById(idUser int) (*model.User, error)
	GetAll() (*[]model.User,error)
	Create(*model.User) error
	Update(*model.User) error
	Delete(idUser int) error
}

type User struct {
	Uow uow.IUnitOfWork
}

func (u *User) GetById(idUser int) (*model.User, error) {

	userEntity, err := u.Uow.GetUserRepository().Get(idUser)
	if err != nil {
		return nil, err
	}

	usr := &model.User{
		Age:  userEntity.Age,
		Name: userEntity.Name,
		Id:userEntity.ID,
	}
	return usr, nil
}

func (u *User) GetAll() (*[]model.User,error) {
	userList,err:=u.Uow.GetUserRepository().GetAll()
	if err!=nil{
		return nil, err
	}

	userModelList :=make([]model.User,len(userList))

	for i,usr :=range userList{
		userModelList[i].Name = usr.Name
		userModelList[i].Age = usr.Age
		userModelList[i].Id =usr.ID
	}
	return &userModelList,nil


}

func (u *User) Create(usr *model.User ) error{

	err:=u.Uow.GetUserRepository().Create(usr)
	if err!=nil{
		return err
	}
	return nil
}

func (u *User) Update(usr *model.User ) error{

	err:=u.Uow.GetUserRepository().Update(usr)
	if err!=nil{
		return err
	}
	return nil
}

func (u *User) Delete(idUser int)  error {

	err := u.Uow.GetUserRepository().Delete(idUser)
	if err != nil {
		return nil
	}
	return  nil
}