package main

import (
	"fmt"

	"github.com/go-orm/gorm"
	_ "github.com/go-orm/gorm/dialects/mysql"
)

type User struct {
	Id   int64
	Name string
	Age  int64
}

func (u *User) TableName() string {
	return "user_info"
}

func main() {
	//1. 连接mysql数据库
	db, err := gorm.Open("mysql", "root:564710@(127.0.0.1:3306)/gin?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.LogMode(true)

	////2. 模型迁移
	// db.AutoMigrate(&User{})
	//
	////3. 插入数据
	//u := User{Name: sql.NullString{String: "张三", Valid: true},Age: 18}
	//db.NewRecord(&u)	//判断主键值是否为空
	//db.Create(&u)
	//db.Create(&User{Name: sql.NullString{String: "李四", Valid: true},Age: 23})

	//4. 查询
	//var usrs []User
	//db.Debug().Select("name,age").Find(&usrs)
	//var usr2 User
	//db.Table("user_info").Select("name").Where("name=?","李四").Scan(&usr2)
	// tb := db.Table("user_info")
	//5. 修改
	// usr := User{Id: 1}
	// tb.Model(&usr).Updates(User{Age: 123})
	sqlstr := `select * from user_info where id in(?) order by field(id,?);`
	ids := []int{10, 4, 11, 1, 5, 33}
	usr := make([]UserInfo, 0, 100)
	db.Raw(sqlstr, ids, ids).Find(&usr)
	fmt.Println(usr)
}

type UserInfo struct {
	Id   int
	Name string
	Age  int
}