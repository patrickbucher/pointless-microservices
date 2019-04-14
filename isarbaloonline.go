package main

import (
	"log"
	"net/http"
	"os"
)

const url = "http://arbalo.ch"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		resp, err := http.Get(url)
		if err != nil {
			log.Println(err)
			w.Write([]byte("No\n"))
			return
		}
		if resp.StatusCode != http.StatusOK {
			log.Printf("status: %s (%d)", resp.Status, resp.StatusCode)
			w.Write([]byte("No\n"))
			return
		}
		w.Write([]byte("Yes\n"))
	})
	port := os.Getenv("PORT")
	if port == "" {
		port = "1234"
	}
	log.Fatal(http.ListenAndServe("0.0.0.0:"+port, nil))
}
