package main

import (
	"fmt"
	"net/http"
	"waysbuck/database"
	"waysbuck/pkg/mysql"

	"github.com/gorilla/mux"
)

func main() {

	//. initial DB
	mysql.DatabaseInit()

	//. run migration
	database.RunMigration()

	r := mux.NewRouter()

	port := 5000

	fmt.Println("server running at port ", port)
	http.ListenAndServe("localhost:5000", r)
}
