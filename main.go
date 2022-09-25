package main

import (
	"example.com/kate/controller"
	"github.com/gorilla/mux"

	"log"
	//"log"
	"net/http"
)

func main() {
	//запуск роутера
	router := mux.NewRouter()
	//router.HandleFunc регистрация первого маршрута, с URL оканчивающимся на "/users" и методом GET, созадет новый экземпляр конструктора
	//контроллера с аргументом DB, прием-передача параметров функции контроллера Getusers
	router.HandleFunc("/do", func(res http.ResponseWriter, req *http.Request) {
		//userCtrl := controller.NewUserCtrl()
		con := controller.NewController()
		con.HandleHttp(res, req)
	}).Methods("GET")
	log.Fatal(http.ListenAndServe(":5000", router))
}
