package db

import "fmt"

const (
	Connection string = "127.0.0.1:xxx..."
)

func ConnectDB()  {
	fmt.Println("数据库连接")
}

func CloseDB()  {
	fmt.Println("数据库关闭")
}