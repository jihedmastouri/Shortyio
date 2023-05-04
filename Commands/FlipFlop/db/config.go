package db

import (
	"database/sql"
)


func init()  {
    // Open a database connection
    db, err := sql.Open("<driver-name>", "<connection-string>")
    if err != nil {
        // handle error
    }
    defer db.Close()
}
