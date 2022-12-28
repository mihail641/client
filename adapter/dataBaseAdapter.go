package adapter

import (
	"bytes"
	"encoding/json"
	"example.com/kate/config"
	"fmt"
	"io"
	"net/url"
	"os"
	"strconv"
	//"io"
	"io/ioutil"
	//"log"
	"example.com/projectApiClient"
	"net/http"
)

//постоянные URL
var UrlMain = config.Get().Url_add

// User структура

// Client структура для работы с БД
type DataBaseAdapter struct {
	//композиция типа для вычисления минимального и максимального id
	BaseAdapter
	HTTPClient http.Client
}

//конструктор адаптера для работы с БД
func NewDataBaseAdapter() *DataBaseAdapter {
	return &DataBaseAdapter{
		HTTPClient: http.Client{},
	}
}
func (m *DataBaseAdapter) Close() {
	return
}

// MakeRequestGet метод получения всех значений БД
func (m *DataBaseAdapter) MakeRequestGet() ([]projectApiClient.User, error) {
	URLGET := UrlMain + "users"
	req, err := http.NewRequest("GET", URLGET, nil)
	if err != nil {
		fmt.Println("Проблема с адресом", err)
		return []projectApiClient.User{}, err
	}
	res, err := m.HTTPClient.Do(req)
	if err != nil {
		fmt.Println("проблема подключения к клиенту", err)
		return []projectApiClient.User{}, err

	}

	defer res.Body.Close()
	if 200 != res.StatusCode {
		return nil, fmt.Errorf("%s", res.Body)
	}
	body, err := ioutil.ReadAll(res.Body) // response body is []byte
	if err != nil {
		fmt.Println("Ошибка перевода ответа в строку", err)
		return []projectApiClient.User{}, err

	}
	fmt.Println(string(body))
	p := []projectApiClient.User{}
	fmt.Println("Печать из функции", string(body))
	err = json.Unmarshal(body, &p)
	if err != nil {
		fmt.Println("Can not unmarshal JSON", err)
		return []projectApiClient.User{}, err
	}
	fmt.Println("Структура", p)
	return p, err
}

// MakeRequestCreate метод адаптера создания нового значения
func (m *DataBaseAdapter) MakeRequestCreate(user projectApiClient.User) (projectApiClient.User, error) {
	URL := UrlMain + "user"
	userBytes, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
		return projectApiClient.User{}, err
	}
	byteRead := bytes.NewReader(userBytes)
	req, err := http.NewRequest("POST", URL, byteRead)
	if err != nil {
		fmt.Println("Проблема чтения заголовка", err)
		return projectApiClient.User{}, err
	}
	res, err := m.HTTPClient.Do(req)
	if err != nil {
		fmt.Println("проблема подключения к клиенту", err)
		return projectApiClient.User{}, err
	}
	defer res.Body.Close()
	if 200 != res.StatusCode {
		return projectApiClient.User{}, fmt.Errorf("%s", res.Body)
	}
	body, err := ioutil.ReadAll(res.Body) // response body is []byte
	if err != nil {
		fmt.Println("Ошибка перевода ответа в строку", err)
		return projectApiClient.User{}, err
	} else {
		fmt.Println(string(body))
		return projectApiClient.User{}, err
	}
}

// MakeRequestDelete метод адаптера удаление значений по максимальному id
func (m *DataBaseAdapter) MakeRequestDelete(IdMax int) (projectApiClient.User, error) {
	URL := UrlMain + "user"
	id := strconv.Itoa(IdMax)
	fmt.Println("Максимально id", id)
	id = url.PathEscape(id)
	URLNew := URL + string("/") + id

	req, err := http.NewRequest("DELETE", URLNew, nil)
	if err != nil {
		fmt.Println(err)
		return projectApiClient.User{}, err
	}
	res, err := m.HTTPClient.Do(req)
	if err != nil {
		fmt.Println("Ошибка подключения к клиенту", err)
		return projectApiClient.User{}, err
	}

	defer res.Body.Close()
	if 200 != res.StatusCode {
		return projectApiClient.User{}, fmt.Errorf("%s", res.Body)
	}
	body, err := ioutil.ReadAll(res.Body) // response body is []byte
	if err != nil {
		fmt.Println("Ошибка перевода ответа в строку", err)
		return projectApiClient.User{}, err
	} else {
		fmt.Println(string(body))
		io.Copy(os.Stdout, res.Body)
		return projectApiClient.User{}, err
	}
}

