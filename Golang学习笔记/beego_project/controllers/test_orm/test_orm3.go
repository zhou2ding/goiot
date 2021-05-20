package testorm

import (
	"beego_project/models"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

//expr表达式用法
type TestOrm3Controller struct {
	beego.Controller
}

func (t *TestOrm3Controller) Get() {
	o := orm.NewOrm()
	// qs := o.QueryTable("article") //两种初始化queryseter（qs）的方法
	qs := o.QueryTable(new(models.Article))

	// art := models.Article{}
	//大小写敏感的where子句中是：= binary，like binary；不敏感：=，like
	// qs.Filter("title__exact", "红楼梦").One(&art)    //把title等于红楼梦的单一查询结果，扔进art这个结构体中。有多个结果时用最先查到的
	// qs.Filter("title__iexact", "红楼梦").One(&art)   //expr加了i，就是大小写不敏感。其他的expr除了gt这些，都有加i的写法
	// qs.Filter("title__contains", "红").One((&art)) //把title包含“红”的单一查询结果，扔进art。
	// qs.Filter("id__gt", "3").One((&art))          //把id大于4的单一查询结果，扔进art。 一共有gt、gte，lt、lte四个
	// qs.Filter("title__startswith", "红").One(&art) //把title已“红”开始的单一查询结果，扔进art。
	// qs.Filter("title__endswith", "梦").One(&art)   //把title已“梦”结尾的单一查询结果，扔进art。
	// qs.Filter("id__in", 3,4,5,6).One(&art) //id在3、4、5、6这个四个数字内的查询结果，给的数字范围也可以是切片、结构体切片
	// qs.Filter("title__isnull", true).One(&art)    //title的值is null的查询结果
	// qs.Filter("title__isnull", false).One(&art)   //title的值is not null的查询结果

	arts := []models.Article{}
	inarr := []int{3, 4, 5, 6}
	qs.Filter("id__in", inarr).All(&arts)
	fmt.Println(arts)

	t.TplName = "test_orm/orm3.html"
}
