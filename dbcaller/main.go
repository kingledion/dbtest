package main

import (
	"database/sql"
	"fmt"
	"math/rand"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/kingledion/dbtest/data"
)

func main() {

	// local db connection
	db, err := sql.Open("mysql", "dbuser:dbpass@tcp(localhost:3306)/testdb")
	defer db.Close()

	// declare db table
	db.Exec("DROP TABLE test_table")

	_, err = db.Exec("CREATE TABLE test_table (id INT, cohort VARCHAR(255))")
	if err != nil {
		fmt.Println("got an error", err)
		return
	}

	// insert some data into the table
	n := 100000
	stmt := "INSERT INTO test_table VALUES (%d, %s)"

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < n; i++ {
		j := rand.Intn(5)
		insertStmt := fmt.Sprintf(stmt, i, fmt.Sprintf("\"group%d\"", j))

		_, err := db.Exec(insertStmt)
		if err != nil {
			fmt.Println("got an error", err)
			return
		}
	}

	grp, err := data.GroupByQuery(db)
	if err != nil {
		fmt.Println("got an error", err)
		return
	}

	fmt.Println(grp)

}
