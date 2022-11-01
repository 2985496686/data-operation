package operation

import (
	"data-operation/mysql/entity"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// ConnectMysql 连接数据库
func ConnectMysql(user, password, ip, post, database string) (*sqlx.DB, error) {

	//database, err := sqlx.Open("数据库类型", "用户名:密码@tcp(地址:端口)/数据库名")
	return sqlx.Open("mysql", user+":"+password+"@tcp("+ip+":"+post+")/"+database)
}

func InsertUser(db *sqlx.DB, args ...any) (int64, int64) {
	result, err := db.Exec("insert into user (name) values(?)", args...)
	if err != nil {
		fmt.Println("error:插入失败！:" + err.Error())
	}
	//获取插入数据后生成的主键id
	id, _ := result.LastInsertId()
	affected, _ := result.RowsAffected()
	return id, affected
}

func SelectUserById(users *[]entity.User, db *sqlx.DB, id int64) {
	err := db.Select(users, "select *from user where id = ?", id)
	if err != nil {
		fmt.Println("error:查询错误！:" + err.Error())
	}
}
