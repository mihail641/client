package adapter

// User структура пользователей с продажами
type User struct {
	ID   int    `xml:"id",json:"id"`
	Name string `xml:"name",json:"name"`
	Sale int    `xml:"sale",json:"sale"`
}
type Document struct {
	Id      int64    `json:"id" :"id"`
	Title   string   `json:"title" :"title"`
	Modules []Module `json:"modules"`
}
type Module struct {
	Id     int64   `json:"id"`
	Title  string  `json:"title"`
	Errors []Error `json:"errors"`
}
type Error struct {
	Id    int64  `json:"id"`
	Title string `json:"title"`
}
