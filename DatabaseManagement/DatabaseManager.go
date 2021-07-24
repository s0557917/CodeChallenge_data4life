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

//This function opens a connection to the database and saves it globally
func ConnectToDatabase() {
	//A connection to the local sqlite database is opened using some flags which improve performance
	//Journal Mode disables the rollback journal, which means no rollback is possible in case of an error.
	//Given that the data in this excercise is meaningless and the aim is to improve performance, this flag makes sense
	//Exclusive locking mode means that the program never releases the lock on the database file.
	//As in this case, there is only one instance writing to the database, concurrency errors cannot not happen.
	//Turning synchronicity off means that if the computer crashes, the database may be corrupted. In this case
	//as the data is not vital, the tradeoff for performance makes sense
	//There are multiple other flags which can be set, I just chose some which seemed to make the most
	//sense in this case, as I just wanted to try out their hit on performance.
	database, err = sql.Open("sqlite3", "./tokens.db?_journal_mode=OFF&_locking_mode=EXCLUSIVE&_synchronous=OFF")
	database.Exec("DELETE FROM tokens")

	if err != nil {
		log.Fatalf("Error while connecting to the database! %s", err)
	}

	//A transaction is created, which will contain all the writes that will be performed.
	//By doing this, the amount of times the connection to the database is reduced to only once.
	//The risk, is that if something happens to the transaction all the data might be gone or corrupted,
	//but as stated before, in this case the data is not vital and the tradeoff makes sense
	transaction, err = database.Begin()
}

//This function gets called eveytime a new token is read, which gets passed as a parameter.
//The function then adds this insertion statement and the value (the token) to the transaction.
//The statement is a constant again, for the same reason explained earlier
func WriteTokensToDatabase(newToken string) {
	transaction.Exec(statement, newToken)
}

//When all the data has been read, the transaction has to be either commited or nothing gets done
//(Journal Mode is set to off, so rolling back is not possible).
//Then, the database connection is closed
func CloseDatabaseConnection() {
	if err != nil {
		fmt.Println("ROLLBACK")
	} else {
		fmt.Println("COMMIT")
		transaction.Commit()
	}
	database.Close()
}
