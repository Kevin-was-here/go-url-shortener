package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var string PSQL_PASSWORD := os.Getenv("PSQL_PASSWORD")

const(
	host = "localhost"
	port = 5432
	user = "kloak"
	password = PSQL_PASSWORD
	dbname = "url_shortener" 
)

func main() {
	
	fmt.Println("Hello, World!")
}
