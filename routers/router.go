package routers

import (
	"github.com/cst05001/fblog/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.BaseController{}, "get:Index")
    beego.Router("/rss", &controllers.BaseController{}, "get:IndexRss")
    beego.Router("/category/:id:int", &controllers.BaseController{}, "get:Category")

    beego.Router("/install", &controllers.BaseController{}, "get:InstallHtml")
    beego.Router("/install", &controllers.BaseController{}, "post:Install")
    beego.Router("/login", &controllers.BaseController{}, "get:LoginHtml")
    beego.Router("/login", &controllers.BaseController{}, "post:Login")
    beego.Router("/logout", &controllers.BaseController{}, "get:Logout")

    beego.Router("/user", &controllers.UserController{}, "get:Get")
    beego.Router("/post", &controllers.PostController{}, "get:AddHtml")
    beego.Router("/post", &controllers.PostController{}, "post:Add")
    beego.Router("/post/:id:int", &controllers.PostController{}, "get:Get")
    beego.Router("/post/:id:int/edit", &controllers.PostController{}, "get:EditHtml")
    beego.Router("/post/:id:int/delete", &controllers.PostController{}, "get:Delete")
    beego.Router("/post/:id:int", &controllers.PostController{}, "post:Edit")

    beego.Router("/tags", &controllers.TagController{}, "get:TagsManagement")
    beego.Router("/tag/:id:int", &controllers.TagController{}, "post:Update")

    beego.Router("/friendlink", &controllers.FriendLinkController{}, "post:Add")
    beego.Router("/friendlink/:id:int", &controllers.FriendLinkController{}, "post:Update")
    beego.Router("/friendlink", &controllers.FriendLinkController{}, "get:AddHtml")
    beego.Router("/friendlinks", &controllers.FriendLinkController{}, "get:Management")
}
