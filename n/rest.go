package n

import (
	"errors"
	"log"
	"net/http"
)

var page string

// go run .\n\rest.go
// C:\users\retno\appdata\local\temp\go-build3531906157\b001\exe\rest.exe
func StartServer() {
	log.Println("Hello, Go!")
	http.Handle("/welcome", http.HandlerFunc(welcome))
	e := http.ListenAndServe(":8999", nil)
	if e != nil && !errors.Is(e, http.ErrServerClosed) {
		log.Fatal(e)
	}
}

func welcome(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("Welcome to Just Enough Norman"))
	//fmt.Fprint(rw, page)
}
