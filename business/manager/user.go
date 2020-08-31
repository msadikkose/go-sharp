package manager

import (
	"go-sharp/dal/uow"
	"go-sharp/model"
)

type TrainData struct {
	sentence string
	sentiment string
	weight int
}

var TrainDataCity = []TrainData {
	{"I love the weather here.", "pos", 1700},
	{"This is an amazing place!", "pos", 2000},
	{"I feel very good about its food and atmosphere.", "pos", 2000},
	{"The location is very accessible.", "pos", 1500},
	{"One of the best cities I've ever been.", "pos", 2000},
	{"Definitely want to visit again.", "pos", 2000},
	{"I do not like this area.", "neg", 500},
	{"I am tired of this city.", "neg", 700},
	{"I can't deal with this town anymore.", "neg", 300},
	{"The weather is terrible.", "neg", 300},
	{"I hate this city.", "neg", 100},
	{"I won't come back!", "neg", 200},
}

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