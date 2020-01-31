package main

import (
	"fmt"
	"hello/main/todo/controller"
	"hello/main/todo/model"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	mux := controller.Register()
	db := model.Connect()
	defer db.Close()
	port := os.Getenv("PORT")
	log.Fatal(http.ListenAndServe(":"+port, mux))
	fmt.Println("Serving...")

}
