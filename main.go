package main

import (
	"example.com/kate/controller"
	"github.com/gorilla/mux"

	"log"
	//"log"
	"net/http"
)

const (
	URL = "http://127.0.0.1:4000/user"
	URLGET="http://127.0.0.1:4000/users"


)
type App struct {
	client *controller.Controller
}

func main() {
	app := App{
		client: controller.NewController(),
	}
	app.StartServer()
}
func (app *App) StartServer () {
	router := mux.NewRouter()
	router.HandleFunc("/do", func(res http.ResponseWriter, req *http.Request) {
		//userCtrl := controller.NewUserCtrl()
		con := controller.NewController()
		con.HandleHttp(res, req)
	}).Methods("GET")

	//router.HandleFunc("/do", controller.HandleHttp).Methods("GET")
	//log.Fatal(http.ListenAndServe(":5000", router))

	log.Fatal(http.ListenAndServe(":5000", router))
	}


