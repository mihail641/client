package controller

import (
	"example.com/kate/adapterType"
	"fmt"
	"github.com/gorilla/mux"
	"strconv"
	//"log"
	"net/http"
)

// DocumentController структура используется для конструктора контроллер
type DocumentController struct {
}

// NewDocumentController конструктор контроллера, возращающий экземпляр структуры Controller
func NewDocumentController(AdapterType adapterType.AdapterType) *DocumentController {
	return &DocumentController{}
}

// GetSimpleTable метод по выводу в браузере  таблицы с фиксированным количеством столбцов и строк
func (d *DocumentController) GetSimpleTable(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Читаю таблицу")
	tableHTML := `<html lang="ru">
	<table border="1" width="600">
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
	<thead>
	<tr>
    <th>Колонка 1</th>
    <th>Колонка 2</th>
	</tr>
	</thead>
    <tbody>
    <tr>
	<td>Значение 1</td>
	<td>Значение 2</td>
	</tr> <!--ряд с ячейками тела таблицы-->
	<tr>
	<td>Значение 3</td>
	<td>Значение 4</td>
	</tr> <!--ряд с ячейками тела таблицы-->
    </tbody>
	</table>
    </html>`
	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	html := []byte(tableHTML)
	fmt.Println(html)
	res.Write(html)
}

// GetComplexTable метод по выводу в браузер html таблицы со слитыми ячейками и столбцами определенным образом
func (d *DocumentController) GetComplexTable(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Читаю таблицу")
	tableHTML := `<html lang="ru">
	<table border="1" width="600">
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
	<thead>
	<tr>
    <th>Значение 1</th>
    <th>Значение 2</th>
    <th>Значение 3</th>
    <th>Значение 4</th>
	</tr>
	</thead>
    <tbody>
    <tr>
	<td>Текст 1</td>
	<td>Текст 2</td>
    <td>Текст 3</td>
	<td>Текст 4</td>
	</tr> <!--ряд с ячейками тела таблицы-->
	<tr>
	<td>Текст 5</td>
	<td rowspan="2">Текст 6</td>
    <td>Текст 7</td>
	<td>Текст 8</td>
	</tr> <!--ряд с ячейками тела таблицы-->
    <tr>
	<td>Текст 9</td>
	<td>Текст 10</td>
	<td>Текст 11</td>
	</tr> <!--ряд с ячейками тела таблицы-->
    <tr>  
	<td>Текст 12</td>
    <td colspan="3">Текст 13</td>
	</tr> <!--ряд с ячейками тела таблицы-->
    </tbody>
	</table>
    </html>`
	//установливаем заголовок «Content-Type: application», т.к.  мы отправляем html таблицу с запросом через роутер
	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	//преобразование строки в массив байт
	html := []byte(tableHTML)
	fmt.Println(html)
	//вывод в браузере таблицы
	res.Write(html)
}
func (d *DocumentController) GetCertainSizeTable(res http.ResponseWriter, req *http.Request) {
	//получения информации из URL о количестве колонок и строк
	params := mux.Vars(req) // we are extracting 'id' of the Course which we are passing in the url

	var sizeCols = params["sizeCols"]
	//конвертация string в int
	sizeColums, err := strconv.Atoi(sizeCols)
	if err != nil {
		m := "Ошибка перевода количества столбцов из string в int "
		fmt.Println(m,err,)
		fmt.Fprintf(res,m,err,)
		return
	}
	var sizeRows = params["sizeRows"]
	//конвертация string в int
	numRows, err := strconv.Atoi(sizeRows)
	if err != nil {
		m := "Ошибка перевода количества строк из string в int "
		fmt.Println(m,err,)
		fmt.Fprintf(res, m, err,)
		return
	}

	//заголовок таблицы
	tableHead := `<html lang="ru">
	<table border="1" width="600">
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    <thead>
	<tr>`
	var table string
	//цикл формирующий количество наименований колонок в зависимости от данных из URL
	for i := 1; i < sizeColums+1; i++ {
		id := strconv.Itoa(i)
		tableName := `
    <th>Колонка название ` + id + `</th>`
		table = table + tableName
	}

	table = tableHead + table
	//начало тела таблицы
	tableBody :=
		`</tr>
	</thead>
	<tbody>
    <tr>`
	tableH := table + tableBody
	counter := 1
	var tableBodyMain string
	var tableMain string
	var tableBodyFinishRow string
	var tableMn string
	//циклы формирующие количество столбцов и строк в зависимости от информации в URL
	for j := 0; j < numRows; j++ {
		tableMain = ""
		for i := 0; i < sizeColums; i++ {
			idText := strconv.Itoa(counter)
			counter = counter + 1
			tableBodyMain = `<td>Значение ` + idText + `</td>`
			tableMain = tableMain + tableBodyMain
		}
		tableBodyFinishRow = ` <tr> </tr> <!--ряд с ячейками тела таблицы-->`
		tableMn = tableMn + tableMain + tableBodyFinishRow
	}
	tableMain = tableH + tableMn
	//окончание таблицы
	tableEnd := `</tbody>
	</table>
	</html>`
	table = tableMain + tableEnd
	//перевод строки в байты
	html := []byte(table)
	//отправка в браузер
	res.Write(html)
}

