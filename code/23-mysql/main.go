package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type user struct {
	id       int
	username string
	password string
}

// 更新用户信息
func updateOne(u *user) bool {
	if u == nil {
		log.Fatal(errors.New("user pointer is nil"))
	}
	sql := "update user_tbl set username = ?, password = ? where id = ?"
	r, err := db.Exec(sql, u.username, u.password, u.id)
	check(err)
	i, err := r.RowsAffected()
	check(err)
	return i == 1
}

// 删除用户信息
func deleteOne(id int) bool {
	sql := "delete from user_tbl where id = ?"
	r, err := db.Exec(sql, id)
	check(err)
	i, err := r.RowsAffected()
	check(err)
	return i == 1
}

// 插入一个用户，插入成功则返回插入之后用户的id
func insertOne(u *user) int64 {
	if u == nil {
		log.Fatal(errors.New("user pointer is nil"))
	}
	sql := "insert into user_tbl values(null, ?, ?)"
	r, err := db.Exec(sql, u.username, u.password)
	check(err)
	i, err := r.LastInsertId()
	check(err)
	return i
}

// 查询单行数据
func queryRow(id int) *user {
	u := user{}
	// sql语句，参数使用占位符
	sql := "select id, username, password from user_tbl where id = ?"
	// 调用QueryRow的时候，并不会有任何的报错信息
	// 报错信息会延迟到Scan方法被调用才返回
	err := db.QueryRow(sql, id).Scan(&u.id, &u.username, &u.password)
	check(err)
	return &u
}

// 查询多行数据
func queryRows(id int) []user {
	users := []user{}
	sql := "select * from user_tbl where id > ?"
	r, err := db.Query(sql, id)
	check(err)
	// 延迟关闭数据库连接
	defer r.Close()
	for r.Next() {
		var u user
		err = r.Scan(&u.id, &u.username, &u.password)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, u)
	}
	return users
}

// 初始化数据库连接
func initDb() error {
	str := "root:cyj070723@tcp(localhost:3306)/go_db?charset=utf8mb4"
	conn, err := sql.Open("mysql", str)
	check(err)
	db = conn
	// 最大连接时长
	db.SetConnMaxLifetime(time.Minute * 3)
	// 最大连接数
	db.SetMaxOpenConns(10)
	// 空闲连接数
	db.SetMaxIdleConns(10)
	return db.Ping()
}

func main() {
	err := initDb()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("成功连接数据库")
	fmt.Println("========================")
	fmt.Println("查询id为1的用户:")
	u := queryRow(1)
	fmt.Println(u)
	fmt.Println("========================")
	fmt.Println("查询id>0的用户:")
	users := queryRows(0)
	fmt.Println(users)
	fmt.Println("========================")
	fmt.Println("插入一个用户:")
	// u1 := user{username: "张三", password: "123456"}
	// i := insertOne(&u1)
	// fmt.Println(i)
	fmt.Println("========================")
	fmt.Println("修改id为3的用户的用户名为李四:")
	u2 := user{id: 3, username: "李四", password: "123456"}
	if b := updateOne(&u2); b {
		fmt.Println("修改成功")
	} else {
		fmt.Println("修改失败")
	}
	fmt.Println("========================")
	fmt.Println("删除id为3的用户:")
	if b := deleteOne(3); b {
		fmt.Println("删除成功")
	} else {
		fmt.Println("删除失败")
	}
}
