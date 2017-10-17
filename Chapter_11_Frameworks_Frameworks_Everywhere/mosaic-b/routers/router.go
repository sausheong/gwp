package routers

import (
	"github.com/sausheong/gwp/Chapter_11_Frameworks_Frameworks_Everywhere/mosaic-b/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
