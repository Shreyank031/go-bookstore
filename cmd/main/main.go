package main

import (
	"github.com/Shreyank031/go-bookstore/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRouter(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9091", r))
}
