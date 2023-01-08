package adapter

import (
	"bufio"
	"example.com/projectApiClient"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

//Структура адаптера при работе с файлом
type FileAdapter struct {
	//композиция типа для вычисления минимального и максимального id
	BaseAdapter
	file   *os.File
	reader *bufio.Reader
}

//Конструктор адаптера при работе с файлом
func NewFileAdapter() (*FileAdapter, error) {
	file, err := os.Open("file_storage")
	if err != nil {
		m := "Ошибка выполнеия открытия файла: %s"
		fmt.Println(m, err)
		return nil, err
	}
	//defer file.Close()
	reader := bufio.NewReader(file)

	return &FileAdapter{
		file:   file,
		reader: reader,
	}, err
}
func (f *FileAdapter) Close() {
	f.file.Close()
	return
}

// MakeRequestGet метод получения всех значений из файла
func (f *FileAdapter) MakeRequestGet() ([]projectApiClient.User, error) {
	//формируется слайс из строк файла, каждая строка записывается как отдельный элемент
	slice := make([]string, 0)
	for {
		str, err := f.reader.ReadString('\n')
		if err != nil {
			break
		}
		// проверка на пустые строки
		if len(strings.TrimSpace(str)) == 0 {
			fmt.Println("Строка пустая")
		}

		fmt.Println(str)
		//TrimRight возвращает срез строки str убирая перенос строки.
		for i := 0; i < len(slice); i++ {
			str = strings.TrimRight(str, "\n")
		}
		slice = append(slice, str)
	}
	k := []projectApiClient.User{}
	//создание слайса из слов
	word := make([]string, 0)
	for i := 1; i < len(slice); i++ {
		fmt.Println(slice[i])
		word = strings.Split(slice[i], "\t")
		//перевод id из string в int
		m, err := strconv.Atoi(word[0])
		if err != nil {
			m := "Ошибка перевода из string в int: %s"
			fmt.Println(m, err)
			return []projectApiClient.User{}, err
		}
		fmt.Println(m)
		p := word[1]
		fmt.Println(p)
		//перевод id из 	//создание слайса из слов
		z, err := strconv.Atoi(word[2])
		if err != nil {
			m := "Ошибка перевода из string в int: %s"
			fmt.Println(m, err)
			return []projectApiClient.User{}, err
		}
		fmt.Println(z)
		//присваивание в структуру User значений из слайса
		k = append(k, projectApiClient.User{m, p, z})
	}
	return k, nil

}

// MakeRequestCreate метод адаптера создания нового значения в файл
func (f *FileAdapter) MakeRequestCreate(user projectApiClient.User) (projectApiClient.User, error) {
	//формируется слайс из 1 строки файла,
	slice := make([]string, 0)
	//считывает первую строку из файла
	str, err := f.reader.ReadString('\n')
	if err != nil {
		fmt.Println("Строка не читается", err)
		return projectApiClient.User{}, err
	}
	//TrimRight возвращает срез 1-ой строки str убирая перенос строки.
	str = strings.TrimRight(str, "\n")
	fmt.Println(str)
	slice = append(slice, str)
	fmt.Println(slice)

	k := []projectApiClient.User{}
	//создание слайса из слов 1 строки
	for i := 0; i < 1; i++ {
		word := strings.Split(slice[i], " ")
		//перевод 3 слова в слайсе из string в int
		m, err := strconv.Atoi(word[2])
		if err != nil {
			m := "Ошибка перевода из string в int: %s"
			fmt.Println(m, err)
			return projectApiClient.User{}, err
		}
		fmt.Println(m)
		//присваивание структуре User id, получившегося при чтении файла, а так же значений Name, Sale из Модели
		k = append(k, projectApiClient.User{m, user.Name, user.Sale})
		//перевод Id, Sale из int в string, добавление переноса строки и табуляции между словами
		values := []string{}
		values = append(values, strconv.Itoa(m))
		values = append(values, user.Name)
		values = append(values, strconv.Itoa(user.Sale))
		values = append(values, "\n")
		msg := strings.Join(values, "\t")
		fmt.Println("Получившаяся строка", msg)
		//запись новой строки в файл
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
	return projectApiClient.User{}, err
}

// MakeRequestDelete метод адаптера удаление значений по максимальному id
func (f *FileAdapter) MakeRequestDelete(IdMax int) (projectApiClient.User, error) {
	//формируется слайс из строк файла, каждая строка записывается как отдельный элемент
	slice := make([]string, 0)
	for {
		str, err := f.reader.ReadString('\n')
		if err != nil {
			break
		}
		//перевод IdMax в string
		IdDelete := strconv.Itoa(IdMax)
		//нахождение строки содержащей удаляемый id
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
	//запись всех строк без удаленной строки в файл
	err := os.WriteFile("file_storage", []byte(msg), 0666)
	if err != nil {
		log.Fatal(err)
	}

	return projectApiClient.User{}, err
}

// MakeRequestUpdate метод адаптера изменения значений БД по минимальному id
func (f *FileAdapter) MakeRequestUpdate(user projectApiClient.User) (projectApiClient.User, error) {
	//формируется слайс из 1 строки файла,
	slice := make([]string, 0)
	//создается строка из данных структуры User переданной из модели, разделяется табуляцией и
	//знаком новой строки в конце, объединяется в 1 строку
	IdUpdate := strconv.Itoa(user.ID)
	stringUpdate := make([]string, 0)
	stringUpdate = append(stringUpdate, IdUpdate)
	stringUpdate = append(stringUpdate, user.Name)
	stringUpdate = append(stringUpdate, strconv.Itoa(user.Sale))
	stringUp := strings.Join(stringUpdate, "\t")
	fmt.Println("Строка в добавление", stringUp)
	for {
		str, err := f.reader.ReadString('\n')
		if err != nil {
			break
		}
		//нахождение строки содержащей изменяемый id
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
	//запись всех строк в файл
	err := os.WriteFile("file_storage", []byte(msg), 0666)
	if err != nil {
		log.Fatal(err)
	}
	return projectApiClient.User{}, err
}
func (f *FileAdapter) GetRezultDocumentation() ([]projectApiClient.Document, error) {
	return nil, nil
}
func (f *FileAdapter) GetDirectories() ([]projectApiClient.Directory, error) {
	return nil, nil
}
