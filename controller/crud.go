package controller

import (
	"encoding/json"
	"fmt"
	"hello/main/todo/model"
	"hello/main/todo/views"
	"net/http"
)

func create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			data := views.PostRequest{}
			//deserialize the data or decoding
			json.NewDecoder(r.Body).Decode(&data)
			fmt.Println(data)
			if err := model.CreateToDo(data.Name, data.ToDo); err != nil {
				w.Write([]byte("Some error"))
				return
			}
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(data)
		} else if r.Method == http.MethodGet {
			name := r.URL.Query().Get("name")
			data, err := model.ReadByName(name)
			if err != nil {
				w.Write([]byte(err.Error()))
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(data)
		} else if r.Method == http.MethodDelete {
			//
			name := r.URL.Path[1:]
			fmt.Println(name)
			if err := model.DeleteToDo(name); err != nil {
				w.Write([]byte("Some error"))
				return
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(struct {
				Status string `json:status`
			}{"Item deleted."})
		}
	}
}
