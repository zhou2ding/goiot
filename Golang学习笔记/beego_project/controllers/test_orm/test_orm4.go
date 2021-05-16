package testorm

import (
	"beego_project/models"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

//QueryTable的其他接口函数
type TestOrm4Controller struct {
	beego.Controller
}

func (t *TestOrm4Controller) Get() {
	o := orm.NewOrm()
	qs := o.QueryTable(new(models.Article))
	art := models.Article{}

	qs.Filter("title__contains", "西").One(&art, "Title", "Author")             //查出的结果只保留title和author字段，其他的不要。&art后面的参数最好用结构体字段名，也可以是数据库字段名，都能查出来
	qs.Filter("title__contains", "西").Filter("author__exact", "意大利").One(&art) //多个筛选条件，相当于mysql中的AND。所有的接口都能这样用

	//All，也可以指定需要显示的字段
	arts := []models.Article{}
	qs.Filter("title__contains", "西").All(&arts)                    //查询的满足条件的多个结果扔进arts切片中
	qs.Exclude("title__exact", "西游记").All(&arts)                    //把title等于西游记的数据排除掉，剩下的扔进arts切片
	qs.Limit(3, 1).All(&arts)                                       //相当于mysql中的Limit 1 3(startIndex size)，正好颠过来；因此，qs.Limit()第二个参数可省略
	qs.Limit(3).Offset(1).All(&arts)                                //效果同上
	qs.GroupBy("author").All(&arts)                                 //分组函数用法后续再补充
	qs.OrderBy("-id").All(&arts)                                    //根据id降序，默认升序；可以有多个排序条件，即可传多个参数。和mysql一样，groupby必须在orderby之前
	qs.Distinct().All(&arts, "author")                              //按指定字段去重
	cnt, err1 := qs.Filter("author__exact", "吴承恩").Count()          //作者等于吴承恩的查询结果的数量
	exist := qs.Filter("author__exact", "兰陵笑笑生").Exist()            //查询结果是否存在
	qs.Filter("id__exact", 12).Update(orm.Params{"author": "列奥尼达"}) //把id等于12的作者改为列奥尼达，{}中能传多个要修改的字段和值
	nums, err2 := qs.Filter("id__exact", 12).Delete()               //把id等于12的数据删除

	//先把需要插的数据扔进arts切片，然后遍历切片，把数据循环插入数据库，然后关闭
	qs.All(&arts, "Title", "Author")
	inster, err3 := qs.PrepareInsert()
	for _, v := range arts {
		id, err := inster.Insert(&v)
		fmt.Println(id, err)
	}
	inster.Close()

	var maps []orm.Params
	qs.Values(&maps) //把每行数据包装成map的形式扔进maps这个map切片中，返回maps的长度和err
	var lists []orm.ParamsList
	qs.ValuesList(&lists) //把每行数据包装成切片扔进lists这个切片中，同上
	var flats orm.ParamsList
	qs.ValuesFlat(&flats, "title") //只把指定字段的整个查询结果扔进flats这个切片中，同上

	fmt.Println(exist, cnt, nums, err1, err2, err3)
	t.TplName = "test_orm/orm4.html"
}
