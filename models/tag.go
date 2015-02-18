package models

import (
    "fmt"
    "github.com/astaxie/beego"
    "github.com/astaxie/beego/orm"
)

type Tag struct {
    Id          int64   `                   orm:"auto;pk"`
    Name        string  `form:"name"`
    Priority    int     `form:"priority"    orm:"default(99)"`
    IsCategory  bool    `form:"iscategory"  orm:"default(false)"`
    Posts       []*Post `                   orm:"reverse(many)"`
}

func NewTag() *Tag {
    t := &Tag{}
    t.IsCategory = false
    t.Priority = 99
    return t
}
func (this *Tag) Add() *Tag {
    t := this.Get()
    if t != nil {
        return t
    }
    o := orm.NewOrm()
    id, err := o.Insert(this)
    if err != nil {
        beego.Error(fmt.Sprintf("models> Tag> Add()> %v\n", err))
        return nil
    }

    this.Id = id
    this = this.Get()
    return this
}

func (this *Tag) Update() *Tag {
    o := orm.NewOrm()
    _, err := o.Update(this, "Name", "IsCategory", "Priority")
    if err != nil {
        beego.Error(fmt.Sprintf("models> Tag> Update()> %v\n", err))
        return nil
    }
    return this
}

// 这个Get是通过Name字段，不是通过Id
func (this *Tag) Get() *Tag {
    o := orm.NewOrm()
    var err error
    if len(this.Name) == 0 {
        err = o.Read(this)
    } else {
        err = o.QueryTable("tag").Filter("name", this.Name).One(this)
    }
    if err != nil {
        return nil
    }
    return this
}

func (this *Tag) AddIfNotExist() *Tag {
    t := this.Get()
    if t != nil {
        return t
    }
    return this.Add()
}
