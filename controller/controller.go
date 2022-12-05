package controller

import (
	"encoding/xml"
	"example.com/kate/adapterType"
	"example.com/kate/model"
	"fmt"
	//"log"
	"net/http"
)

// Controller структура используется для конструктора контроллер
type Controller struct {
	controller *model.Model
}

// NewController конструктор контроллера, возращающий экземпляр структуры Controller
func NewController(AdapterType adapterType.AdapterType) *Controller {
	return &Controller{
		controller: model.NewModel(AdapterType),
	}
}

// HandleHttp метод контроллера для запуска алгоритма модели, и возврата в роутер данных в формате xml
func (m *Controller) HandleHttp(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Сервер запустился")
	//ClientAlgorithmTake метод модели
	var t, err = m.controller.ClientAlgorithmTake()
	if err != nil {
		m := "Ошибка выполнеия контроллера: %s"
		fmt.Println(m, err)
		fmt.Fprintf(res, m, err)
		return
	}
	fmt.Println("Срез структуры перед XML", t)
	//	установливаем заголовок «Content-Type: application/xml», т.к. потому что мы отправляем данные XML с запросом через роутер
	res.Header().Set("Content-Type", "application/xml")
	if err != nil {
		res.Header().Set("Content-Type", "text/html")
	}
	//кодирование в xml результата выполнения метода и передача в пакет main
	xml.NewEncoder(res).Encode(&t)
}
