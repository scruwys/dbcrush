package main

import (
	"database/sql"
	"fmt"
	"log"

	// we have to import the driver, but don't use it in our code
	// so we use the `_` symbol
	"github.com/go-sql-driver/mysql"
	"github.com/spf13/cobra"
)

func Run(cmd *cobra.Command, args []string) {
	cfg := mysql.Config{
		User:                 flagUsername,
		Passwd:               flagPassword,
		Net:                  "tcp",
		Addr:                 fmt.Sprintf("%s:%d", flagHostname, flagPort),
		DBName:               flagDatabase,
		AllowNativePasswords: true,
	}

	// Get a database handle.
	db, err := sql.Open("mysql", cfg.FormatDSN())

	if err != nil {
		log.Fatalf("could not connect to database: %v", err)
	}

	// To verify the connection to our database instance, we can call the `Ping`
	// method. If no error is returned, we can assume a successful connection
	if err := db.Ping(); err != nil {
		log.Fatalf("unable to reach database: %v", err)
	}

	fmt.Println("database is reachable")
}
