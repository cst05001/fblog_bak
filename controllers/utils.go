package controllers

import (
    "fmt"
	"github.com/astaxie/beego"
    "github.com/astaxie/beego/orm"
    "github.com/cst05001/fblog/models"
    "github.com/russross/blackfriday"
)

func GetPost(id int) *models.Post {
    postId := int64(id)
    p := models.NewPost()
    p.Id = postId
    p = p.Get()
    if p == nil {
        return nil
    }
    return p
}

func GetPostInHtml(id int) *models.Post {
    p := GetPost(id)
    p.Content = string(blackfriday.MarkdownBasic([]byte(p.Content)))
    return p
}

func PutBaseInfo(this beego.Controller) {
    blog := models.NewBlog()
    this.Data["blog"] = blog
}

func PutCategories(this beego.Controller) error {
    o := orm.NewOrm()
    categories := make([]models.Tag, 0)
    _, err := o.QueryTable("tag").Filter("iscategory", true).OrderBy("priority").RelatedSel().All(&categories)
    if err != nil {
        beego.Error("controller> Utils> GetCategories()> %v\n", err)
        return err
    }
    this.Data["categories"] = categories
    return nil
}

func PutFriendLinks(this beego.Controller) error {
    o := orm.NewOrm()
    friendLinks := make([]models.FriendLink, 0)
    _, err := o.QueryTable("friend_link").OrderBy("priority").All(&friendLinks)
    if err != nil {
        beego.Error("controller> Utils> GetFriendLink()> %v\n", err)
        return err
    }
    this.Data["friendlinks"] = friendLinks
    return nil
}


func PutAllPostsInHtml(this beego.Controller) error {
    posts := make([]models.Post, 0)
    o := orm.NewOrm()
    _, err := o.QueryTable("post").OrderBy("-timestamp").RelatedSel().All(&posts)
    if err != nil {
        beego.Error("controller> Utils> GetCategories()> %v\n", err)
        return err
    }
    for i, _ := range posts {
        posts[i].Content = string(blackfriday.MarkdownBasic([]byte(posts[i].Content)))
    }

    this.Data["posts"] = posts
    return nil
}

func GetUserFromSession(this beego.Controller) *models.User {
    v := this.GetSession("user")
    if v != nil {
        u := v.(*models.User)
        return u
    }
    return nil
}

func PutPerm(this beego.Controller, post *models.Post) {
    user := GetUserFromSession(this)
    if user == nil {
        return
    }

    if user.HasPerm(models.PERM_NEWPOST) {
        this.Data["PERM_NEWPOST"] = true
    }
    if user.HasPerm(models.PERM_EDITPOST, post) {
        this.Data["PERM_EDITPOST"] = true
    }
    if user.HasPerm(models.PERM_DELPOST, post) {
        this.Data["PERM_DELPOST"] = true
    }
    if user.HasPerm(models.PERM_EDITPOST, post) {
        this.Data["PERM_EDITPOST"] = true
    }
    if user.HasPerm(models.PERM_EDITTAG) {
        this.Data["PERM_EDITTAG"] = true
    }
    if user.HasPerm(models.PERM_EDITFRIENDLINK) {
        this.Data["PERM_EDITFRIENDLINK"] = true
    }
}


func IsUser(this beego.Controller, render bool) bool {
    if GetUserFromSession(this) != nil {
        return true
    }
    if render {
        this.Ctx.Output.Header("Content-Type", "text/html; charset=utf-8")
        this.Ctx.WriteString("请先<a href=\"/login\">登录</a>")
    }
    return false
}

func IsInstalled(this beego.Controller) bool {
    // User表存在，则放弃安装。
    o := orm.NewOrm()
    sum, err := o.QueryTable("blog").Count()
    if err != nil {
        beego.Error(fmt.Sprintf("controller> Base> Install> %v\n", err))
        return true
    }
    if sum > 0 {
        return true
    }
    return false
}