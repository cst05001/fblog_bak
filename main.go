package main

import (
	_ "github.com/cst05001/fblog/routers"
	_ "github.com/cst05001/fblog/models"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

