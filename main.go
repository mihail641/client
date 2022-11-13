package main

import (
	"example.com/kate/adapter"
	"example.com/kate/config"
	"example.com/kate/controller"
	"flag"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	var c config.Config
	c = config.Get()
	config.FaultAdapterTypeUrl(c)
	flagComand := c.ConcreteAdapterType
	//получение и считывание значения флага, возможные значения флага берутся из adapterType
	var concreteAdapterType string
	flag.StringVar(&concreteAdapterType, "concreteAdapterType", string(flagComand), "")
	flag.Parse()
	var p adapter.AdapterType
	//присваивание считанного значения флага структуре adapterType
	p = adapter.AdapterType(concreteAdapterType)
	//adType:=make([]string,0)
	//adType=append(adType, adapter.FileAdapterType)
	//adType=append(adType,adapter.DataBaseAdapterType)
	//запуск роутера
	router := mux.NewRouter()
	//router.HandleFunc регистрация первого маршрута, с URL оканчивающимся на "/users" и методом GET, созадет новый экземпляр конструктора
	//контроллера с аргументом DB, прием-передача параметров функции контроллера Getusers
	router.HandleFunc("/do", func(res http.ResponseWriter, req *http.Request) {
		//userCtrl := controller.NewUserCtrl()
		con := controller.NewController(p)
		con.HandleHttp(res, req)
	}).Methods("GET")
	log.Fatal(http.ListenAndServe(":5000", router))

}
