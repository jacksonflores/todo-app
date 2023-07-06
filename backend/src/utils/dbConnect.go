package utils

import (
	"database/sql"
)

func DBConnect() (*sql.DB, error) {
	db, err := sql.Open("mysql", SettingsStuff.DB)
	if err != nil {
		return nil, err
	}

	return db, nil
}
