package controller

import (
	"net/http"
)

//Register function
func Register() *http.ServeMux {

	mux := http.NewServeMux()
	mux.HandleFunc("/ping", ping())
	mux.HandleFunc("/", create())
	return mux
}
