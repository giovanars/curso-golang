package main

import (
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"strings"
	"strconv"
	"fmt"
	"database/sql"
	"log"
	"json"
)


type Usuario struct {
	ID int `json:"id"`
	Nome string `json:"nome"`
}


func UsuarioHandler(w http.ResponseWriter, r *http.Request) {

	sid := strings.TrimPrefix(r.URL.Path, "/usuarios/")
	id, _ := strconv.Atoi(sid)

	switch {
	case r.Method ==  "GET" && id > 0:
		usuarioPorId(w, r, id);
	case r.Method == "GET":
		usuarioTodos(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Desculpa ..")
	}

}

func usuarioPorId(w http.ResponseWriter, r *http.Request, id int){
	db, err := sql.Open("mysql", "root:123456@/cursogo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var u Usuario
	db.QueryRow("SELECT id, nome FROM usuarios WHERE id = ?", id).Scan(&u.ID, &u.Nome)

	json, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(json))
}


func usuarioTodos(w http.ResponseWriter, r *http.Request){
	db, err := sql.Open("mysql", "root:123456@/cursogo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, _ := db.Query("SELECT id, nome FROM usuarios")
	defer rows.Close()

	var u []Usuario
	for rows.Next() {
		var usuario Usuario
		rows.Scan(&usuario.ID, &Usuario.Nome)
		u.append(u, usuario)
	}

	json, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(json))
}