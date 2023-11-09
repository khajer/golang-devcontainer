package main

import (
	"database/sql"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

const (
	host     = "host.docker.internal" // change localhost to when use devContainer
	port     = 5432
	user     = "postgres"
	password = "mysecretpassword"
	dbname   = "test_db"
)

func indexHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "This is my website!\n")
}
func main() {
	fmt.Println("started")
	http.HandleFunc("/", indexHandle)

	ConnectDB()
	err := http.ListenAndServe(":3000", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}

}
func ConnectDB() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("connect database successful")
}
