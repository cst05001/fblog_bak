package models

import (
    "fmt"
    "time"
    "github.com/astaxie/beego"
    "github.com/astaxie/beego/orm"
)
type Post struct {
    Id          int64       `                                   orm:"auto;pk"`
    User        *User       `                                   orm:"rel(fk)"`
    Title       string      `form:"title"   valid:"Required"    orm:"index"`
    Content     string      `form:"content" valid:"Required"    orm:"type(text)"`
    Timestamp   time.Time   `                                   orm:"null;auto_now_add;type(datetime)"`
    Tags        []*Tag      `                                   orm:"rel(m2m)"`
}

func NewPost() *Post {
    p := &Post{}
    return p
}

func (this *Post) Add() *Post {
    o := orm.NewOrm()

    // 检查Tag，如果存在则填充Id，不存在则创建。
    for k, _ := range this.Tags {
        this.Tags[k] = this.Tags[k].AddIfNotExist()
    }
    id, err := o.Insert(this)
    if err != nil {
        beego.Error(fmt.Sprintf("models> Post> Update> err: %v\n", err))
        return nil
    }
    this.Id = id

    m2m := o.QueryM2M(this, "Tags")
    _, err = m2m.Clear()
    if err != nil {
        beego.Error(fmt.Sprintf("models> Post> Add()> err: %v\n", err))
        return nil
    }
    _, err = m2m.Add(this.Tags)
    if err != nil {
        beego.Error(fmt.Sprintf("models> Post> Add()> err: %v\n", err))
        return nil
    }


    this = this.Get()
    /*
    err = o.Read(this)
    if err != nil {
        beego.Error(fmt.Sprintf("models> Post> Update> err: %v\n", err))
        return nil
    }
    */
    return this
}

func (this *Post) Update() *Post {
    for k, _ := range this.Tags {
        this.Tags[k] = this.Tags[k].Add()
    }
    o := orm.NewOrm()
    m2m := o.QueryM2M(this, "Tags")
    _, err := m2m.Clear()
    if err != nil {
        beego.Error(fmt.Sprintf("models> Post> Update> err: %v\n", err))
        return nil
    }
    _, err = m2m.Add(this.Tags)
    if err != nil {
        beego.Error(fmt.Sprintf("models> Post> Update> err: %v\n", err))
        return nil
    }
    _, err = o.Update(this, "title", "content", "user")
    if err != nil {
        beego.Error(fmt.Sprintf("models> Post> Update> err: %v\n", err))
        return nil
    }

    /*
    err = o.Read(this)
    if err != nil {
        beego.Error(fmt.Sprintf("models> Post> Update> err: %v\n", err))
        return nil
    }
    */
    this = this.Get()
    return this
}

func (this *Post) Get() *Post {
    o := orm.NewOrm()
    err := o.QueryTable("post").RelatedSel().Filter("id", this.Id).One(this)
    if err != nil {
        beego.Error(fmt.Sprintf("models> Post> Get()> error: %v\n", err))
        return nil
    }

    _, err = o.LoadRelated(this, "Tags")
    if err != nil {
        beego.Error(fmt.Sprintf("models> Post> Get> LoadRelated> error: %v\n", err))
    }
    beego.Debug(fmt.Sprintf("models> Post> get> Tags: %v", this.Tags))
    return this
}

func (this *Post) Delete() bool {
    o := orm.NewOrm()


    err := o.Read(this)
    if err != nil {
        beego.Error(fmt.Sprintf("models> Post> Delete()> err: %v\n", err))
        return false
    }

    // 删除文章和Tags的关联
    m2m := o.QueryM2M(this, "Tags")
    _, err = m2m.Clear()
    if err != nil {
        beego.Error(fmt.Sprintf("models> Post> Delete()> err: %v\n", err))
        return false
    }
    /*
    m2m.Remove(this.Tags)
    if _, err = o.Delete(this); err == nil {
        return true
    }
    */
    // 删除文章
    _, err = o.Delete(this)
    if err == nil {
        return true
    }
    beego.Error(fmt.Sprintf("models> Post> Delete()> err: %v\n", err))
    return false
}
