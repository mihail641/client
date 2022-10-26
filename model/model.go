package model

import (
	"example.com/kate/adapter"
	"fmt"
)

//Model структура используется для конструктора контроллер
type Model struct {
	adapter adapter.IAdapter
}

// NewModel конструктор модели
func NewModel(concreteAdapterType adapter.AdapterType) *Model {
	var m adapter.IAdapter

	switch concreteAdapterType {
	case adapter.DB:
		m = adapter.NewDataBaseAdapter()
	case adapter.File:
		m = adapter.NewFileAdapter()
	}
	return &Model{adapter: m}

}

// ClientAlgorithmTake метод модели
func (d *Model) ClientAlgorithmTake() ([]adapter.User, error) {

	//MakeRequestGet получение из адаптера всех пользователей БД
	user, err := d.adapter.MakeRequestGet()
	if err != nil {
		m := "Ошибка выполнеия 1 функции получения информации о всех пользователях"
		fmt.Println(m, err)
		return []adapter.User{}, err
	}
	fmt.Println("Переданная структура", user)
	//отправление структуры БД в метод адаптера Min, для получения минимального id
	IdMin := d.adapter.Min(user)
	//отправление структуры БД в метод адаптера Max, для получения максимального id
	IdMax := d.adapter.Max(user)
	fmt.Println("Максимальное значение", IdMax)
	fmt.Println("Минимальное значение", IdMin)
	//Обращение к методу адаптера к изменению самого минимального по id значения БД
	var user1 adapter.User
	user1 = adapter.User{
		ID:   IdMin,
		Name: "Vova",
		Sale: 654,
	}
	fmt.Println("user.ID", IdMin)

	_, err = d.adapter.MakeRequestUpdate(user1)
	if err != nil {
		m := "Ошибка выполнеия  функции изменения пользователя"
		fmt.Println(m, err)
		return []adapter.User{}, err
	}
	//обращение к модели адаптера к удалению максимального по id значения БД
	_, err = d.adapter.MakeRequestDelete(IdMax)
	if err != nil {
		m := "Ошибка выполнеия  функции удаления пользователя"
		fmt.Println(m, err)
		return []adapter.User{}, err
	}
	var user3 adapter.User
	user3 = adapter.User{
		Name: "RED",
		Sale: 895,
	}
	//обращение к модели адаптера к созданию нового значения БД
	_, err = d.adapter.MakeRequestCreate(user3)
	if err != nil {
		m := "Ошибка выполнеия  функции создания пользователя"
		fmt.Println(m, err)
		return []adapter.User{}, err
	}
	//обращение к модели адаптера к получению новых значений БД
	users, err := d.adapter.MakeRequestGet()
	if err != nil {
		m := "Ошибка выполнеия 2 функции получения информации о всех пользователях: %s"
		fmt.Println(m, err)
		return []adapter.User{}, err
	}
	fmt.Println(users)
	return users, nil
}
