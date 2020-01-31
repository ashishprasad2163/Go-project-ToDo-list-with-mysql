package controller

import (
	"encoding/json"
	"fmt"
	"hello/main/todo/views"
	"net/http"
)

func ping() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			data := views.Response{
				Code: http.StatusOK,
				Body: "pong",
			}
			w.WriteHeader(200)
			fmt.Println("request generated.")
			json.NewEncoder(w).Encode(data)
		}
	}
}
