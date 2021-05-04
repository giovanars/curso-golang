package main

import (
	"log"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func main()  {
	db, err := sql.Open("mysql", "root:123456@/cursogo")
	if err != nil {
		panic(err)
	}
	defer db.Close();

	tx, _ := db.Begin()
	stmt, _ := db.Prepare("insert into usuarios(id, nome) values(?,?)")

	stmt.Exec(2020, "Bia")
	stmt.Exec(2003, "Carlos")

	_, err = stmt.Exec(1, "Tiago")

	if err != nil{
		tx.Rollback()
		log.Fatal(err)
	}

	tx.Commit()
}