package model

import (
	"hello/main/todo/views"
)

//ReadAll to find database by no id means all db
func ReadAll() ([]views.PostRequest, error) {
	rows, err := con.Query("SELECT * FROM TODO")
	if err != nil {
		return nil, err
	}
	todos := []views.PostRequest{}
	for rows.Next() {
		data := views.PostRequest{}
		rows.Scan(&data.Name, &data.ToDo)
		todos = append(todos, data)
	}
	return todos, nil
}

//ReadByName created to find database by id as name
func ReadByName(name string) ([]views.PostRequest, error) {
	rows, err := con.Query("SELECT * FROM TODO WHERE name=?", name)
	if err != nil {
		return nil, err
	}
	todos := []views.PostRequest{}
	for rows.Next() {
		data := views.PostRequest{}
		rows.Scan(&data.Name, &data.ToDo)
		todos = append(todos, data)
	}
	return todos, nil
}
