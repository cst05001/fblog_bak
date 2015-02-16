package models

import (
    "fmt"
     _ "github.com/mattn/go-sqlite3"
    "github.com/astaxie/beego" 
    "github.com/astaxie/beego/orm" 
)

func init() {
    orm.RunCommand() // ORM CLI支持

    // 注册表模型
    orm.RegisterModel(new(User))
    orm.RegisterModel(new(Post))
    orm.RegisterModel(new(Tag))
    orm.RegisterModel(new(Blog))
    orm.RegisterModel(new(FriendLink))

    //数据库连接参数
    dbname := "default"
    driver := "sqlite3"
    maxIdle := 10
    maxConn := 10
    force := false
    verbose := true
    orm.RegisterDataBase(dbname, driver, "fblog.sqlite", maxIdle, maxConn) // 创建连接

    err := orm.RunSyncdb(dbname, force ,verbose) //自动建表
    if err != nil {
        beego.Error(fmt.Sprintf("models> init()> orm.RunSyncdb> err: %v\n", err))
    }
}
