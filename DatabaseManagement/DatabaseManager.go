package DatabaseManagement

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var database *sql.DB
var err error
var transaction *sql.Tx

const statement = "INSERT OR IGNORE INTO tokens (token_value) VALUES (?)"

func ConnectToDatabase() {
	database, err = sql.Open("sqlite3", "./tokens.db?_journal_mode=OFF&_locking_mode=EXCLUSIVE&_synchronous=OFF")
	database.Exec("DELETE FROM tokens")

	if err != nil {
		log.Fatalf("Error while connecting to the database! %s", err)
	}

	transaction, err = database.Begin()
}

func WriteTokensToDatabase(newToken string) {
	transaction.Exec(statement, newToken)
}

func CloseDatabaseConnection() {
	if err != nil {
		fmt.Println("ROLLBACK")
		transaction.Rollback()
	} else {
		fmt.Println("COMMIT")
		transaction.Commit()
	}
	database.Close()
}
