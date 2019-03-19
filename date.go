package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/date", func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		w.Write([]byte(now.Format("2006/01/02\n")))
	})
	log.Fatal(http.ListenAndServe("0.0.0.0:3112", nil))
}
