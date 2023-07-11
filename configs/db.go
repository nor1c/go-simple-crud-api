package config

import (
	"database/sql"
)

func main() {
	db, err := sql.Open("mysql", "root:@/go-simple-crud")

	if err != nil {
		panic(err)
	}
}
