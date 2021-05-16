package utils

import (
	"beego_project/models"
	"strconv"

	"github.com/astaxie/beego/orm"
)

func Str2Int(s string) int {
	res, _ := strconv.ParseInt(s, 10, 0)
	return int(res)
}

func QueryRows(key, value string) []models.Article {
	o := orm.NewOrm()
	articles := []models.Article{}
	r := o.Raw("select * from article where "+key+" = ?", "")
	r.SetArgs(value).QueryRows(&articles)
	return articles
}
