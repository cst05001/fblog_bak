package models

import (
    "fmt"
	"github.com/astaxie/beego"
    "github.com/astaxie/beego/orm"
)
type Blog struct {
    Id      int
    Name    string
}

func NewBlog() *Blog {
    blog := &Blog{}
    blog.Get()
    return blog
}

func (this *Blog) Add() bool {
    o := orm.NewOrm()
    sum, err := o.QueryTable("blog").Count()
    if sum > 1 {
        return false
    }
    _, err = o.Insert(this)
    if err != nil {
        beego.Error(fmt.Sprintf("models> Blog> Add()> %v\n", err))
        return false
    }
    return true
}

func (this *Blog) Get() *Blog {
    o := orm.NewOrm()
    err := o.QueryTable("blog").One(this)
    if err != nil {
        beego.Error(fmt.Sprintf("models> Blog> Get()> %v\n", err))
        return nil
    }
    return this
}

func (this *Blog) Update() *Blog {
    o := orm.NewOrm()
    _, err := o.Update(this)
    if err != nil {
        beego.Error(fmt.Sprintf("models> Blog> Update()> %v\n", err))
        return nil
    }
    return this
}
