package main

import (
	"database/sql"
	"fmt"
	//"database/sql/driver"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

const (
	sqlUser = "root"
	sqlPwd  = "lrf64787"
)

// root:123456%^@tcp(39.107.87.114:3306)/book_manager?charset=utf8
func initSql() (db *sql.DB, err error) {
	db, err = sql.Open("mysql", sqlUser+":"+sqlPwd+"@tcp(127.0.0.1:3306)/beego?charset=utf8")
	if err != nil {
		err = errors.Wrap(err, "sql Open fail")
		return
	}
	//defer db.Close()
	err = db.Ping()
	if err != nil {
		err = errors.Wrap(err, "sql link fail")
		return
	}

	return
}

type account struct {
	Id       int64
	carid    string
	status   int8
	password string
}

func queryAc(db *sql.DB, sqlStr string) (err error) {
	fmt.Println("sqlStr:", sqlStr)
	rows, err := db.Query(sqlStr)
	if err != nil {
		return errors.Wrap(err, "db query fail")
	}

	for rows.Next() {
		var ac account
		err = rows.Scan(&ac.Id, &ac.carid, &ac.status, &ac.password)
		if err != nil {
			switch {
			case err == sql.ErrNoRows:
				err = errors.Wrap(err, "ErrNoRows")
			default:
				err = errors.Wrap(err, "scan error")
			}
			return
		}
		fmt.Println(ac)
	}

	fmt.Println("queryAc finished")

	return
}
