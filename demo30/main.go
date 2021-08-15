package main

import (
	myMod "./mod"
	. "./foo"
	"./db"
	"fmt"
)

func main() {
	fmt.Println("start main")
	
	// 引用db包
	fmt.Println(db.Connection)
	db.ConnectDB()
	db.CloseDB()

	// 引用mod包
	fmt.Println(myMod.Msg)

	// 引用foo包
	fmt.Println(Count)
}