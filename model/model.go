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
	m := "Ошибка выполнеия 1 функции получения информации о всех пользователях"
	fmt.Println(m, err)
	return []User{},err
}
fmt.Println("Переданная структура", user)
IdMin := d.model.Min(user)
IdMax := d.model.Max(user)
fmt.Println("Максимальное значение", IdMax)
fmt.Println("Минимальное значение", IdMin)

d.model.MakeRequestUpdate(IdMin)
d.model.MakeRequestDelete(IdMax)
d.model.MakeRequestCreate()
t, err:=d.model.MakeRequestGet()
	if err != nil {
		m := "Ошибка выполнеия 2 функции получения информации о всех пользователях: %s"
		fmt.Println(m, err)
		return []User{},err
	}
	fmt.Println(t)

return []User{}, nil
}
