package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/time", func(w http.ResponseWriter, r *http.Request) {
		now := time.Now().Format("15:04:05 CET")
		w.Write([]byte(fmt.Sprintf("Current Time: %s\n", now)))
	})
	log.Fatal(http.ListenAndServe("0.0.0.0:2359", nil))
}
