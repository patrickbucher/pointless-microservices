package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
)

func main() {
	http.HandleFunc("/dice", func(w http.ResponseWriter, r *http.Request) {
		eyes := rand.Intn(6) + 1
		w.Write([]byte(fmt.Sprintf("%d\n", eyes)))
	})
	log.Fatal(http.ListenAndServe("0.0.0.0:6666", nil))
}
