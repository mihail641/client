package main

import (
	"example.com/kate/adapterType"
	"example.com/kate/config"
	"example.com/kate/controller"
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	var c config.Config
	c = config.Get()
	flagComand := c.ConcreteAdapterType
	//получение и считывание значения флага, возможные значения флага берутся из adapterType
	var concreteAdapterType string
	flag.StringVar(
		&concreteAdapterType,
		"concreteAdapterType",
		string(flagComand),
		"",
	)
	flag.Parse()
	var p adapterType.AdapterType
	//присваивание считанного значения флага структуре adapterType
	p = adapterType.AdapterType(concreteAdapterType)
	//adType:=make([]string,0)
	//adType=append(adType, adapter.FileAdapterType)
	//adType=append(adType,adapter.DataBaseAdapterType)
	//запуск роутера
	router := mux.NewRouter()
	fmt.Println("А теперь запуск сервера")
	//router.HandleFunc регистрация первого маршрута, с URL оканчивающимся на "/users" и методом GET, создает новый экземпляр конструктора
	//контроллера с аргументом DB, прием-передача параметров функции контроллера Getusers
	router.HandleFunc(
		"/do",
		func(res http.ResponseWriter, req *http.Request) {
			//userCtrl := controller.NewUserCtrl()
			con := controller.NewController(p)
			con.HandleHttp(
				res,
				req,
			)
		},
	).Methods("GET")
	//router.HandleFunc- регистрация маршрута, с URL оканчивающимся на /document и методом GET, создает новый экземпляр конструктора
	//контроллера, прием-передача параметров метода GetSimpleTable по получению пустой таблицы html
	router.HandleFunc(
		"/document",
		func(res http.ResponseWriter, req *http.Request) {
			con := controller.NewDocumentController(p)
			con.GetSimpleTable(
				res,
				req,
			)
		},
	).Methods("GET")
	//router.HandleFunc-регистрация маршрута, с URL оканчивающимся на /complex и методом GET, создает новый экземпляр конструктора
	//контроллера, прием-передача параметров метода GetComplexTable по получению пустой таблицы html со слитыми в определенном порядке ячейками
	router.HandleFunc(
		"/complex",
		func(res http.ResponseWriter, req *http.Request) {
			//userCtrl := controller.NewUserCtrl()
			con := controller.NewDocumentController(p)
			con.GetComplexTable(
				res,
				req,
			)
		},
	).Methods("GET")
	//router.HandleFunc регистрация маршрута, с URL оканчивающимся на /cols/{sizeCols}/rows/{sizeRows} и методом GET, создает новый экземпляр конструктора
	//контроллера, прием-передача параметров метода GetCertainSizeTable по получению пустой таблицы html с заданными через URL количество столбцов и строк
	router.HandleFunc(
		"/cols/{sizeCols}/rows/{sizeRows}",
		func(res http.ResponseWriter, req *http.Request) {
			//userCtrl := controller.NewUserCtrl()
			con := controller.NewDocumentController(p)
			con.GetCertainSizeTable(
				res,
				req,
			)
		},
	).Methods("GET")
	//router.HandleFunc регистрация маршрута, с URL оканчивающимся на /documentation и методом GET, создает новый экземпляр конструктора
	//контроллера, прием-передача параметров метода GetDocumentationTable по получению таблицы html с документами, модулями и ошибками
	log.Println("Starting HTTP server on :5000")
	log.Fatal(
		http.ListenAndServe(
			":5000",
			router,
		),
	)

}
