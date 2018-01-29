package routers

import (
	"github.com/BookStackCN/BookStack/BookStack/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
