package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/date", func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		today := now.Format("2006/01/02\n")
		w.Write([]byte(fmt.Sprintf("Today's Date: %s", today)))
	})
	log.Fatal(http.ListenAndServe("0.0.0.0:3112", nil))
}
