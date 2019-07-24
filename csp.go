package main

import (
	"flag"
	"net/http"
	"os"
    "fmt"
)

func main() {
	csp := flag.String("csp", "", "Content-Security-Policy")
	flag.Parse()

	fileServer := http.FileServer(http.Dir("."))
	http.Handle("/", setCSPHeader(*csp, fileServer))

    fmt.Println("🚀 http://localhost:1213")
    fmt.Println("https://content-security-policy.com/")
	if err := http.ListenAndServe(":1213", nil); err != nil {
		os.Exit(1)
	}
}

func setCSPHeader(csp string, h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Security-Policy", csp)
		h.ServeHTTP(w, r)
	}
}
