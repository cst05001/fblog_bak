package controllers

import (
    "fmt"
    "strings"
    "strconv"
	"github.com/astaxie/beego"
	"github.com/cst05001/fblog/models"
)

type PostController struct {
	beego.Controller
}

func (this *PostController) Get() {
    i, err := strconv.Atoi(this.Ctx.Input.Param(":id"))
    if err != nil {
        beego.Error("controllers> Post> Get(): %v\n", err)
        this.Ctx.WriteString("无法获得文章ID")
        return
    }

    PutBaseInfo(this.Controller)
    PutCategories(this.Controller)

    post := GetPostInHtml(i)
    PutPerm(this.Controller, post)

    this.Data["post"] = post

    this.TplNames = "post.html"
    this.Render()
}


func (this *PostController) AddHtml() {
    PutBaseInfo(this.Controller)
    PutPerm(this.Controller, nil)
    if _, ok := this.Data["PERM_NEWPOST"]; ! ok {
        this.Ctx.WriteString("没有发表文章的权限")
        return
    }
    this.Data["PERM_NEWPOST"] = true

    this.TplNames = "newpost.html"
    this.Render()
}

func (this *PostController) Add() {
    if ! IsUser(this.Controller, true) {
        return
    }
    post := models.NewPost()
    if err := this.ParseForm(post); err != nil {
        beego.Error(fmt.Sprintf("controller> post> new()> err: %v\n", err))
        return
    }
    post.User = GetUserFromSession(this.Controller)

    PutPerm(this.Controller, post)

    if _, ok := this.Data["PERM_NEWPOST"]; ! ok {
        this.Ctx.WriteString("没有发表文章的权限")
        return
    }


    // 更新 Tags
    ts := strings.Split(this.GetString("tags"), ",")
    tags := make([]*models.Tag, 0)
    for k, _ := range ts {
        ts[k] = strings.TrimSpace(ts[k])
        beego.Debug(fmt.Sprintf("controller> Post> Edit> Trim: %s\n", ts[k]))
        tag := models.NewTag()
        tag.Name = ts[k]
        tags = append(tags, tag)
    }
    post.Tags = tags


    post = post.Add()

    if post != nil {
        this.Ctx.Output.Header("Content-Type", "text/html; charset=utf-8")
        this.Ctx.WriteString(fmt.Sprintf("发布成功。<a href=\"/post/%d\">链接</a>", post.Id))
    } else {
        this.Ctx.WriteString("发布失败")
    }
}

func (this *PostController) EditHtml() {
    i, err := strconv.Atoi(this.Ctx.Input.Param(":id"))
    if err != nil {
        beego.Error(fmt.Sprintf("controller> Post> EditHtml()> err: %v\n"), err)
        return
    }

    post := GetPost(i)
    PutBaseInfo(this.Controller)
    PutPerm(this.Controller, post)
    PutCategories(this.Controller)

    if ! IsUser(this.Controller, true) {
        return
    }

    if _, ok := this.Data["PERM_EDITPOST"]; ! ok {
        this.Ctx.WriteString("没有更新文章的权限")
        return
    }


    this.Data["post"] = post
    this.TplNames = "editpost.html"
    this.Render()

}

func (this *PostController) Edit() {
    user := GetUserFromSession(this.Controller)
    if user == nil {
        this.Ctx.Output.Header("Content-Type", "text/html; charset=utf-8")
        this.Ctx.WriteString("请先<a href=\"/login\">登录</a>")
        return
    }


    i, err := strconv.Atoi(this.Ctx.Input.Param(":id"))
    if err != nil {
        beego.Error(fmt.Sprintf("controller> Post> EditHtml()> err: %v\n"), err)
        return
    }

    // 获取 Post
    u := GetUserFromSession(this.Controller)
    p := GetPost(i)
    if err := this.ParseForm(p); err != nil {
        beego.Error(fmt.Sprintf("controller> post> new()> err: %v\n", err))
        return
    }
    p.User = u

    if ! u.HasPerm(models.PERM_EDITPOST, p) {
        this.Ctx.WriteString("没有更新文章的权限")
        return
    }

    // 更新 Tags
    ts := strings.Split(this.GetString("tags"), ",")
    tags := make([]*models.Tag, 0)
    for k, _ := range ts {
        ts[k] = strings.TrimSpace(ts[k])
        beego.Debug(fmt.Sprintf("controller> Post> Edit> Trim: %s\n", ts[k]))
        tag := models.NewTag()
        tag.Name = ts[k]
        tags = append(tags, tag)
    }
    p.Tags = tags

    // 更新 Post
    p = p.Update()
    if p == nil {
        this.Ctx.WriteString("更新失败")
        return
    }
    this.Ctx.Output.Header("Content-Type", "text/html; charset=utf-8")
    this.Ctx.WriteString(fmt.Sprintf("更新成功,<a href=\"/post/%d\">链接</a>", p.Id))

}

func (this *PostController) Delete() {
    if ! IsUser(this.Controller, true) {
        return
    }

    id, err := strconv.Atoi(this.Ctx.Input.Param(":id"))
    if err != nil {
        beego.Error(fmt.Sprintf("controller> Post> Delete()> err: %v\n"), err)
        return
    }

    p := GetPost(id)
    u := GetUserFromSession(this.Controller)
    p.User = u

    PutPerm(this.Controller, p)

    if _, ok := this.Data["PERM_DELPOST"]; ! ok {
        this.Ctx.WriteString("没有删除文章的权限")
        return
    }



    if p.Delete() {
        this.Ctx.Output.Header("Content-Type", "text/html; charset=utf-8")
        this.Ctx.WriteString("删除成功。<a href=\"/\">回首页</a>。")
    } else {
        this.Ctx.WriteString("删除失败")
    }
}
