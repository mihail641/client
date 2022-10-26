package adapter

type User struct {
	ID   int    `xml:"id",json:"id"`
	Name string `xml:"name",json:"name"`
	Sale int    `xml:"sale",json:"sale"`
}

type IAdapter interface {
	MakeRequestGet() ([]User, error)
	MakeRequestCreate(user User) (User, error)
	MakeRequestUpdate(user User) (User, error)
	MakeRequestDelete(idMax int) (User, error)
	Max(p []User) int
	Min(p []User) int
}
