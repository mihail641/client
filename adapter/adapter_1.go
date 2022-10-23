package adapter

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type FileAdapter struct {
	File os.File
}

//
//type User struct {
//	ID   int    `xml:"id",json:"id"`
//	Name string `xml:"name",json:"name"`
//	Sale int    `xml:"sale",json:"sale"`
//}

func NewFileAdapter() *FileAdapter {
	return &FileAdapter{}
}
func file() (File *os.File) {
	File, err := os.Open("file_storage")
	if err != nil {
		m := "Ошибка выполнеия открытия файла: %s"
		fmt.Println(m, err)
	}
	defer File.Close()

	return File
}
func (l *FileAdapter) MakeRequestGet() ([]User, error) {
	File := file()
	//File, err := os.Open("file_storage")
	//if err != nil {
	//	m := "Ошибка выполнеия открытия файла: %s"
	//	fmt.Println(m, err)
	//}
	//defer File.Close()
	reader := bufio.NewReader(File)
	slice := make([]string, 0)
	for {
		str, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		fmt.Println(str)
		for i := 0; i < len(slice); i++ {
			str = strings.TrimRight(str, "\n")
		}
		slice = append(slice, str)
	}
	k := []User{}
	word := make([]string, 0)
	for i := 1; i < len(slice); i++ {
		fmt.Println(slice[i])
		word = strings.Split(slice[i], "\t")
		m, err := strconv.Atoi(word[0])
		if err != nil {
			m := "Ошибка перевода из string в int: %s"
			fmt.Println(m, err)
			return []User{}, err
		}
		fmt.Println(m)
		p := word[1]
		fmt.Println(p)
		z, err := strconv.Atoi(word[2])
		if err != nil {
			m := "Ошибка перевода из string в int: %s"
			fmt.Println(m, err)
			return []User{}, err
		}
		fmt.Println(z)
		k = append(k, User{m, p, z})
	}
	return k, nil

}
func (f *FileAdapter) MakeRequestCreate(user User) (User, error) {
	File := file()
	//var (
	//	File *os.File
	//)
	//File, err := os.Open("file_storage")
	//if err != nil {
	//	m := "Ошибка выполнеия открытия файла: %s"
	//	fmt.Println(m, err)
	//}
	//defer File.Close()
	reader := bufio.NewReader(File)
	slice := make([]string, 0)
	//for {
	str, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Строка не читается", err)

	}
	str = strings.TrimRight(str, "\n")
	fmt.Println(str)
	slice = append(slice, str)
	fmt.Println(slice)

	k := []User{}
	for i := 0; i < 1; i++ {
		word := strings.Split(slice[i], " ")
		m, err := strconv.Atoi(word[2])
		if err != nil {
			m := "Ошибка перевода из string в int: %s"
			fmt.Println(m, err)
			return User{}, err
		}
		fmt.Println(m)
		k = append(k, User{m, user.Name, user.Sale})
		//
		//fmt.Println(k)

		values := []string{}
		values = append(values, strconv.Itoa(m))
		values = append(values, user.Name)
		values = append(values, strconv.Itoa(user.Sale))
		values = append(values, "\n")
		msg := strings.Join(values, "\t")
		fmt.Println("Получившаяся строка", msg)

		f, err := os.OpenFile("file_storage", os.O_APPEND|os.O_WRONLY, 0600)
		if err != nil {
			fmt.Println("Ошибка открытия файла для записи", err)
		}
		defer f.Close()

		if _, err = f.WriteString(msg); err != nil {
			if err != nil {
				fmt.Println("Ошибка записи", err)
			}
		}

	}

	return User{}, err
}
func (f *FileAdapter) MakeRequestDelete(IdMax int) (User, error) {
	File := file()
	//var (
	//	File *os.File
	//)
	//File, err := os.Open("file_storage")
	//if err != nil {
	//	m := "Ошибка выполнеия открытия файла: %s"
	//	fmt.Println(m, err)
	//}
	//defer File.Close()
	reader := bufio.NewReader(File)
	slice := make([]string, 0)
	for {
		str, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		IdDelete := strconv.Itoa(IdMax)
		r, err := regexp.MatchString(IdDelete, str)
		if err != nil {
			fmt.Println("Нахождение строки", err)
		}
		fmt.Println(r)
		if r != true {
			for i := 0; i < len(slice); i++ {
				str = strings.TrimRight(str, "\n")
			}
			slice = append(slice, str)

		}

	}
	msg := strings.Join(slice, "\n")
	fmt.Println("Получившаяся строка", msg)

	err := os.WriteFile("file_storage", []byte(msg), 0666)
	if err != nil {
		log.Fatal(err)
	}

	return User{}, err
}

func (f *FileAdapter) MakeRequestUpdate(user User) (User, error) {
	File := file()
	//var (
	//	File *os.File
	//)
	//File, err := os.Open("file_storage")
	//if err != nil {
	//	m := "Ошибка выполнеия открытия файла: %s"
	//	fmt.Println(m, err)
	//}
	//defer File.Close()
	reader := bufio.NewReader(File)
	slice := make([]string, 0)
	IdUpdate := strconv.Itoa(user.ID)
	stringUpdate := make([]string, 0)
	stringUpdate = append(stringUpdate, IdUpdate)
	stringUpdate = append(stringUpdate, user.Name)
	stringUpdate = append(stringUpdate, strconv.Itoa(user.Sale))
	stringUp := strings.Join(stringUpdate, "\t")
	fmt.Println("Строка в добавление", stringUp)
	for {
		str, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		r, err := regexp.MatchString(IdUpdate, str)
		if err != nil {
			fmt.Println("Нахождение строки", err)
		}
		if r != true {
			for i := 0; i < len(slice); i++ {
				str = strings.TrimRight(str, "\n")
			}
			slice = append(slice, str)
		}
	}
	slice = append(slice, stringUp)
	msg := strings.Join(slice, "\n")
	fmt.Println(msg)
	err := os.WriteFile("file_storage", []byte(msg), 0666)
	if err != nil {
		log.Fatal(err)
	}
	return User{}, err
}
