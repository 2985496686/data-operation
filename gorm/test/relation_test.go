package test

import (
	"data-operation/gorm/model"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"testing"
)

func TestRelation(t *testing.T) {

	db, err := gorm.Open(mysql.Open("root:111111@/study2?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		fmt.Println("数据库连接错误:" + err.Error())
	}
	stu1 := model.Student{
		UserName: "Joke",
		ClassNo:  4,
		Score:    98,
		Class:    model.Class{ClassId: 5, ClassName: "高三九班"},
	}
	//该操作会先插入外键指向的实体:Class，再插入Student
	//下面两个操作虽然允许插入外部关联的实体，但是如果外部关联的实体存在的话并不会更新
	//db.Create(&stu1)
	//db.Save(&stu1)

	//若想要更新外部关联的数据，使用FullSaveAssociations
	db.Session(&gorm.Session{FullSaveAssociations: true}).Save(&stu1)

	//在插入时若不想插入关联实体，如下：
	db.Omit("Class").Save(&stu1)
}
