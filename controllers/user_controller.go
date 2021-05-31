package controllers

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/spf13/cast"
	"iris-demo/models"
	"iris-demo/service"
	"log"
)

type UserController struct {
	Ctx iris.Context
	Service service.UserService
}
func NewUserController() *UserController {
	return &UserController{Service:service.NewUserService()}
}

// GetList 查询所有
func (self *UserController) GetList()(result *models.Result) {
	return self.Service.GetUserList()
}

// PostSaveUser 保存 and 修改
func (self *UserController) PostSaveUser()(result models.Result)  {
	var user models.User
	if err := self.Ctx.ReadJSON(&user); err != nil {
		log.Println(err)
		result.Msg = "数据错误"
		return
	}

	return self.Service.PostSaveUser(user)
}

// GetUserById 根据id查询
func (self *UserController) GetUserById(ctx context.Context)(result models.Result)  {
	id := ctx.PostValue("id")
	if id == "" {
		result.Code = 400
		result.Msg = "缺少参数id"
		return
	}
	return self.Service.GetUserById(cast.ToUint(id))
}

// PostDelUser 根据id删除
func (self *UserController) PostDelUser(ctx context.Context)(result models.Result)  {
	id := ctx.PostValue("id")
	if id == "" {
		result.Code = 400
		result.Msg = "缺少参数id"
		return
	}
	return self.Service.DelUser(cast.ToUint(id))
}