package main

import (
	"fmt"
)

func main() {
	sqlCon, err := initSql()
	defer sqlCon.Close()
	if err != nil {
		fmt.Println("initSql error:", err)
		fmt.Printf("%+v", err, "\n")
	}

	err = queryAc(sqlCon, "select * from account where id=2")
	if err != nil {
		fmt.Println("query sql error :", err)
		fmt.Printf("%+v", err, "\n")
	}
}
