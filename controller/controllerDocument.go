package controller

import (
	"example.com/kate/adapterType"
	"example.com/kate/model"
	"fmt"
	"github.com/gorilla/mux"
	"strconv"
	//"log"
	"net/http"
)

type Directory struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Directories []Directory
}

// DocumentController структура используется для конструктора контроллер
type DocumentController struct {
	model *model.Model
}

// NewDocumentController конструктор контроллера, возращающий экземпляр структуры Controller
func NewDocumentController(AdapterType adapterType.AdapterType) *DocumentController {
	return &DocumentController{model: model.NewModel(AdapterType)}
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
	res.Header().Set("Content-Type", "text/html")
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
	res.Header().Set("Content-Type", "text/html")
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
		fmt.Println(m, err)
		fmt.Fprintf(res, m, err)
		return
	}
	var sizeRows = params["sizeRows"]
	//конвертация string в int
	numRows, err := strconv.Atoi(sizeRows)
	if err != nil {
		m := "Ошибка перевода количества строк из string в int "
		fmt.Println(m, err)
		fmt.Fprintf(res, m, err)
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

// GetDocumentationTable  метод по созданию таблицы html в зависимости от встроенных структур БД
func (d *DocumentController) GetDocumentationTable(res http.ResponseWriter, req *http.Request) {
	//присваивание экземпляру структуры значений слайса структуры Document из метода controller.GetRezultDocumentation()
	documents, err := d.model.GetRezultDocumentation()
	if err != nil {
		m := "Ошибка выполнеия контроллера: %s"
		fmt.Println(m, err)
		fmt.Fprintf(res, m, err)
		return
	}
	//заголовок html таблицы
	tableHead := `<html lang="ru">
    <table border="1" width="600">
	<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
	<thead>
	<tr>
	<th>Documents</th>
	<th>Modules</th>
	<th>Errors</th>
	</tr>
	</thead>
	<tbody>
	<tr>`
	//вычисление количества документов в слайсе
	lenDocuments := len(documents)
	fmt.Println("Len документов", lenDocuments)
	//создание пустой строки
	var tableDocuments string
	//цикл для получения доктупа к каждому значению документа
	for keyDocument := range documents {
		fmt.Println("Количество документов", keyDocument)
		//module-объявление переменной
		module := documents[keyDocument].Modules
		//вычисление количество модулей вложенных в конкретный документ
		lenModules := len(module)
		fmt.Println("Len модулей", lenModules)
		//создание пустых строк
		var tableModules string
		var lenDocumentsString string
		var sliceErrors int
		var colspan string
		var colspanDoc string
		//цикл для получения доктупа к каждому значению модуля, в определенном документе
		for keyModules := range module {
			fmt.Println("Количество модулей", keyModules)
			//error-объявление переменной
			error := documents[keyDocument].Modules[keyModules].Errors
			//вычисление количества ошибок в конкретном документе и конкретном модуле
			lenErrors := len(error)
			//создание пустой строки
			var tableErrors string
			//цикл для получения доктупа к каждому значению ошибки
			for keyErrors := range error {
				fmt.Println("Количетво ошибок", keyErrors)
				//создание строки html таблицы ошибок
				tableError := `<td> ` + error[keyErrors].Title + `</td>
                </tr>`
				//объединение значений строк ошибок относящикся к одному модулю
				tableErrors = tableErrors + tableError
			}
			//условия слияния срок и столбцов, при условии, что у модуля отсутсвуют вложенные структуры
			if lenErrors == 0 {
				lenErrors = 1
				colspan = "2"

			} else {
				colspan = "1"
			}
			//перевод количества ошибок вложенных в конкретнный модуль из int в string
			lenErrorsString := strconv.Itoa(lenErrors)
			//счетчик ошибок принадлежащих конкретному документу и нескольким модулям
			sliceErrors = sliceErrors + lenErrors
			fmt.Println("Слайс ошибок", sliceErrors)
			fmt.Println("Len ошибок", lenErrors)

			//строка html таблицы объединябщая столько строк сколько ошибок вложены в конкретный модуль
			tableModule := `    <td rowspan=` + lenErrorsString + ` colspan=` + colspan + `>` + module[keyModules].
				Title + `</td>`
			//конкотинация строк
			tableModules = tableModules + tableModule + tableErrors
			//перевод int в string
			lenDocumentsString = strconv.Itoa(sliceErrors)

		}
		//условие слияние строк если к документу не привязан ни один модуль
		if sliceErrors == 0 {
			colspanDoc = `3`
			lenDocumentsString = `1`
		}
		fmt.Println("КОЛИЧЕСТВО ОБЪЕДИНЕННЫХ СТРОК В ДОКУМЕНТЕ", lenDocuments)
		//часть в таблице где добавляются документы
		tableDocument := `    <tr>
        <td rowspan=` + lenDocumentsString + ` colspan=` + colspanDoc + `>` + documents[keyDocument].Title + `</td>`
		tableDocuments = tableDocuments + tableDocument + tableModules
	}
	//окончание таблицы
	endTable := `
	</tbody>
    </table>
    </html>`
	//конкатинация строк из которой состоит таблица
	allTable := tableHead + tableDocuments + endTable
	//перевод строки в байты
	html := []byte(allTable)
	//отправка в браузер
	res.Write(html)

}

// метод для построение таблицы html
//при условии что вложенная структура имеет рекурсионный вид и количество связей не задано
func (d *DocumentController) GetDirectoriesTable(res http.ResponseWriter, req *http.Request) {
	//струтура для примера
	var directories = []Directory{
		{
			Id:    1,
			Title: "1",
			Directories: []Directory{
				{
					Id:    2,
					Title: "2",
					Directories: []Directory{
						{Id: 12, Title: "12"},
						{Id: 15, Title: "15"},
					},
				},
				{
					Id:    3,
					Title: "3",
					Directories: []Directory{
						{
							Id:    4,
							Title: "4",
							Directories: []Directory{
								{
									Id:    5,
									Title: "5",
									Directories: []Directory{
										{
											Id:    6,
											Title: "5 столбец 1 строка всего 2"},
										{
											Id:    7,
											Title: "7",
											Directories: []Directory{
												{
													Id: 8, Title: "8",
													Directories: []Directory{
														{Id: 9, Title: "9", Directories: []Directory{
															{
																Title:       "Уровень 1 Директория 1",
																Directories: nil,
															},
															{
																Title: "Уровень 1 Директория 2",
																Directories: []Directory{
																	{
																		Title: "Уровень 2 Директория 1",
																		Directories: []Directory{
																			{
																				Title:       "Уровень 3 Директория 1",
																				Directories: nil,
																			},
																			{
																				Title:       "Уровень 3 Директория 2",
																				Directories: nil,
																			},
																		},
																	},
																	{
																		Title:       "Уровень 2 Директория 2",
																		Directories: nil,
																	},
																},
															},
															{
																Title:       "Уровень 1 Директория 3",
																Directories: nil,
															},
														}}}},
												{
													Id: 10, Title: "10",
													Directories: []Directory{{Id: 11,
														Title: "11"},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
				{
					Id:    16,
					Title: "16",
					Directories: []Directory{
						{Id: 18, Title: "18"},
						{Id: 20, Title: "20"},
					},
				},
			},
		},
		{
			Title:       "Уровень 1 Директория 1",
			Directories: nil,
		},
		{
			Title: "Уровень 1 Директория 2",
			Directories: []Directory{
				{
					Title: "Уровень 2 Директория 1",
					Directories: []Directory{
						{
							Title:       "Уровень 3 Директория 1",
							Directories: nil,
						},
						{
							Title:       "Уровень 3 Директория 2",
							Directories: nil,
						},
					},
				},
				{
					Title:       "Уровень 2 Директория 2",
					Directories: nil,
				},
			},
		},
		{
			Title:       "Уровень 1 Директория 3",
			Directories: nil,
		},
	}
	//заголовок таблицы
	tableHead := `<html lang="ru">
	<table border="1" width="600">
	<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
	<thead>
	<tr>
	`
	var table string
	sliceColom := make([]int, 0)

	//цикл для извлечения количесва столбцов
	for key := range directories {
		level := 0
		getSliceColoms(directories[key].Directories, level, &sliceColom)
	}
	fmt.Println("sliceColom*", sliceColom)
	// вычисление максимального номера столбца
	var colomsNum int
	sumMax := (sliceColom)[0]
	for _, value := range sliceColom {
		colomsNum = 0
		if value > sumMax {
			sumMax = value
		}
		fmt.Println("sumMax", sumMax)
		colomsNum = colomsNum + sumMax
	}
	//отправка номера максимального столбца в функцию по построению заголовка таблицы
	stringBegin := colomsTable(colomsNum)
	//цикл для построения таблицы
	for key := range directories {
		//получает из функции количество строк которые необходимо объединить в первой строки
		firstRow := numHeads(directories[key].Directories)
		step := 1
		//перевод int в string
		firstRowString := strconv.Itoa(firstRow)
		//получает из функции количество колонок которые необходимо объединить у первой строки
		firstColspan, lineSeparation, _ := colspan(directories[key], colomsNum, step)
		if firstRow == 0 {
			firstRowString = `1`
		}
		//условие,
		//в случае если в первом столбце есть структура у которой нет вложенной структура столбцы объединяются до конца таблицы
		if directories[key].Directories == nil {
			firstColspan = colomsNum
		}
		firstColspanString := strconv.Itoa(firstColspan)
		if firstColspan == 0 {
			firstColspanString = `1`
		}
		//вывод певого столбца
		tableBeginer := `<td rowspan=` + firstRowString + ` colspan=` + firstColspanString + `>` + directories[key].
			Title + `</td>`
		//получение из функции строк всей таблицы
		tableMidle := paintRow(directories[key].Directories, colomsNum, step)
		//объединение строки из первого столбца со строй полоченной на предыдущем шаге
		table = table + tableBeginer + lineSeparation + tableMidle
	}
	//объединение всей таблицы
	tableAll := tableHead + stringBegin + `</tr>
</thead>
<tbody>
<tr>
` + table +
		`</tbody>
</table>
</html>`
	html := []byte(tableAll)
	//отправка в браузер
	res.Write(html)
}

//функция, которая рисует строки
func paintRow(directories []Directory, coloms int, step int) (tableMidls string) {

	var tableMidle string
	//цикл для вычисления "голов таблицы и их подсчета, а так же конкатенации строк таблицы
	for key := range directories {
		numRow := numHeads(directories[key].Directories)
		if numRow == 0 {
			numRow = 1
		}
		rowString := strconv.Itoa(numRow)
		numColom, lineSeparation, level := colspan(directories[key], coloms+1, step)
		colomString := strconv.Itoa(numColom)
		if numColom == 0 {
			colomString = `1`
		}
		tableString := `
<td rowspan=` + rowString + ` colspan=` + colomString + `>` + directories[key].Title + `</td>`

		table := paintRow(directories[key].Directories, coloms, level)
		//fmt.Println("Исходящая таблица внутри цикла", tableMidle)

		tableMidle = tableMidle + tableString + lineSeparation + table
		//tableMidls = tableMidls + tableString

	}

	return tableMidle
}

//функция которая показывает количество объединеных слобцов у каждой "головы"
func colspan(directories Directory, coloms int, step int) (int, string, int) {
	var lineSeparator string
	var lineSeparators string
	num := 0
	step = step + 1
	//условие если нашли "голову" после нее необходимо поставить символ новой строки
	if directories.Directories == nil {
		num = coloms - step
		lineSeparator = `
</tr>
<tr>`

	} else {
		num = 0
		lineSeparator = ``
	}
	lineSeparators = lineSeparators + lineSeparator

	//num = num + numRow

	return num, lineSeparators, step
}

//фукция считающая количество столбцов в каждой директории
func getSliceColoms(directories []Directory, level int, sliceColomn *[]int) {
	fmt.Println("Входящая директория в getSliceColoms", directories)

	level = level + 1
	*sliceColomn = append(*sliceColomn, level)

	for key := range directories {

		getSliceColoms(directories[key].Directories, level, sliceColomn)
	}
	fmt.Println("level", level)

}

//функция выдающая количество строчек при слиянии
func numHeads(directories []Directory) int {
	var num int
	for i := 0; i < len(directories); i++ {
		if directories[i].Directories == nil {
			num = 1 + num
		}
		numRow := numHeads(directories[i].Directories)
		num = num + numRow
	}
	return num
}

//функция которая рисует колонки таблицы
func colomsTable(colomsNum int) string {
	var tableBegins string
	for i := 0; i < colomsNum; i++ {
		numColom := strconv.Itoa(i)
		tableBegins = tableBegins +
			`<th>Directory` + numColom + `</th>
`
	}
	return tableBegins
}
