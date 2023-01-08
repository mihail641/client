package adapter

import "example.com/projectApiClient"

// IAdapter интерфейс объединяющие методы двух объектов
type IAdapter interface {
	MakeRequestGet() ([]projectApiClient.User, error)
	MakeRequestCreate(user projectApiClient.User) (projectApiClient.User, error)
	MakeRequestUpdate(user projectApiClient.User) (projectApiClient.User, error)
	MakeRequestDelete(idMax int) (projectApiClient.User, error)
	GetRezultDocumentation() ([]projectApiClient.Document, error)
	GetDirectories() ([]projectApiClient.Directory, error)
	Max(p []projectApiClient.User) int
	Min(p []projectApiClient.User) int
	Close()
}
