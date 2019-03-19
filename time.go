package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/time", func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		w.Write([]byte(now.Format("15:04:05 CET\n")))
	})
	log.Fatal(http.ListenAndServe("0.0.0.0:2359", nil))
}
