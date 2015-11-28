package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

type Posts struct {
	Id      int64
	Content string
	Author  string
}

func init() {
	orm.RegisterDriver("postgres", orm.DR_Postgres)
	orm.RegisterDataBase("default", "postgres", "host=babar.elephantsql.com user=qjnqlnxs dbname=qjnqlnxs password=VSIjUweR1jWi3lf33R9EmlYVppgFJrwu sslmode=disable")
	orm.RegisterModel(new(Posts))
}

func AddPost(p *Posts) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(p)
	return
}

func GetPost(p *Posts) (err error) {
	o := orm.NewOrm()
	err = o.Read(p)
	return
}

func UpdatePost(p *Posts) (err error) {
	o := orm.NewOrm()
	_, err = o.Update(p)
	return
}

func DeletePost(p *Posts) (err error) {
	o := orm.NewOrm()
	_, err = o.Delete(p)
	return
}

