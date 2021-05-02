package data

import (
	"database/sql"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// unindexed
// i => 100    :   0.5 ms/op
// i => 1000   :   2.3 ms/op
// i => 10000  :  20.1 ms/op
// i => 100000 : 215.6 ms/op
// indexed
// i => 100    :  0.2 ms/op
// i => 1000   :  0.7 ms/op
// i => 10000  :  6.0 ms/op
// i => 100000 : 59.6 ms/op
func BenchmarkGroupByQuery(b *testing.B) {

	// local db connection
	db, err := sql.Open("mysql", "dbuser:dbpass@tcp(localhost:3306)/testdb")
	defer db.Close()

	// declare db table
	db.Exec("DROP TABLE test_table")

	_, err = db.Exec("CREATE TABLE test_table (id INT, cohort VARCHAR(255))")
	if err != nil {
		b.Fatal(err)
	}
	_, err = db.Exec("CREATE INDEX idx_cohort ON test_table (cohort)")
	if err != nil {
		b.Fatal(err)
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
			b.Error(err)
		}
	}

	b.ResetTimer()

	// access and print results
	for i := 0; i < b.N; i++ {
		GroupByQuery(db)
	}
}
