package mysql

import (
	"database/sql"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"os"
)

var (
	Ect *sql.DB
)

func init() {
	params := strings.Join(
		[]string{
			"parseTime=true",
			"loc=America%2FSao_Paulo",
			"multiStatements=true",
		}, "&",
	)
	var err error
	host, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	if host == "Go" {
		Ect, err = sql.Open("mysql", "root:1000sonhosreais@tcp(localhost:3306)/ect?"+params)
	} else {
		Ect, err = sql.Open("mysql", "root:Semaver13@tcp(localhost:3306)/ect?"+params)
	}
	if err != nil {
		panic(err)
	}
}
