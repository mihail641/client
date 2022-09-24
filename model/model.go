package model

import (
	"example.com/kate/adapter"
	"fmt"
)

type Model struct {
	model *adapter.Client
}

func NewModel() *Model {
	return &Model{
		model: adapter.NewClient(),
	}
}
func (d *Model) ClientAlgorithmTake() ([]adapter.User, error) {
	user, err := d.model.MakeRequestGet()
	if err != nil {
		m := "Ошибка выполнеия 1 функции получения информации о всех пользователях"
		fmt.Println(m, err)
		return []adapter.User{}, err
	}
	fmt.Println("Переданная структура", user)
	IdMin := d.model.Min(user)
	IdMax := d.model.Max(user)
	fmt.Println("Максимальное значение", IdMax)
	fmt.Println("Минимальное значение", IdMin)
	_, err = d.model.MakeRequestUpdate(IdMin)
	if err != nil {
		m := "Ошибка выполнеия  функции изменения пользователя"
		fmt.Println(m, err)
		return []adapter.User{}, err
	}
	_, err = d.model.MakeRequestDelete(IdMax)
	if err != nil {
		m := "Ошибка выполнеия  функции удаления пользователя"
		fmt.Println(m, err)
		return []adapter.User{}, err
	}
	_, err = d.model.MakeRequestCreate()
	if err != nil {
		m := "Ошибка выполнеия  функции создания пользователя"
		fmt.Println(m, err)
		return []adapter.User{}, err
	}
	t, err := d.model.MakeRequestGet()
	if err != nil {
		m := "Ошибка выполнеия 2 функции получения информации о всех пользователях: %s"
		fmt.Println(m, err)
		return []adapter.User{}, err
	}
	fmt.Println(t)
	return []adapter.User{}, nil
}
