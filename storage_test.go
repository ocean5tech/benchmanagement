package main

import (
	"database/sql"
	"fmt"
	"log"
	"testing"

	_ "github.com/lib/pq"
)

func TestDatabaseConnection(t *testing.T) {

	//connectionString := "postgres://postgres:rocketyou1@localhost:5432/postgres?search_path=benchmanagement"
	connectionString := "host=localhost port=5432 user=postgres dbname=postgres password=rocketyou1 sslmode=disable search_path=benchmanagement"
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		t.Errorf("Failed to open database: %v", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		t.Errorf("Failed to connect to database: %v", err)
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
		var PM *string
		var supervisor *string
		var status *string
		var homepage *string
		err := rows.Scan(&id, &name, PM, supervisor, status, homepage)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, name)
	}
	fmt.Println("Connection to database established!")
}
