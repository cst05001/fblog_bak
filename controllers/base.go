package controllers

import (
    "fmt"
	"github.com/astaxie/beego"
    "github.com/astaxie/beego/orm"
    "github.com/astaxie/beego/validation"
    "github.com/cst05001/fblog/models"
)

type BaseController struct {
	beego.Controller
}

func (this *BaseController) InstallHtml() {
    /*
        安装博客的html
    */

    // User表存在，则放弃安装。
    o := orm.NewOrm()
    sum, err := o.QueryTable("user").Count()
    if err != nil {
        beego.Error(fmt.Sprintf("controller> Base> Post> err: %v\n", err))
        return
    }
    if sum > 0 {
        this.Ctx.WriteString("数据库已经初始化过，请不要重复操作。")
        return
    }

	this.TplNames = "install.html"
    this.Render()
}

func (this *BaseController) Install() {
    /*
        安装博客
    */

    // User表存在，则放弃安装。
    if IsInstalled(this.Controller) {
        this.Ctx.WriteString("数据库已经初始化过，请不要重复操作。")
        return
    }

    // 新建用户
    u := models.NewUser()
    if err := this.ParseForm(u); err != nil {
        beego.Error(fmt.Sprintf("controller> Base> Install()> err: %v\n", err))
        return
    }

    // 验证用户信息
    valid := validation.Validation{}
    valid.Required(u.Email, "需要提供电子邮箱")
    b, err := valid.Valid(u)
    if err != nil {
        beego.Error(fmt.Sprintf("controller> Base> Install()> err: %v\n", err))
        return
    }
    if !b {
        for _, err := range valid.Errors {
            beego.Error(fmt.Sprintf("controller> Base> Post> err :%v\n", err))
            this.Ctx.WriteString(fmt.Sprintf("%v<br/>", err))
        }
        this.Ctx.WriteString("表单验证失败")
        return
    }
    u.Role = models.ROLE_ADMIN
    u.Add() // 写入数据

    blogName := this.GetString("blogname")
    blog := &models.Blog{}
    blog.Name = blogName
    if ! blog.Add() {
        this.Ctx.WriteString("初始化博客信息失败")
        return
    }

    this.Ctx.WriteString("初始化成功")
}

func (this *BaseController) LoginHtml() {
    /*
        登陆界面的 html
    */
    this.TplNames = "login.html"
    this.Render()
}

func (this *BaseController) Login() {
    /*
        登陆验证
    */

    u := models.NewUser()
    if err := this.ParseForm(u); err != nil {
        beego.Error(fmt.Sprintf("controller> Base> Login()> err: %v\n", err))
        return
    }

    // 验证用户信息
    valid := validation.Validation{}
    b, err := valid.Valid(u)
    if err != nil {
        beego.Error(fmt.Sprintf("controller> Base> Install()> err: %v\n", err))
        return
    }
    if !b {
        for _, err := range valid.Errors {
            beego.Error(fmt.Sprintf("controller> Base> Post> err :%v\n", err))
            this.Ctx.WriteString(fmt.Sprintf("%v", err))
        }
        this.Ctx.WriteString("表单验证失败")
        return
    }

    u = u.Auth()
    if u == nil {
        beego.Debug(fmt.Sprintf("controller> Base> Login()> Auth Failed: %v\n", u))
        this.Ctx.WriteString("用户登陆信息验证失败")
        return
    }
    this.SetSession("user", u)
    this.Ctx.Output.Header("Content-Type", "text/html; charset=utf-8")
    this.Ctx.WriteString("登陆成功，去<a href=\"/user\">用户中心</a>。")
}

func (this *BaseController) Logout() {
    this.DelSession("user")
    this.Ctx.Output.Header("Content-Type", "text/html; charset=utf-8")
    this.Ctx.WriteString("已经注销,回到<a href=\"/\">首页</a>.")
}

