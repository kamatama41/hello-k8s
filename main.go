package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		sleep, ok := r.URL.Query()["sleep"]
		if ok {
			if d, err := time.ParseDuration(sleep[0]); err == nil {
				log.Printf("Sleep %s", d)
				time.Sleep(d)
			}
		}
		num, ok := r.URL.Query()["num"]
		if !ok {
			num = []string{"-"}
		}

		fmt.Fprintf(w, "Hello World, %s.", num[0])
	})
	log.Println("Starting server...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
