package main

import (
	"example.com/kate/controller"
	"github.com/gorilla/mux"

	"log"
	//"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/do", func(res http.ResponseWriter, req *http.Request) {
		//userCtrl := controller.NewUserCtrl()
		con := controller.NewController()
		con.HandleHttp(res, req)
	}).Methods("GET")
	log.Fatal(http.ListenAndServe(":5000", router))
}
