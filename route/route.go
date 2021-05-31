package route

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"iris-demo/controllers"
)

func InitRouter(app *iris.Application) {
	bathUrl := "/api"
	mvc.New(app.Party(bathUrl +"/user")).Handle(controllers.NewUserController())
}