// MakeRequestUpdate метод адаптера изменения значений БД по минимальному id
func (m *DataBaseAdapter) MakeRequestUpdate(user projectApiClient.User) (projectApiClient.User, error) {
	URL := UrlMain + "user"

	userBytes, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
		return projectApiClient.User{}, err
	}
	byteRead := bytes.NewReader(userBytes)
	req, err := http.NewRequest("PUT", URL, byteRead)

	if err != nil {
		fmt.Println(err)
		return projectApiClient.User{}, err
	}
	res, err := m.HTTPClient.Do(req)
	if err != nil {
		fmt.Println("Проблема подключения к клиенту", err)
		return projectApiClient.User{}, err
	}
	defer res.Body.Close()
	if 200 != res.StatusCode {
		return projectApiClient.User{}, fmt.Errorf("%s", res.Body)
	}
	body, err := ioutil.ReadAll(res.Body) // response body is []byte
	if err != nil {
		fmt.Println("Ошибка перевода ответа в строку", err)
		return projectApiClient.User{}, err
	}
	fmt.Println(string(body))
	err = json.Unmarshal(body, &user)
	if err != nil {
		return projectApiClient.User{}, fmt.Errorf("can't parse body as JSON: %w", err)
	}
	return user, err
}

// GetRezultDocumentation метод получающий от API все документы и привязанные к нему Модули и Ошибки
func (m *DataBaseAdapter) GetRezultDocumentation() ([]projectApiClient.Document, error) {
	URLGET := UrlMain + "full"
	req, err := http.NewRequest("GET", URLGET, nil)
	if err != nil {
		fmt.Println("Проблема с адресом", err)
		return []projectApiClient.Document{}, err
	}
	res, err := m.HTTPClient.Do(req)
	if err != nil {
		fmt.Println("проблема подключения к клиенту", err)
		return []projectApiClient.Document{}, err

	}
	defer res.Body.Close()
	if 200 != res.StatusCode {
		return nil, fmt.Errorf("%s", res.Body)
	}
	body, err := ioutil.ReadAll(res.Body) // response body is []byte
	if err != nil {
		fmt.Println("Ошибка перевода ответа в строку", err)
		return []projectApiClient.Document{}, err

	}
	fmt.Println(string(body))
	p := []projectApiClient.Document{}
	fmt.Println("Печать из функции", string(body))
	err = json.Unmarshal(body, &p)
	if err != nil {
		fmt.Println("Can not unmarshal JSON", err)
		return []projectApiClient.Document{}, err
	}

	return p, nil
}

// GetDirectoriesSlice метод модели получающий от API слайс директорий
func (m *DataBaseAdapter) GetDirectories() ([]projectApiClient.Directory, error) {

	URLGET := UrlMain + "directories"
	fmt.Println("URLGET", URLGET)
	req, err := http.NewRequest("GET", URLGET, nil)
	if err != nil {
		fmt.Println("Проблема с адресом", err)
		return nil, err
	}
	res, err := m.HTTPClient.Do(req)
	if err != nil {
		fmt.Println("проблема подключения к клиенту", err)
		return nil, err

	}
	defer res.Body.Close()
	if 200 != res.StatusCode {
		return nil, fmt.Errorf("%s", res.Body)
	}
	body, err := ioutil.ReadAll(res.Body) // response body is []byte
	if err != nil {
		fmt.Println("Ошибка перевода ответа в строку", err)
		return nil, err

	}
	fmt.Println(string(body))
	directories := []projectApiClient.Directory{}
	err = json.Unmarshal(body, &directories)
	if err != nil {
		fmt.Println("Can not unmarshal JSON", err)
		return []projectApiClient.Directory{}, err
	}
	return directories, nil
}
