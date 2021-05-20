package testorm

import (
	"beego_project/models"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

//QueryTable的其他接口函数
type TestOrmO2OController struct {
	beego.Controller
}

func (t *TestOrmO2OController) Prepare() {
	t.EnableXSRF = false
}

func (t *TestOrmO2OController) Get() {
	o := orm.NewOrm()
	// uid := 2 //一般由前端获取，下面的更新和删除要用到
	pid := 3 //正向查询和反向查询要用到

	//orm插入
	// usr := models.UserOne{UserName: "张三", Age: 18}
	// profile := models.ProfileOne{Phone: "13810293412", UserOne: &usr}
	// o.Insert(&usr)
	// o.Insert(&profile)

	//原生sql插入
	// r1 := o.Raw("insert into user_one(user_name,age) values(?,?)", "李四", "28")
	// ret, _ := r1.Exec()
	// user_one_id, _ := ret.LastInsertId()
	// r2 := o.Raw("insert into profile_one(phone,user_one_id) values(?,?)", "", user_one_id)
	// r2.Exec()

	//更新：orm
	// qs := o.QueryTable(new(models.ProfileOne))
	// qs.Filter("user_one_id__exact", 3).Update(orm.Params{"user_one_id": uid})

	//更新：原生sql
	// r3 := o.Raw("update profile_one set user_one_id = ? where user_one_id = ?", uid, 3)
	// r3.Exec()

	//删除：orm
	// qs := o.QueryTable(new(models.UserOne))
	// qs.Filter("id__exact", uid).Delete() //自带级联删除功能，即删了主表的数据会自动删子表的关联数据

	//删除：原生sql
	// r4 := o.Raw("delete from user_one where id = ?", uid)
	// r5 := o.Raw("delete from profile_one where user_one_id = ?", uid)
	// r4.Exec()
	// r5.Exec()

	//正向查询：RelatedSel()，根据子表的外键，把子表和主表都查到
	// po := models.ProfileOne{}
	// qs := o.QueryTable("profile_one")
	// qs.Filter("pid__exact", pid).RelatedSel().One(&po)
	// fmt.Println(po, po.UserOne)

	//反向查询：ProfileOne__Id，根据子表的外键，只能把主表查到
	usr := models.UserOne{}
	qs := o.QueryTable("user_one")
	qs.Filter("ProfileOne__Id", pid).One(&usr)
	fmt.Println(usr, usr.ProfileOne)

	t.TplName = "test_orm/o2o.html"
}
