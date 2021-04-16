package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type user struct {
	id   int
	name string
	age  int
}

func initDB() (err error) {
	dsn := "root:564710@tcp(localhost)/go_sql"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return
	}
	err = db.Ping()
	if err != nil {
		return
	}
	return
}

func queryOne(id int) {
	var u1 user
	sqlStr := `select id, name, age from user where id=?;`
	db.QueryRow(sqlStr, id).Scan(&u1.id, &u1.name, &u1.age)
	fmt.Printf("%#v\n", u1)
}

func queryMore(n int) {
	sqlStr := `select id, name, age from user where id>?;`
	rows, err := db.Query(sqlStr, n)
	if err != nil {
		fmt.Printf("query more failed, error:%v\n", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var u1 user
		rows.Scan(&u1.id, &u1.name, &u1.age)
		fmt.Printf("%#v\n", u1)
	}
}

func insert(name string, age int) {
	sqlStr := `insert into user(name,age) values(?,?)`
	ret, err := db.Exec(sqlStr, name, age)
	if err != nil {
		fmt.Printf("insert failed, error:%v\n", err)
		return
	}
	id, _ := ret.LastInsertId()
	fmt.Println("插入数据的id:", id)
}

func update(id, age int) {
	sqlStr := `update user set age = ? where id > ?`
	ret, err := db.Exec(sqlStr, age, id)
	if err != nil {
		fmt.Printf("update failed, error:%v\n", err)
		return
	}
	n, _ := ret.RowsAffected()
	fmt.Printf("修改了%d行", n)
}

func delete(id int) {
	sqlStr := `delete from user where id = ?`
	ret, err := db.Exec(sqlStr, id)
	if err != nil {
		fmt.Printf("delete failed, error:%v\n", err)
		return
	}
	n, _ := ret.RowsAffected()
	fmt.Printf("修改了%d行\n", n)
}

func prepareInsert(data map[string]int) {
	sqlStr := `insert into user(name,age) values(?,?)`
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare failed, error:%v\n", err)
		return
	}
	defer stmt.Close()
	for k, v := range data {
		stmt.Exec(k, v)
	}
}

func transaction() {
	tx, err := db.Begin()
	if err != nil {
		fmt.Printf("start transaction failed, error:%v\v", err)
		return
	}
	sqlStr1 := `insert into user(name,age) values('孔子',1234)`
	sqlStr2 := `insert into user(name,age) values('孟子',4321)`
	_, err = tx.Exec(sqlStr1)
	if err != nil {
		tx.Rollback()
		return
	}
	_, err = tx.Exec(sqlStr2)
	if err != nil {
		tx.Rollback()
		return
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return
	}
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init DB failed, error:%v\v", err)
		return
	}
	defer db.Close()
	fmt.Println("连接数据库成功")
	// queryOne(2)
	// queryMore(1)
	// insert("孙七", 886)
	// insert("吴八", 80000)
	// update(3, 9000)
	// delete(2)
	// queryMore(0)
	data := map[string]int{
		"赵大": 10,
		"钱乙": 15,
		"郑丙": 33,
	}
	prepareInsert(data)
}
