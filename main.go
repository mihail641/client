package main

import (
	"example.com/kate/adapter"
	"example.com/kate/controller"
	"flag"
	"github.com/gorilla/mux"
	"log"
	//"log"
	"net/http"
)

var AdapterType string

func main() {
	flag.StringVar(&AdapterType, "AdapterType", adapter.FileAdapterType, adapter.DataBaseAdapterType)
	flag.Parse()

	//запуск роутера
	router := mux.NewRouter()
	//router.HandleFunc регистрация первого маршрута, с URL оканчивающимся на "/users" и методом GET, созадет новый экземпляр конструктора
	//контроллера с аргументом DB, прием-передача параметров функции контроллера Getusers
	router.HandleFunc("/do", func(res http.ResponseWriter, req *http.Request) {
		//userCtrl := controller.NewUserCtrl()
		con := controller.NewController(AdapterType)
		con.HandleHttp(res, req)
	}).Methods("GET")
	log.Fatal(http.ListenAndServe(":5000", router))
}
