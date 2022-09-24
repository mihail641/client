package adapter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/url"
	"os"
	"strconv"
	//"io"
	"io/ioutil"
	//"log"
	"net/http"
)

const (
	URL    = "http://127.0.0.1:4000/user"
	URLGET = "http://127.0.0.1:4000/users"
)

type User struct {
	ID   int    `xml:"id_xml",json:"id"`
	Name string `xml:"name_xml",json:"name"`
	Sale int    `xml:"sale_xml",json:"sale_xml"`
}

type Client struct {
	HTTPClient http.Client
}

func NewClient() *Client {
	return &Client{
		HTTPClient: http.Client{},
	}
}
func (m *Client) MakeRequestGet() ([]User, error) {
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
func (m *Client) MakeRequestCreate() (User, error) {

	var user User
	user = User{
		Name: "RED",
		Sale: 895,
	}
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
	}
	fmt.Println(string(body))
	return User{}, err
}
func (m *Client) MakeRequestDelete(IdMax int) (User, error) {
	id := strconv.Itoa(IdMax)
	fmt.Println("Максимально id", id)
	id = url.PathEscape(id)
	u, err := url.Parse(URL)
	if err != nil {
		log.Fatal(err)
		return User{}, err
	}
	rel, err := u.Parse(id)
	if err != nil {
		log.Fatal(err)
		return User{}, err
	}
	fmt.Println(rel)
	udlStr := rel.String()
	req, err := http.NewRequest("DELETE", udlStr, nil)
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
	}
	fmt.Println(string(body))
	io.Copy(os.Stdout, res.Body)
	return User{}, err
}
func (m *Client) MakeRequestUpdate(IdMin int) (User, error) {
	var user User
	user = User{
		ID:   IdMin,
		Name: "Vova",
		Sale: 654,
	}
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
func (m *Client) Min(p []User) int {

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

func (m *Client) Max(p []User) int {
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
