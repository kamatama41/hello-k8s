package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		sleep, ok := r.URL.Query()["sleep"]
		if ok {
			if d, err := time.ParseDuration(sleep[0]); err == nil {
				time.Sleep(d)
			}
		}
		fmt.Fprintf(w, "Hello, World")
	})
	http.ListenAndServe(":8080", nil)
}
