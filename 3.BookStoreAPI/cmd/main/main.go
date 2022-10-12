package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/mehrdad-zade/GoLang/3.BookStoreAPI/pkg/routs"
)

func main() {
	r := mux.NewRouter()
	routs.RegisterBookStoreRouts(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8080", r))

}

