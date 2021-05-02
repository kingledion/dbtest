package data

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type GroupCount struct {
	Cohort string
	Cnt    uint
}

func GroupByQuery(db *sql.DB) ([]GroupCount, error) {

	var counts []GroupCount

	res, err := db.Query(`
		SELECT cohort, COUNT(cohort) AS cnt
		FROM test_table
		WHERE cohort = "group0"
		GROUP BY cohort
	`)
	defer res.Close()

	if err != nil {
		log.Println(err)
		return []GroupCount{}, err
	}

	for res.Next() {
		var gc GroupCount
		err := res.Scan(&gc.Cohort, &gc.Cnt)
		if err != nil {
			return []GroupCount{}, err
		}

		counts = append(counts, gc)
	}

	return counts, nil

}
