package main

import (
	"log"
	"net/http"

	"go-server/handler"

	_ "github.com/go-sql-driver/mysql"
)

type Employee struct {
	Id   int
	Name string
	City string
}

func main() {
	log.Println("Server started on: http://localhost:8080")
	http.HandleFunc("/", handler.Index)
	http.HandleFunc("/show", handler.Show)
	http.HandleFunc("/new", handler.New)
	http.HandleFunc("/edit", handler.Edit)
	http.HandleFunc("/insert", handler.Insert)
	http.HandleFunc("/update", handler.Update)
	http.HandleFunc("/delete", handler.Delete)
	http.ListenAndServe(":8080", nil)
}
