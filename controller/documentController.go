package controller

import (
	"fmt"
	//"log"
	"net/http"
)

// DocumentController структура используется для конструктора контроллер
type DocumentController struct {
}

// NewDocumentController конструктор контроллера, возращающий экземпляр структуры Controller
func NewDocumentController() *DocumentController {
	return &DocumentController{}
}
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
