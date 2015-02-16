package controllers

import (
    //"fmt"
	"github.com/astaxie/beego"
)

type IndexController struct {
	beego.Controller
}

func (this *IndexController) Get() {
    PutBaseInfo(this.Controller)
    PutAllPostsInHtml(this.Controller)
    PutCategories(this.Controller)
    PutFriendLinks(this.Controller)
	this.TplNames = "index.html"
    this.Render()
}
