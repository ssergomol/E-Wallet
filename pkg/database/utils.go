package database

import (
	"database/sql"
	"io/ioutil"
	"strings"
)

func LoadSQLFile(db *sql.DB) error {
	file, err := ioutil.ReadFile("pkg/database/init_db.sql")

	if err != nil {
		return err
	}

	requests := strings.Split(string(file), ";")

	for _, request := range requests {
		_, err := db.Exec(request)
		if err != nil {
			return err
		}
	}
	return nil
}
