package main

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
)

func main(){
	db, err := sql.Open("mysql", "root:123456@/cursogo")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	stmt, _:= db.Prepare("update usuarios set nome = ? where id = ?")

	stmt.Exec("Cleiton", 1)
	stmt.Exec("Valeska", 2)
}