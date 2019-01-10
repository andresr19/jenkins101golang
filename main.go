package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Sum func
func Sum(a, b int) int {
	return a + b
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("%d", Sum(9929, 3))))
}

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/", index)
	http.ListenAndServe(":7777", router)

}
