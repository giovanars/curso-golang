package main

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
)

type usuario struct {
	id int
	nome string
}

func main(){
	db, err := sql.Open("mysql", "root:123456@/cursogo")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, _ := db.Query("SELECT * FROM usuarios where id > ?", 5)
	defer rows.Close()

	for rows.Next() {
		var u usuario
		rows.Scan(&u.id, &u.nome)
		fmt.Println(u)
	}

	
}