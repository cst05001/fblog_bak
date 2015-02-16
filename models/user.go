package models

import (
    "fmt"
    "encoding/hex"
    "crypto/md5"
    "github.com/astaxie/beego"
    "github.com/astaxie/beego/orm"
)

const (
    ROLE_VISITOR = iota
    ROLE_EDITOR
    ROLE_ADMIN
)

const (
    /*
        权限常量定义
    */
    PERM_NEWPOST = iota
    PERM_EDITPOST
    PERM_DELPOST
    PERM_EDITTAG
    PERM_NEWFRIENDLINK
    PERM_DELFRIENDLINK
    PERM_EDITFRIENDLINK
)

type User struct {
    Id      int64   `form:"-"                                           orm:"auto;pk"`
    Name        string  `form:"name"        valid:"Required;AlphaNumeric"   orm:"index"`
    Password    string  `form:"password"    valid:"Required"`
    Email       string  `form:"email"                                       orm:"null"`
    Timestamp   string  `form:"-"                                           orm:"auto_now_add;type(timestamp)"`
    Role        int     `form:"-"                                           orm:"default(0)"`
    Posts       []*Post `form:"-"                                           orm:"reverse(many);on_delete(cascade)"`
}

func NewUser() *User {
    u := &User{Role: ROLE_VISITOR}
    return u
}

func (this *User) Add() *User {
    /*
        创建用户
    */

    beego.Debug("models> User> Add()")

    // 用 MD5 加密密码
    passwordHasher := md5.New()
    passwordHasher.Write([]byte(this.Password))
    this.Password = hex.EncodeToString(passwordHasher.Sum(nil))

    // 插入新用户
    o := orm.NewOrm()
    id, err := o.Insert(this)
    if err != nil {
        beego.Error(fmt.Sprintf("models> User> Add> insert: %v\n", err))
        return nil
    }
    beego.Debug(fmt.Sprintf("models> User> Add> insert: %v\n", this))

    // 获取用户信息
    this.Id = id
    err = o.Read(this)
    if err != nil {
        beego.Error(fmt.Sprintf("models> User> Add> read: %v\n", err))
        return nil
    }

    return this
}

func (this *User) Auth() *User {
    /*
        获取用户信息，一般用来做登陆验证
    */

    // 用 MD5 加密密码
    passwordHasher := md5.New()
    passwordHasher.Write([]byte(this.Password))
    this.Password = hex.EncodeToString(passwordHasher.Sum(nil))

    // 通过查询用户名和密码都匹配的方式，验证用户
    o := orm.NewOrm()
    err := o.QueryTable("user").Filter("name", this.Name).Filter("password", this.Password).One(this)
    if err != nil {
        beego.Error(fmt.Sprintf("models> User> Get> one: %v\n", err))
        return nil
    }
    return this
}

func (this *User) HasPerm(perm int, vs ...interface{}) bool {
    switch perm {
    case PERM_NEWPOST:
        if this.Role == ROLE_EDITOR || this.Role == ROLE_ADMIN {
            return true
        }
        return false
    case PERM_EDITPOST:
        v := vs[0]
        post := v.(*Post)
        if post != nil && post.User.Id == this.Id {
            return true
        }
        return false
    case PERM_DELPOST:
        if len(vs) == 0 {
            return false
        }
        v := vs[0]
        if v == nil {
            return false
        }
        post := v.(*Post)
        if post != nil && post.User.Id == this.Id {
            return true
        }
        return false
    case PERM_EDITTAG:
        if this.Role == ROLE_ADMIN {
            return true
        }
        return false
    case PERM_NEWFRIENDLINK:
        if this.Role == ROLE_ADMIN {
            return true
        }
        return false
    case PERM_EDITFRIENDLINK:
        if this.Role == ROLE_ADMIN {
            return true
        }
        return false
    case PERM_DELFRIENDLINK:
        if this.Role == ROLE_ADMIN {
            return true
        }
        return false
    default:
        beego.Error(fmt.Sprintf("No such perm type: %s\n", perm))
        return false
    }
}
