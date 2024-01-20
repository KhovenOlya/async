package main

import (
	"buildpermits/internal/pkg/api"
	"net/http"
)

func main() {
	http.HandleFunc("/security-decision", api.Check)
	http.ListenAndServe("localhost:8080", nil)
}
