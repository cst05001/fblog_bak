package controllers

import (
    "fmt"
    "strconv"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/cst05001/fblog/models"
)

type TagController struct {
	beego.Controller
}

func (this *TagController) Update() {
    PutCategories(this.Controller)
    if !IsUser(this.Controller, true) {
        return
    }
    u := GetUserFromSession(this.Controller)
    if ! u.HasPerm(models.PERM_EDITTAG) {
        this.Ctx.WriteString("没有编辑标签的权限")
        return
    }
    this.Data["PERM_EDITTAG"] = true

    i, err := strconv.Atoi(this.Ctx.Input.Param(":id"))
    if err != nil {
        beego.Error(fmt.Sprintf("controller> Post> Get()> %v\n"), err)
        return
    }
    tagId := int64(i)

    t := models.NewTag()
    t.Id = tagId
    if err := this.ParseForm(t); err != nil {
        beego.Error(fmt.Sprintf("controller> post> new()> err: %v\n", err))
        return
    }

    /*
    t = t.Get()
    if t == nil {
        this.Ctx.WriteString("没有找到标签")
        return
    }
    */

    t = t.Update()
    if t == nil {
        this.Ctx.WriteString("更新标签失败")
        return
    }

    this.Ctx.Output.Header("Content-Type", "text/html; charset=utf-8")
    this.Ctx.WriteString("更新标签成功。返回<a href=\"/tags\">标签管理</a>。")
}

func (this *TagController) TagsManagement() {
    PutBaseInfo(this.Controller)
    PutCategories(this.Controller)
    if !IsUser(this.Controller, true) {
        return
    }

    PutPerm(this.Controller, nil)
    if _, ok := this.Data["PERM_EDITTAG"]; ! ok {
        this.Ctx.WriteString("没有编辑标签的权限")
        return
    }

    tags := make([]models.Tag, 0)
    o := orm.NewOrm()
    //_, err := o.QueryTable("tag").RelatedSel().All(&tags)
    _, err := o.QueryTable("tag").All(&tags)
    if err != nil {
        beego.Error("controller> Tag> TagsManagement> %v\n", err)
        return
    }
    this.Data["tags"] = tags
    this.TplNames = "tagsmanagement.html"
    this.Render()
}
