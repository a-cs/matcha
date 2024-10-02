package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {
	dataSourceName := "user=myuser password=mysecretpassword host=localhost port=5432 dbname=mydatabase sslmode=disable"
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// Executar uma consulta e imprimir os resultados
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var email, password, user_name, account_status string
		var active_matchs interface{}
		if err := rows.Scan(&id, &email, &password, &user_name, &active_matchs, &account_status); err != nil {
			panic(err.Error())
		}
		fmt.Println(id, email, password, user_name, active_matchs, account_status)
	}
}
