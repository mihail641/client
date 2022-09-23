package model

import (
	"example.com/kate/adapter"
	"fmt"
)

type Adapter struct {
	model*adapter.Client

}
type User struct {
	users *adapter.User
}

func NewModel() *Adapter {
	return &Adapter{
		model: adapter.NewClient(),
	}
	
}


func (d*Adapter)ClientAlgorithmTake()([]User,error) {

user, err := d.model.MakeRequestGet()
if err != nil {
fmt.Println(err)
}
fmt.Println("Переданная структура", user)
IdMin := d.model.Min(user)
if err != nil {
fmt.Println(err)
}
IdMax := d.model.Max(user)
fmt.Println("Максимальное значение", IdMax)
fmt.Println("Минимальное значение", IdMin)

d.model.MakeRequestUpdate(IdMin)
d.model.MakeRequestDelete(IdMax)
d.model.MakeRequestCreate()
	if err != nil {
		fmt.Println(err)
	}
t, err:=d.model.MakeRequestGet()
if err != nil {
fmt.Println(err)
}
	fmt.Println(t)

return []User{}, err
}
