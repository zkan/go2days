package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {

	hand := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		fmt.Fprint(os.Stdout, "processing request\n")

		<-ctx.Done()

		fmt.Fprint(os.Stderr, "request cancelled\n")
	})
	s := &http.Server{
		Addr:           ":8000",
		Handler:        hand,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		IdleTimeout:    10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
