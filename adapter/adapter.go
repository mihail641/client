package adapter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"os"
	"strconv"
	//"io"
	"io/ioutil"
	//"log"
	"net/http"
)

//постоянные URL
const (
	URL    = "http://127.0.0.1:4000/user"
	URLGET = "http://127.0.0.1:4000/users"
)

// User структура
type User struct {
	ID   int    `xml:"id",json:"id"`
	Name string `xml:"name",json:"name"`
	Sale int    `xml:"sale",json:"sale"`
}

// Client структура
type Adapter struct {
	HTTPClient http.Client
}

//конструктор адаптера
func NewAdapter() *Adapter {
	return &Adapter{
		HTTPClient: http.Client{},
	}
}

// MakeRequestGet метод получения всех значений БД
func (m *Adapter) MakeRequestGet() ([]User, error) {
	req, err := http.NewRequest("GET", URLGET, nil)
	if err != nil {
		fmt.Println("Проблема с адресом", err)
		return []User{}, err
	}
	res, err := m.HTTPClient.Do(req)
	if err != nil {
		fmt.Println("проблема подключения к клиенту", err)
		return []User{}, err

	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body) // response body is []byte
	if err != nil {
		fmt.Println("Ошибка перевода ответа в строку", err)
		return []User{}, err

	}
	fmt.Println(string(body))
	p := []User{}
	fmt.Println("Печать из функции", string(body))
	err = json.Unmarshal(body, &p)
	if err != nil {
		fmt.Println("Can not unmarshal JSON", err)
		return []User{}, err
	}
	fmt.Println("Структура", p)
	return p, err
}

// MakeRequestCreate метод адаптера создания нового значения
func (m *Adapter) MakeRequestCreate(user User) (User, error) {

	//var user User
	//user = User{
	//	Name: "RED",
	//	Sale: 895,
	//}
	userBytes, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
		return User{}, err
	}
	byteRead := bytes.NewReader(userBytes)
	req, err := http.NewRequest("POST", URL, byteRead)
	if err != nil {
		fmt.Println("Проблема чтения заголовка", err)
		return User{}, err
	}
	res, err := m.HTTPClient.Do(req)
	if err != nil {
		fmt.Println("проблема подключения к клиенту", err)
		return User{}, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body) // response body is []byte
	if err != nil {
		fmt.Println("Ошибка перевода ответа в строку", err)
		return User{}, err
	} else {
		fmt.Println(string(body))
		return User{}, err
	}
}

// MakeRequestDelete метод адаптера удаление значений по максимальному id
func (m *Adapter) MakeRequestDelete(IdMax int) (User, error) {
	id := strconv.Itoa(IdMax)
	fmt.Println("Максимально id", id)
	id = url.PathEscape(id)
	URLNew := URL + string("/") + id

	req, err := http.NewRequest("DELETE", URLNew, nil)
	if err != nil {
		fmt.Println(err)
		return User{}, err
	}
	res, err := m.HTTPClient.Do(req)
	if err != nil {
		fmt.Println("Ошибка подключения к клиенту", err)
		return User{}, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body) // response body is []byte
	if err != nil {
		fmt.Println("Ошибка перевода ответа в строку", err)
		return User{}, err
	} else {
		fmt.Println(string(body))
		io.Copy(os.Stdout, res.Body)
		return User{}, err
	}
}

// MakeRequestUpdate метод адаптера изменения значений БД по минимальному id
func (m *Adapter) MakeRequestUpdate(user User) (User, error) {
	//var user User
	//user.ID = IdMin
	//user = User{
	//	ID:   user.ID,
	//	Name: "Vova",
	//	Sale: 654,
	//}
	//fmt.Println("user.ID", IdMin)

	userBytes, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
		return User{}, err
	}
	byteRead := bytes.NewReader(userBytes)
	req, err := http.NewRequest("PUT", URL, byteRead)

	if err != nil {
		fmt.Println(err)
		return User{}, err
	}
	res, err := m.HTTPClient.Do(req)
	if err != nil {
		fmt.Println("Проблема подключения к клиенту", err)
		return User{}, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body) // response body is []byte
	if err != nil {
		fmt.Println("Ошибка перевода ответа в строку", err)
		return User{}, err
	}
	fmt.Println(string(body))
	err = json.Unmarshal(body, &user)
	if err != nil {
		return User{}, fmt.Errorf("can't parse body as JSON: %w", err)
	}
	return user, err
}

// Min метод адаптера нахождения минимального id в структуре
func (m *Adapter) Min(p []User) int {

	var k []int

	for _, rec := range p {
		k = append(k, rec.ID)
	}
	IdMin := k[0]
	for _, value := range k {
		if value < IdMin {
			IdMin = value
		}
	}
	return IdMin
}

// Max метод адаптера по определению в БД максимального значения id
func (m *Adapter) Max(p []User) int {
	var k []int
	for _, rec := range p {
		k = append(
			k,
			rec.ID,
		)
	}
	IdMax := k[0]
	for _, value := range k {
		if value > IdMax {
			IdMax = value
		}
	}
	fmt.Println(IdMax)
	return IdMax
}
