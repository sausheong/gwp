package main

import (
	_ "github.com/sausheong/gwp/Chapter_11_Frameworks_Frameworks_Everywhere/ws-b/docs"
	_ "github.com/sausheong/gwp/Chapter_11_Frameworks_Frameworks_Everywhere/ws-b/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.DirectoryIndex = true
	beego.StaticDir["/swagger"] = "swagger"
	beego.Run()
}
