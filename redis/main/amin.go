package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func main() {
	//192.168.1.171:6379
	conn, err := redis.Dial("tcp", "192.168.1.171:6379")
	conn.Do("auth", "111111")
	if err != nil {
		fmt.Println("数据库连接错误:" + err.Error())
	}
	defer conn.Close()
	conn.Do("set", "hello", "张三")
	hello, _ := conn.Do("get", "hello")
	fmt.Println(hello)
}
