package main

import (
	"database/sql"
	"fmt"

	// 當 _ "github.com/lib/pq" 有錯誤時，先執行以下語法試試：
	// go get github.com/lib/pq

	// 若仍然有錯誤，則依序執行以下語法：
	// go mod init pq
	// go mod edit -require github.com/lib/pq
	// go get github.com/lib/pq
	// go run main.go

	// 使用 _ 可以避免 Go 在編譯時，提示未使用的錯誤
	_ "github.com/lib/pq"
)

const (
	// Initialize connection constants.
	HOST     = "chmpostgresql.postgres.database.azure.com"
	DATABASE = "vote"
	USER     = "myadmin@chmpostgresql"
	PASSWORD = "Pa55w0rd"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// Initialize connection string.
	var connectionString string = fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=require", HOST, USER, PASSWORD, DATABASE)

	// Initialize connection object.
	db, err := sql.Open("postgres", connectionString)
	checkError(err)

	err = db.Ping()
	checkError(err)
	fmt.Println("Successfully created connection to database")

	// Drop previous table of same name if one exists.
	_, err = db.Exec("DROP TABLE IF EXISTS inventory;")
	checkError(err)
	fmt.Println("Finished dropping table (if existed)")

	// Create table.
	_, err = db.Exec("CREATE TABLE inventory (id serial PRIMARY KEY, name VARCHAR(50), quantity INTEGER);")
	checkError(err)
	fmt.Println("Finished creating table")

	// Insert some data into table.
	sql_statement := "INSERT INTO inventory (name, quantity) VALUES ($1, $2);"
	_, err = db.Exec(sql_statement, "banana", 150)
	checkError(err)
	_, err = db.Exec(sql_statement, "orange", 154)
	checkError(err)
	_, err = db.Exec(sql_statement, "apple", 100)
	checkError(err)
	fmt.Println("Inserted 3 rows of data")
}
