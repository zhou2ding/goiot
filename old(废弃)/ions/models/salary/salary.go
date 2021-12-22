package salary

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Salary struct {
	Id int `orm:"pk;auto"`
	CardId string `orm:"size(64)"`
	BaseSalary float64 `orm:"digits(12);decimals(2)"`
	WorkDays float64 `orm:"digits(3);decimals(1)"`
	DaysOff float64 `orm:"digits(3);decimals(1)"`
	DaysLeave float64 `orm:"digits(3);decimals(1)"`
	Bonus float64 `orm:"digits(8);decimals(2)"`
	RentSubsidy float64 `orm:"digits(6);decimals(2)"`
	TransSubsidy float64 `orm:"digits(6);decimals(2)"`
	SocialSec float64 `orm:"digits(6);decimals(2)"`
	HouseFund float64 `orm:"digits(6);decimals(2)"`
	Tax float64 `orm:"digits(6);decimals(2)"`
	Fine float64 `orm:"digits(6);decimals(2)"`
	NetSalary float64 `orm:"digits(10);decimals(2)"`
	PayDate string `orm:"size(64)"`
	CreateTime time.Time `orm:"type(datetime);auto_now"`
}

func (s *Salary) TableName() string {
	return "sys_salary"
}

func init() {
	orm.RegisterModel(new(Salary))
}
