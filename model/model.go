package model

import (
	"example.com/kate/adapter"
	"fmt"
)

//Model структура используется для конструктора контроллер
type Model struct {
	model *adapter.Client
}

// NewModel конструктор модели
func NewModel() *Model {
	return &Model{
		model: adapter.NewClient(),
	}
}

// ClientAlgorithmTake метод модели
func (d *Model) ClientAlgorithmTake() ([]adapter.User, error) {
	//MakeRequestGet получение из адаптера всех пользователей БД
	user, err := d.model.MakeRequestGet()
	if err != nil {
		m := "Ошибка выполнеия 1 функции получения информации о всех пользователях"
		fmt.Println(m, err)
		return []adapter.User{}, err
	}
	fmt.Println("Переданная структура", user)
	//отправление структуры БД в метод адаптера Min, для получения минимального id
	IdMin := d.model.Min(user)
	//отправление структуры БД в метод адаптера Max, для получения максимального id
	IdMax := d.model.Max(user)
	fmt.Println("Максимальное значение", IdMax)
	fmt.Println("Минимальное значение", IdMin)
	//Обращение к методу адаптера к изменению самого минимального по id значения БД
	_, err = d.model.MakeRequestUpdate(IdMin)
	if err != nil {
		m := "Ошибка выполнеия  функции изменения пользователя"
		fmt.Println(m, err)
		return []adapter.User{}, err
	}
	//обращение к модели адаптера к удалению максимального по id значения БД
	_, err = d.model.MakeRequestDelete(IdMax)
	if err != nil {
		m := "Ошибка выполнеия  функции удаления пользователя"
		fmt.Println(m, err)
		return []adapter.User{}, err
	}
	//обращение к модели адаптера к созданию нового значения БД
	_, err = d.model.MakeRequestCreate()
	if err != nil {
		m := "Ошибка выполнеия  функции создания пользователя"
		fmt.Println(m, err)
		return []adapter.User{}, err
	}
	//обращение к модели адаптера к получению новых значений БД
	t, err := d.model.MakeRequestGet()
	if err != nil {
		m := "Ошибка выполнеия 2 функции получения информации о всех пользователях: %s"
		fmt.Println(m, err)
		return []adapter.User{}, err
	}
	fmt.Println(t)
	return t, nil
}
