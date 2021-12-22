package profit

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Profit struct {
	Id int `orm:"pk;auto"`
	Mon string `orm:"size(16)"`
	Sales float64 `orm:"digits(10);decimals(2)"`
	StuIncr int
	Django int
	VueDjango int
	Celery int
	CreateDate time.Time `orm:"type(datetime);auto_now"`
}

func (p *Profit) TableName() string {
	return "sys_caiwu_data"
}

func init() {
	orm.RegisterModel(new(Profit))
}
