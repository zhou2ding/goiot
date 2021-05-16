package testorm

import (
	"beego_project/models"
	"beego_project/utils"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type TestOrm5Controller struct {
	beego.Controller
}

func (t *TestOrm5Controller) Get() {
	o := orm.NewOrm()

	//增删改
	r1 := o.Raw("insert into article(title,author) values(?,?)", "金瓶梅", "兰陵笑笑生")
	r2 := o.Raw("update article set title = ? where id = ?", "乱世佳人", 12)
	r3 := o.Raw("delete from article where id = ?", 12)
	ret1, err1 := r1.Exec() //返回sql执行的结果和err
	ret2, err2 := r2.Exec()
	ret3, err3 := r3.Exec()
	fmt.Println(ret1, ret2, ret3, err1, err2, err3)
	fmt.Println(ret1.LastInsertId())
	fmt.Println(ret1.RowsAffected())

	//查询单条数据
	r4 := o.Raw("SELECT * FROM article where id = ?", 12)
	//方式1，基本用不到
	var (
		id            int
		title, author string
	)
	r4.QueryRow(&id, &title, &author)
	//方式2，明确知道是一条时使用
	art := models.Article{}
	r4.QueryRow(&art)

	//查询多条数据，最常用
	arts := []models.Article{}
	r5 := o.Raw("select * from article where title like ?", "%西%")
	r5.QueryRows(&arts)
	//避免重复写sql语句，作用是替换上面SQL语句中?的值
	arts1 := []models.Article{}
	r5.SetArgs("%游%").QueryRows(&arts1)
	//封装成自用函数
	arts2 := utils.QueryRows("id", "4")
	fmt.Println("工具查询结果：", arts2)
	//values、valueList、valueFlat，作用和queryset的同名方法一样

	//RowsToStruct：第二三个参数是把字段的值分别作为key、value；用于统计数据的表转换到结构体
	//RowsToMap：传参同上；用于统计数据的表转换到map
	to_struct := ToStruct{}
	to_map := make(orm.Params)
	r6 := o.Raw("SELECT name,age FROM to_struct")
	nums1, _ := r6.RowsToStruct(&to_struct, "name", "age")
	nums2, _ := r6.RowsToMap(&to_map, "namge", "age")

	//一次准备多次执行，类似queryseter的PrepareInsert，这个还能update和delete
	r7 := o.Raw("insert into article(title,author) values(?,?)")
	rp, _ := r7.Prepare()
	rp.Exec("小李飞刀", "古龙")
	rp.Exec("温瑞安", "四大名捕")
	rp.Close()

	fmt.Printf("%v\n,%v\n,%v\n,%v\n,%v\n,to_struct:%v,to_map:%v,nums1:%v,nums2:%v\n", id, title, author, art, arts, to_struct, to_map, nums1, nums2)
	t.TplName = "test_orm/orm5.html"
}

type ToStruct struct {
	//结构体的两个字段对应表中作为key的列的值，表结构如下
	//id   name    age
	//1    total    100
	//2    page     5
	Total string
	Page  string
}
