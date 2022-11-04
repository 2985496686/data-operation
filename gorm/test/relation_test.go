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
	fmt.Println(stu1)
	//该操作会先插入外键指向的实体:Class，再插入Student
	//下面两个操作虽然允许插入外部关联的实体，但是如果外部关联的实体存在的话并不会更新
	//db.Create(&stu1)
	//db.Save(&stu1)

	//若想要更新外部关联的数据，使用FullSaveAssociations
	//db.Session(&gorm.Session{FullSaveAssociations: true}).Save(&stu1)

	//在插入时若不想插入关联实体，如下：
	//db.Omit("Class").Save(&stu1)

	//关联查询

	//多对一查询或一对一查询
	//通过一个stu来查询与他关联的class
	//Model方法用来条件筛选，Association方法的参数是需要关联查询的字段
	var stu2 model.Student
	db.Where("id = ?", 5).Find(&stu2)
	db.Model(&stu2).Association("Class").Find(&stu2.Class)
	fmt.Println(stu2)

	//一对多查询
	var class model.Class
	db.Where("class_name", "高三三班").Find(&class)
	db.Model(&class).Association("Students").Find(&class.Students)
	//db.Where("class_name = ?", "高三三班").Preload("Students").Find(&class)
	fmt.Println(class)

	//预加载的方式进行关联查询
	//默认情况下gorm框架是开启了预加载的
	//多对一查询
	var stu3 model.Student
	db.Where("id = 5").Preload("Class").Find(&stu3)
	fmt.Println(stu3)

	//一对多查询
	var class2 model.Class
	db.Where("class_name = ?", "高三三班").Preload("Students").Find(&class2)
	fmt.Println(class2)

}
