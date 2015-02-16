package controllers

import (
    "fmt"
    "strconv"
	"github.com/astaxie/beego"
	"github.com/cst05001/fblog/models"
)

type FriendLinkController struct {
	beego.Controller
}

func (this *FriendLinkController) Management() {
    if !IsUser(this.Controller, true) {
        return
    }

    u := GetUserFromSession(this.Controller)
    if ! u.HasPerm(models.PERM_EDITFRIENDLINK) {
        this.Ctx.WriteString("没有编辑友情链接的权限")
        return
    }

    PutBaseInfo(this.Controller)
    PutCategories(this.Controller)
    PutFriendLinks(this.Controller)
	this.TplNames = "friendlinkmanagement.html"
    this.Render()
}

func (this *FriendLinkController) Update() {
    if !IsUser(this.Controller, true) {
        return
    }
    u := GetUserFromSession(this.Controller)
    if ! u.HasPerm(models.PERM_EDITFRIENDLINK) {
        this.Ctx.WriteString("没有编辑友情链接的权限")
        return
    }

    id, err := strconv.Atoi(this.Ctx.Input.Param(":id"))
    if err != nil {
        beego.Error(fmt.Sprintf("controller> FriendLink> Update()> %v\n"), err)
        return
    }

    friendLink := models.NewFriendLink()
    friendLink.Id = int64(id)
    if err := this.ParseForm(friendLink); err != nil {
        beego.Error(fmt.Sprintf("controller> FriendLink> Update()> %v\n", err))
        return
    }

    friendLink = friendLink.Update()

    if friendLink == nil {
        this.Ctx.WriteString("更新友情链接失败")
        return
    }

    this.Ctx.Output.Header("Content-Type", "text/html; charset=utf-8")
    this.Ctx.WriteString("更新友情链接成功。返回<a href=\"/friendlinks\">友情链接管理</a>。")
}

func (this *FriendLinkController) AddHtml() {
    if ! IsUser(this.Controller, true) {
        return
    }

    user := GetUserFromSession(this.Controller)
    if ! user.HasPerm(models.PERM_NEWFRIENDLINK) {
        this.Ctx.WriteString("没有权限新建友情链接")
        return
    }

	this.TplNames = "addfriendtag.html"
    this.Render()
}

func (this *FriendLinkController) Add() {
    if ! IsUser(this.Controller, true) {
        return
    }

    user := GetUserFromSession(this.Controller)
    if ! user.HasPerm(models.PERM_NEWFRIENDLINK) {
        this.Ctx.WriteString("没有权限新建友情链接")
        return
    }

    friendLink := models.NewFriendLink()
    if err := this.ParseForm(friendLink); err != nil {
        beego.Error(fmt.Sprintf("controller> FriendLink> Add()> err: %v\n", err))
        return
    }
    friendLink = friendLink.Add()
    if friendLink != nil {
        this.Ctx.Output.Header("Content-Type", "text/html; charset=utf-8")
        this.Ctx.WriteString("添加友情链接成功。\n<a href=\"/user\">用户中心</a>\n<a href=\"/\">首页</a>")
    } else {
        this.Ctx.WriteString("添加友情链接失败")
    }
}
