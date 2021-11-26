package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func Init() (*sql.DB, error) {
	// Create the database handle, confirm driver is present
	db, err := sql.Open("mysql", "rabbit:@%")
	return db,err
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
