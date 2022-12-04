package adapter

// User структура пользователей с продажами
type User struct {
	ID   int    `xml:"id",json:"id"`
	Name string `xml:"name",json:"name"`
	Sale int    `xml:"sale",json:"sale"`
}

