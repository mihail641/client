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
	"strings"

	//"io"
	"io/ioutil"
	//"log"
	"net/http"
)

const (
	URL = "http://127.0.0.1:4000/user"
	URLGET="http://127.0.0.1:4000/users"


)
type User struct {
	ID   int    `xml:"id_xml",json:"id"`
	Name string `xml:"name_xml",json:"name"`
	Sale int    `xml:"sale_xml",json:"sale_xml"`
}

type Client struct {
	//Get(url string) (*http.Response, error)
	//Do(req *http.Request) (*http.Response, error)
	//Provider        Provider
	HTTPClient                 http.Client

}
//type Provider interface {
//	Get(url string) (*http.Response, error)
//	Do(req *http.Request) (*http.Response, error)
//}
func NewClient() *Client {
	//httpClient := &http.Client{}
	//client := &Client{
	//	Provider:        httpClient}
	return &Client{
		HTTPClient: http.Client{},
	}
}
func (m*Client) ClientServ (c*User) (bool, error)  {
	url := &url.URL{
		Scheme: "http",
		Path:   "menu.html",
	}
	res, err := m.Provider.Get(url.String())
	if err != nil {
		return false, err
	}
	body := res.Body
	defer body.Close()
	buf, err := ioutil.ReadAll(body)
	if err != nil {
		return false, err
	}
	return strings.Contains(string(buf), "<title>WebRelay</title>"), nil
}
func (m*Client)MakeRequestGet()([]User, error)  {
	var req, err = http.NewRequest("GET", URLGET, nil)
	if err != nil {
		fmt.Println("Проблема с адресом", err)
	}

		body, err := ioutil.ReadAll(resp.Body) // response body is []byte
	if err != nil {
		fmt.Println(err)
	}
	p:=[]User {}

	fmt.Println("Печать из функции", string(body))
	err = json.Unmarshal(body, &p)
	if err != nil {
		fmt.Println("Can not unmarshal JSON", err)
	}
	fmt.Println("Структура", p)
	return p , err
}
func (m*Client)MakeRequestCreate()(User, error)  {

	var user User
	user=User{
		Name: "RED",
		Sale: 895,
	}
	userBytes, err:= json.Marshal(user)
	if err != nil {
		fmt.Println(err)

	}
	byteRead:=bytes.NewReader(userBytes)
	req, err:= http.NewRequest("POST", URL, byteRead)
	if err != nil {
		fmt.Println(err)
	}
	//resp, err := m.Client.Do(req)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body) // response body is []byte
	if err !=nil {
		fmt.Println(err)
	}
	fmt.Println(string(body))
	return User{}, err
}
func (m*Client) MakeRequestDelete (IdMax int) (User, error){

	id:= strconv.Itoa(IdMax)
	fmt.Println("Максимально id",id)
	id=url.PathEscape(id)


	u, err := url.Parse(URL)
	if err != nil {
		log.Fatal(err)
	}

	rel, err := u.Parse(id)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(rel)

	udlStr:=rel.String()

	req, err:= http.NewRequest("DELETE", udlStr, nil)

	if err != nil {
		fmt.Println(err)

	}
	//resp, err := m.Client.Do(req)
	//if err != nil {
	//	fmt.Println(err)
	//
	//}
	//defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body) // response body is []byte
	if err !=nil {
		fmt.Println(err)

	}
	fmt.Println(string(body))


	//fmt.Println(PrettyPrint(result))

	// Loop through the data node for the FirstName
	io.Copy(os.Stdout, resp.Body)
	return User{}, err
}
func (m*Client) MakeRequestUpdate (IdMin int) (User, error) {
	var user User
	user=User{
		ID: IdMin,
		Name: "Vova",
		Sale: 654,
	}
	userBytes, err:= json.Marshal(user)
	if err != nil {
		fmt.Println(err)
	}
	byteRead:=bytes.NewReader(userBytes)
	req, err:= http.NewRequest("PUT", URL, byteRead)

	if err != nil {
		fmt.Println(err)
	}
	//resp, err := m.Client.Do(req)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//defer resp.Body.Close()
	//body, err := ioutil.ReadAll(resp.Body) // response body is []byte
	//if err !=nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(string(body))
	err = json.Unmarshal(body, &user)
	if err != nil {
		return User{}, fmt.Errorf("can't parse body as JSON: %w", err)
	}
	return user, err
}
func (m*Client)Min (p []User) int {

	var k []int

	for _, rec := range p {
		k= append(k, rec.ID)
	}
	IdMin:=k[0]
	for _, value := range k {
		if value < IdMin {
			IdMin = value
		}
	}
	return IdMin
}

func (m*Client) Max (p []User)  int {




	var k []int

	for _, rec := range p {
		k= append(
			k,
			rec.ID,
		)
	}

	IdMax:=k[0]

	for _, value := range k {
		if value > IdMax {
			IdMax = value
		}
	}
	fmt.Println(IdMax)
	return IdMax
}
