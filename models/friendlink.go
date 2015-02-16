package models

import (
    "fmt"
    "github.com/astaxie/beego"
    "github.com/astaxie/beego/orm"
)

type FriendLink struct {
    Id          int64   `form:"-"           orm:"auto;pk"`
    Name        string  `form:"name"`
    Url         string  `form:"url"`
    Priority    int     `form:"priority"    orm:"default(99)"`
}

func NewFriendLink() *FriendLink {
    fl := &FriendLink{}
    return fl
}

func (this *FriendLink) Add() *FriendLink {
    o := orm.NewOrm()
    id, err := o.Insert(this)
    if err != nil {
        beego.Error(fmt.Sprintf("models> FriendLink> Add()> %v\n", err))
        return nil
    }

    this.Id = id
    this = this.Get()
    return this
}

func (this *FriendLink) Get() *FriendLink {
    o := orm.NewOrm()
    err := o.Read(this)
    if err != nil {
        beego.Error(fmt.Sprintf("models> FriendLink> Get()> %v\n", err))
        return nil
    }
    return this
}

func (this *FriendLink) Update() *FriendLink {
    o := orm.NewOrm()
    _, err := o.Update(this, "name", "url", "priority")

    if err != nil {
        beego.Error(fmt.Sprintf("models> Post> Update> err: %v\n", err))
        return nil
    }
    return this
}
