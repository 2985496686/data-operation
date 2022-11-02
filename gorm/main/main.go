package main

import (
	"data-operation/gorm/model"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
)

// 批量插入
func InsertUsers(db *gorm.DB, studs *[]model.Student) {
	db.Create(studs)
}

func main() {
	//数据库连接
	//"user:password@/dbname?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open("root:111111@/study2?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		fmt.Println("数据库连接失败！")
		return
	}
	stu := model.Student{
		UserName: "张三",
		ClassNo:  1,
		Score:    95,
	}
	//插入数据
	//result := db.Create(&stu)
	/*
		只插入UserName和ClassNo两个字段
		db.Select("UserName", "ClassNo").Create(&stu)
		不插入UserName和ClassNo两个字段
		db.Omit("UserName", "ClassNo").Create(&stu)
	*/
	//err = result.Error
	if err != nil {
		fmt.Println(err.Error())
	}

	//批量插入 方法一
	//studs := []model.Student{{UserName: "张三", ClassNo: 2, Score: 87},
	//	{UserName: "李四", ClassNo: 1, Score: 67},
	//	{UserName: "王五", ClassNo: 3, Score: 97}}
	//InsertUsers(db, &studs)

	//批量插入 方法二 通过map批量插入
	//db.Model(&model.Student{}).Create([]map[string]interface{}{
	//	{"UserName": "李四", "ClassNo": 1, "Score": 67},
	//	{"UserName": "王硕", "ClassNo": 2, "Score": 97},
	//	{"UserName": "宇腾", "ClassNo": 3, "Score": 87},
	//})

	//save操作必须指定主键，用于修改指定主键的用户信息，若主键不存在，则会进行插入
	//db.Save(stu)

	//跳过钩子方法执行插入
	db.Session(&gorm.Session{SkipHooks: true}).Create(&stu)

	db.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "id"}},
	})
}
