package routers

import (
	"github.com/sausheong/mosaic-b/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
