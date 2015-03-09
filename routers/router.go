package routers

import (
	"github.com/Kusold/test-config-issue/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
