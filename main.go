package main

import (
	"buyit/dao"
	_ "buyit/routers"

	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	dao.Init()
	beego.Run()
}

