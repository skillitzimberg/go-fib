package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// Index handles requests at "/api".
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\nTo generate the Fibonacci sequence use the URL '/api/fibonnaci/{someNum} where someNum is an integer greater than 0.\n")
}

// Fibonacci handles requests at "/api/fibonacci/:number".
func Fibonacci(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	num, err := strconv.Atoi(ps.ByName("number"))
	check(err, "strconv.Atoi(ps.ByName('number')")

	fib, err := fib(num)
	check(err, "fib(num)")

	fibJSON, err := json.Marshal(fib)
	check(err, "json.Marshal(fib)")

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%v\n", string(fibJSON))
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/hello/:name", Fibonacci)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func check(err error, from string) {
	if err != nil {
		log.Printf("Error from %s: %s", from, err)
	}
}
