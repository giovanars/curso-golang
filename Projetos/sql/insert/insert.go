package main

import (
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"database/sql"
)


func main(){
	db, err := sql.Open("mysql", "root:123456!@/cursogo")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	stmt, _ := db.Prepare("insert into usuarios(nome) values(?)")
	stmt.Exec("Maria")
	stmt.Exec("João")

	res, _ := stmt.Exec("Pedro")

	id, _ := res.LastInsertId()
	lineEffecteds, _ := res.RowsAffected()
	fmt.Println(id)
	fmt.Println(lineEffecteds)
}