package main

import (
	_ "github.com/sausheong/mosaic-b/routers"
	"github.com/sausheong/mosaic-b/mosaic"
	"github.com/astaxie/beego"
)

func main() {
	go mosaic.TilesDB()
	beego.Run()
}

