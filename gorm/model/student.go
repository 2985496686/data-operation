package model

import (
	"fmt"
	"gorm.io/gorm"
)

// Student
// gorm框架默认的约定:
//  1. Student 对应到数据库中的students表
//     我们可以使用方法TableName来设置对应的表名
//  2. 结构体中字段想要对应到数据库表字段，首先字段名首字母必须大写，默认采用蛇形小写，如：UserName 对应 user_name
//     可以使用标签来设置字段名对应关系:`gorm:"column:username"`
//
// /
type Student struct {
	Id       int    `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	UserName string `gorm:"column:username"`
	ClassNo  int    `gorm:"column:classno"`
	Score    int    `gorm:"column:score"`
	Class    Class  `gorm:"references:ClassId;foreignKey:ClassNo"` //foreignKey:该结构体中的外键  references：外联实体的主键或唯一键
}

func (stu Student) TableName() string {
	return "u_stu"
}

//钩子方法，会自动开启是事务，钩子方法返回error时会自动回滚

// BeforeCreate 插入方法(Create方法)执行前执行
func (stu *Student) BeforeCreate(db *gorm.DB) (err error) {
	fmt.Println("BeforeCreate执行！")
	return
}

// AfterCreate 插入方法(Create方法)执行后执行
func (stu *Student) AfterCreate(db *gorm.DB) (err error) {
	fmt.Println("AfterCreate执行！")
	//	return errors.New("error!")
	return
}
