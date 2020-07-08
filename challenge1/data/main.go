package main

import (
	"fmt"
	"net/http"
	"os"
	"sync/atomic"
)

// We're going to use a mutex so that we can handle concurrent increments
// We can do only this because this app will only have 1 replica,
// and doesn't have to provide any persistence guarantees.
var clicks uint64

func main() {
	http.HandleFunc("/clicks", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			atomic.AddUint64(&clicks, 1)
		}

		fmt.Fprint(w, atomic.LoadUint64(&clicks))
	})

	var addr = fmt.Sprintf(":%s", os.Getenv("PORT"))
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		panic(err)
	}
}
