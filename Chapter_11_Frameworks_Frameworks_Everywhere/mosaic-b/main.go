package main

import (
	_ "github.com/sausheong/gwp/Chapter_11_Frameworks_Frameworks_Everywhere/mosaic-b/routers"
	"github.com/sausheong/gwp/Chapter_11_Frameworks_Frameworks_Everywhere/mosaic-b/mosaic"
	"github.com/astaxie/beego"
)

func main() {
	go mosaic.TilesDB()
	beego.Run()
}

