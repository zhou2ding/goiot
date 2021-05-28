package testorm

import (
	"beego_project/models"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

//orm的CRUD基础用法
type TestOrm1Controller struct {
	beego.Controller
}

func (t *TestOrm1Controller) Get() {
	o := orm.NewOrm()
	o.Using("default") //不调用这个的话默认使用default
	//新增一条数据，循环插入时，如果有插入出错，要continue下，否则循环会被打破
	// usr := models.User{Name: "张三", Age: 18, Addr: "湖南"}
	// id, err := o.Insert(&usr)
	// fmt.Println(err, id)

	//新增多条，原子操作，中间有一条数据插入出错的话，整个都会执行失败
	// usrs := []models.User{{Name: "李四", Age: 18, Addr: "长沙"}, {Name: "王五", Age: 22, Addr: "常德"}}
	// cnt, err := o.InsertMulti(100, usrs) //一次最多插入100条(不要太大)，不够的话一次性插完，超的话每次100条。返回插入的条数
	// fmt.Println(err, cnt)

	//没有则插入，有则更新。逻辑：只根据有唯一约束的字段（如Id）查询，有则更新成其他的字段的值，没有则插入
	// usr := models.User{Id: 17, Name: "钱七", Age: 96, Addr: "江苏"} //Id这个主键字段必须带上，否则永远都是插入
	// o.InsertOrUpdate(&usr, "Id")                                //区别于查询，这里也要带上Id

	//查询一条数据，查询时只要带了主键字段Id，read中就不用传额外参数（如果要查的结果在数据库有多条，也只能查出一条来）
	find_usr := models.User{Id: 15, Name: "李四"}
	err := o.Read(&find_usr)
	fmt.Println(find_usr, err)
	// find_usr1 := models.User{Name: "张三", Addr: "湖南"}
	// err := o.Read(&find_usr1, "Name", "Addr") //read函数默认使用主键查询，要使用其他字段的话要加上结构体的字段名作为后续参数，主键字段无需指定
	// if err == orm.ErrNoRows {
	// 	fmt.Println("查询不到数据") //查询不到返回：<QuerySeter> no row found错误
	// } else {
	// 	fmt.Println(find_usr1)
	// }

	//若查询不到，则返回true和新创建的id，并创建；若查询到，返回false和查到的id
	// findcreate_usr := models.User{Name: "孙拔", Age: 8000, Addr: "傲来"}
	// is_new, id, err := o.ReadOrCreate(&findcreate_usr, "Name", "Addr") //返回是否新建的，id和错误；
	// fmt.Println(is_new, id, err)

	//更新数据，要先查询，才能更新。返回更新的数据的条数
	// usr := models.User{Name: "张三"}
	// err := o.Read(&usr, "Name")
	// if err == nil {
	// 	usr.Name = "张三丰"
	// 	usr.Addr = "武当山"
	// 	num, err := o.Update(&usr)
	// 	// num, err := o.Update(&usr, "Name") //只修改指定字段
	// 	fmt.Printf("更新了%v条数据，错误：%v\n", num, err)
	// } else {
	// 	fmt.Println("查询出错了")
	// }

	//删除数据，也是先查询，再删除，只不过删除的函数自带了查询功能，不用手动Read。返回删除的条数和err
	// usr := models.User{Name: "李四"}
	// num, err := o.Delete(&usr, "Name") //如果有多条数据满足条件的话，则都删掉
	// fmt.Printf("删除了%v条数据，错误：%v\n", num, err)

	t.TplName = "test_orm/orm1.html"
}
