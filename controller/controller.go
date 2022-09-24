package controller

import (
	"encoding/xml"
	"example.com/kate/model"
	"fmt"
	//"log"
	"net/http"
)

type Controller struct {
	controller *model.Model
}

func NewController() *Controller {
	return &Controller{
		controller: model.NewModel(),
	}
}
func (m *Controller) HandleHttp(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Сервер запустился")
	var t, err = m.controller.ClientAlgorithmTake()
	if err != nil {
		m := "Ошибка выполнеия крнтроллера: %s"
		fmt.Println(m, err)
		fmt.Fprintf(res, m, err)
		return
	}
	fmt.Println("Срез структуры перед XML", t)
	res.Header().Set("Content-Type", "application/xml")
	xml.NewEncoder(res).Encode(&t)
}
