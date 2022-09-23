package controller

import (
	"encoding/xml"
	"example.com/kate/model"
	"fmt"
	//"log"
	"net/http"
)

type User struct {
	users *model.User
}

type Controller struct {
	controller *model.Adapter
}

func NewController() *Controller {
	return &Controller{
		controller:model.NewModel(),
	}
}

func (m*Controller) HandleHttp(res http.ResponseWriter, req *http.Request)  {
	//client := http.Client{}
	//if err != nil {
	//	fmt.Println(err)
	//
	//}
	//defer res.Body.Close()
	//body, err := ioutil.ReadAll(res.Body) // response body is []byte
	//if err !=nil {
	//	fmt.Println(err)
	//
	//}
	//fmt.Println(string(body))
	//

	var t, err = m.controller.ClientAlgorithmTake()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Срез структуры перед XML", t)
	res.Header().Set("Content-Type", "application/xml")
	xml.NewEncoder(res).Encode(&User{})
}

