package adapter

import "fmt"

type BaseAdapter struct {
}
// Min метод вычисление минимального id
func (m *BaseAdapter) Min(p []User) int {

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
// Max метод адаптера по определению  максимального значения id
func (m *BaseAdapter) Max(p []User) int {
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
