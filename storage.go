package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func ConnectDB() error {
	// 连接字符串格式为：postgres://user:password@host:port/database
	connectionString := "host=localhost port=5432 user=postgres dbname=postgres password=rocketyou1 sslmode=disable"

	// 打开数据库连接
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 检查数据库连接是否成功
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT * FROM demo")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	// 打印查询结果
	for rows.Next() {
		var id int
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, name)
	}
	return nil
}
