package routers

import (
	"xinzhiyuntest/beegotest/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
