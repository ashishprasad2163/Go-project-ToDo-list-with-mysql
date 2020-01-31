package main

import (
	"fmt"
	"hello/main/todo/controller"
	"hello/main/todo/model"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	mux := controller.Register()
	db := model.Connect()
	defer db.Close()
	log.Fatal(http.ListenAndServe(":3000", mux))
	fmt.Println("Serving...")

}
