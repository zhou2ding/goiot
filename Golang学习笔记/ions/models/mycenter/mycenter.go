package mycenter

import "time"

type MyCenter struct {
	Id int `orm:"pk;auto"`
	CardId string `orm:"size(64)"`
	BaseSalary float64 `orm:"digits(12);decimal(2)"`
	WorkDays float64 `orm:"digits(3);decimal(1)"`
	DaysOff float64 `orm:"digits(3);decimal(1)"`
	DaysLeave float64 `orm:"digits(3);decimal(1)"`
	Bonus float64 `orm:"digits(8);decimal(2)"`
	RentSubsidy float64 `orm:"digits(6);decimal(2)"`
	TransSubsidy float64 `orm:"digits(6);decimal(2)"`
	SocialSec float64 `orm:"digits(6);decimal(2)"`
	HouseFund float64 `orm:"digits(6);decimal(2)"`
	Tax float64 `orm:"digits(6);decimal(2)"`
	Fine float64 `orm:"digits(6);decimal(2)"`
	NetSalary float64 `orm:"digits(6);decimal(2)"`
	PayDate time.Time `orm:"type(date);auto_now"`
	CreateTime time.Time `orm:"type(datetime);auto_now"`
}