package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	var status = "OK"

	http.HandleFunc("/api/service2/untested-request", func(w http.ResponseWriter, r *http.Request) {
		status = "FAILED"
		fmt.Fprintf(w, "Service successfully broken")
	})
	http.HandleFunc("/api/service2", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, status)
		if status != "OK" {
			time.Sleep(10 * time.Second)
		}
		fmt.Fprintf(w, "Hello from go server %s \n", status)
	})
	http.ListenAndServe(":8080", nil)
}
