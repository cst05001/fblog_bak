package controllers

import (
	"github.com/astaxie/beego"
	"github.com/cst05001/fblog/models"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) Get() {
    PutBaseInfo(this.Controller)
    PutFriendLinks(this.Controller)
    PutCategories(this.Controller)
    v := this.GetSession("user")
    if v == nil {
        this.Ctx.Output.Header("Content-Type", "text/html; charset=utf-8")
        this.Ctx.WriteString("请先<a href=\"/login\">登录</a>")
        return
    }
    u := v.(*models.User)
    this.Data["USER"] = u
    PutPerm(this.Controller, nil)
	this.TplNames = "user.html"
    this.Render()
}
