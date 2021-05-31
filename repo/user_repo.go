package repo

import (
	"iris-demo/datasource"
	"iris-demo/models"
)

type UserRepository interface {
	GetUserList() *[]models.User
	SaveUser(book models.User)(err error)
	GetUserById(id uint)(book models.User,err error)
	DeleteUser(id uint)(err error)
	GetUserByName(name string )*[]models.User
}

func NewUserRepository() UserRepository{
	return &userRepository{}
}
var db = datasource.GetDB()

type userRepository struct {}

func (self userRepository) GetUserList()*[]models.User{
	user:= new([]models.User)
	err:=db.Raw(`select * FROM user`).Scan(user).RowsAffected
	if err > 0 {
		return user
	}else{
		return nil
	}
}
func (self userRepository)GetUserByName(name string )*[]models.User{
	user := new([]models.User)
	err := db.Raw(`select * FROM user where user.name = ?`,name).Scan(user).RowsAffected
	if err > 0 {
		return user
	}else{
		return nil
	}
}


func (self userRepository) SaveUser(user models.User)(err error){
	if user.ID != 0{
		err := db.Save(&user).Error
		return err
	}else {
		err := db.Create(&user).Error
		return err
	}
}
func (self userRepository) GetUserById(id uint)(user models.User,err error){
	err = db.First(&user,id).Error
	return user,err
}
func (self userRepository) DeleteUser(id uint)(err error){
	user:= new(models.User)
	user.ID = id
	err = db.Unscoped().Delete(&user).Error
	return err
}
