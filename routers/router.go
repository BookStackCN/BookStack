package routers

import (
	"github.com/BookStackCN/BookStack/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/read/*", &controllers.MainController{}, "get:MarkdownRender")
}
