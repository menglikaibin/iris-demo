package service

import (
	"iris-demo/models"
	"iris-demo/repo"
)
type UserService interface {
	GetUserList () *models.Result
	PostSaveUser(user models.User) (result models.Result)
	GetUserById(id uint) (result models.Result)
	DelUser(id uint) (result models.Result)
}

type userService struct {}

func NewUserService() UserService{
	return &userService{}
}

var userRepo = repo.NewUserRepository()

func (self userService)GetUserList()  *models.Result{
	books := userRepo.GetUserList()
	result := new (models.Result)
	result.Data = books
	result.Code = 200
	result.Msg ="SUCCESS"
	return result
}
func (self userService) PostSaveUser(user models.User)(result models.Result){
	err := userRepo.SaveUser(user)
	if err != nil{
		result.Code = 400
		result.Msg = err.Error()
	}else{
		result.Code = 200
		result.Msg ="SUCCESS"
		user := userRepo.GetUserByName(user.Name)
		result.Data = user
	}
	return
}
func (self userService) GetUserById(id uint)(result models.Result){
	user,err := userRepo.GetUserById(id)
	if err!= nil{
		result.Code = 400
		result.Msg = err.Error()
	}else{
		result.Data = user
		result.Code = 200
		result.Msg ="SUCCESS"
	}
	return result
}
func (self userService) DelUser(id uint)(result models.Result){
	err := userRepo.DeleteUser(id)
	if err!= nil{
		result.Code = 400
		result.Msg = err.Error()
	}else{
		result.Code = 200
		result.Msg ="SUCCESS"
		list := userRepo.GetUserList()
		result.Data = list

	}
	return
}
 