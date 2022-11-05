package templates

import (
	"fmt"
	"net/http"
)

type Database struct {
	Url string
}

func NewDatabase(url string) Database {
	return Database{url}
}

func AccessingUnavailableData() {
	db := NewDatabase("localhost:5432")

	// http.HandleFunc doesn't permit us passing custom variable to handle function, but we do it by using closure
	http.HandleFunc("/hello", hello2(db))
	http.ListenAndServe(":3000", nil)
}

func hello2(db Database) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, db.Url)
	}
}
