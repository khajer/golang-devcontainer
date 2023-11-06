package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

func indexHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "This is my website!\n")
}
func main() {
	fmt.Println("started")
	http.HandleFunc("/", indexHandle)

	err := http.ListenAndServe(":3000", nil)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}

}
