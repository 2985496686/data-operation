package main

import (
	"data-operation/mysql/entity"
	"data-operation/mysql/operation"
	"fmt"
)

func main() {
	db, err := operation.ConnectMysql("root", "111111", "localhost", "3306", "study2")
	/*
		transaction, err := db.Begin() ------- 开启事务
		transaction.Rollback() ----------- 回滚事务
		transaction.Commit() ------- 提交
	*/

	if err != nil {
		fmt.Println("error:connect failed:" + err.Error())
		return
	}
	id, count := operation.InsertUser(db, "张三")
	fmt.Printf("主键王硕id:%d  影响行数:%d\n", id, count)

	//查询数据库内容
	var users []entity.User
	operation.SelectUserById(&users, db, 1)
	fmt.Println(users)
}
