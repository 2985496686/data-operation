package test

import (
	"data-operation/gorm/model"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"testing"
)

func TestSelect(t *testing.T) {
	db, err := gorm.Open(mysql.Open("root:111111@/study2?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		fmt.Println("数据库连接错误:" + err.Error())
	}
	var stu1 model.Student
	//First 会按照主键升序排序取第一个 查询相当于: SELECT * FROM `u_stu` ORDER BY `u_stu`.`id` LIMIT 1
	db.First(&stu1)
	fmt.Println(stu1)

	var stu2 model.Student
	//Last 会按照主键降序取第一个 相当于：SELECT * FROM `u_stu` ORDER BY `u_stu`.`id` DESC LIMIT 1
	db.Last(&stu2)
	fmt.Println(stu2)

	var stu3 model.Student
	//Take 不会进行排序，直接取记录中的第一条 相当于：SELECT * FROM `u_stu` LIMIT 1
	db.Take(&stu3)
	fmt.Println(stu3)

	//Find 相当于SELECT * FROM `u_stu`
	var users1 []model.Student
	db.Find(&users1)
	fmt.Println(users1)

	//根据条件查询
	//方法一:使用Take、First、Take时，指定user的相应字段 注意：默认的零值不能被当作筛选条件
	//方法二:使用Find查询
	var users2 []model.Student
	db.Find(&users2, "classno = ?", 1)
	fmt.Println(users2)
	//方法三: 使用Where
	var users3 []model.Student
	db.Where("classno = ?", 1).Find(&users3)
	fmt.Println(users3)
	//方法四：使用Where & Map
	var users4 []model.Student
	db.Where(map[string]interface{}{"classno": 1}).Find(&users4)
	fmt.Println(users4)

	//NOT 相当于：SELECT * FROM `u_stu` WHERE NOT classno = 1
	var users5 []model.Student
	db.Not("classno = ?", 1).Find(&users5)
	fmt.Println(users5)

	//limit和find：SELECT * FROM `u_stu` LIMIT 3 OFFSET 2
	var users6 []model.Student
	db.Offset(2).Limit(3).Find(&users6)
	fmt.Println(users6)

	//注意：想要将查询结果Scan或Find进入该结构体，要满足:1.结构体字段首字母大写  2. 结构体字段要与查询字段相同
	var result []struct {
		Classno int
		Max     int
	}
	//Group、Having
	db.Table("u_stu").
		Select("classno,max(score) max").
		Group("classno").
		Having("max(score) > ?", 95).Scan(&result)
	fmt.Println(result)

	var users7 model.Student
	db.Distinct("username", "classno").Find(&users7)
	fmt.Println(users7)

	//高级查询
	//FirstOrInit  获取升序排列的第一条记录，若获取失败，通过筛选条件初始化结构体
	var stu4 model.Student
	db.Attrs("classno", 1). //当获取失败时，Attrs会进行属性赋值
				Assign("score", 100). //无论获取成功还是失败，Assign都会进行属性赋值
				FirstOrInit(&stu4, map[string]interface{}{"username": "Joke"})
	fmt.Println(stu4)

	//FirstOrCreate 获取升序排序的第一条记录，若获取失败，则创建记录
	// Attrs和Assign的用法和FirstOrInit中类似
	var stu5 model.Student
	db.Attrs("classno", 1).
		Assign("score", 100).
		FirstOrCreate(&stu5, map[string]interface{}{"username": "Joke"})

}
