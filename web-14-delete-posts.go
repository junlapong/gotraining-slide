package main

import (
	"context"
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql" // import with _ for only register driver
)

func main() {
	ctx := context.Background()
	check := func(err error) {
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
	}
	// Open connection to MySQL Database
	db, err := sql.Open("mysql", "root:@/blogdb")
	check(err)

	// BEGIN DELETE OMIT
	stmt, err := db.PrepareContext(ctx, "DELETE FROM posts WHERE id = ?")
	check(err)

	_, err = stmt.ExecContext(ctx, 1)
	check(err)
	// END DELETE OMIT
}
