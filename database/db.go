package database

import (
	"database/sql"
)

func Init() (*sql.DB, error) {
	// Create the database handle, confirm driver is present
	sql.Drivers()

	return sql.Open("mariaDB", "rabbit@%")
	/*
		defer func(db *sql.DB) {
			err := db.Close()
			if err != nil {

			}
		}(db)

		// Connect and check the server version
		var version string
		db.QueryRow("SELECT VERSION()").Scan(&version)
		fmt.Println("Connected to:", version)
	*/

}
