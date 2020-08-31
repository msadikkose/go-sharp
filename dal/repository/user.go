package repository

import (
	"go-sharp/ent"
	"go-sharp/ent/user"
	"go-sharp/model"
	"sync"
)

type IUser interface {
	Get(idUser int) (*ent.User, error)
	GetAll() ([]*ent.User, error)
	Create(userEnt *model.User) error
	Delete(idUser int)  error
	Update(usr *model.User) error
}

type User struct {
	Base
}

func (u *User) Get(idUser int) (*ent.User, error) {
	var usr *ent.User
	var err error
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		usr, err = u.Transaction.User.Query().Where(user.IDEQ(idUser)).Where(user.IsDeleted(false)).Only(u.Ctx)
	}()
	wg.Wait()
	if err != nil {
		return nil, err
	}
	return usr, nil
}

func (u *User) GetAll() ([]*ent.User, error) {

	var usr []*ent.User
	var err error
	wg := make(chan int)
	go func() {
		usr, err = u.Transaction.User.Query().Where(user.IsDeleted(false)).Order(ent.Asc(user.FieldID)).All(u.Ctx)
		wg <- 1
	}()
	<-wg
	if err != nil {
		return nil, err
	}
	return usr, nil
}
func (u *User) Create(userModel *model.User) error {
	var err error
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		_, err = u.Transaction.User.
			Create().
			SetAge(userModel.Age).
			SetName(userModel.Name).
			SetIsDeleted(false).
			Save(u.Ctx)
		defer wg.Done()
	}()
	wg.Wait()
	if err !=nil{
		return err
	}
	return nil
}

func (u *User) Update(userModel *model.User) error {
	var err error
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		if(userModel.Id>0) {
			_, err = u.Transaction.User.
				UpdateOneID(userModel.Id).
				SetAge(userModel.Age).
				SetName(userModel.Name).
				SetIsDeleted(false).
				Save(u.Ctx)
		}
		defer wg.Done()
	}()
	wg.Wait()
	if err != nil{
		return err
	}
	return nil
}

func (u *User) Delete(idUser int)  error {

	var err error
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		_, err = u.Transaction.User.
			UpdateOneID(idUser).
			SetIsDeleted(true).
			Save(u.Ctx)

	}()
	wg.Wait()
	if err != nil {
		return nil
	}

	return  nil
}
