package controller

import (
	"example.com/projectApiClient"
	"regexp"
	"testing"
)

func Test_getTableResult(t *testing.T) {
	tests := []struct {
		directories []projectApiClient.Directory
		want        string
	}{
		{directories: []projectApiClient.Directory{
			{Id: 1, Title: "Music", Directories: []projectApiClient.Directory{
				{Id: 2, Title: "Папка 3", Directories: []projectApiClient.Directory{
					{Id: 3, Title: "Папка 4", Directories: []projectApiClient.Directory{
						{Id: 4, Title: "Папка 5"},
						{Id: 4, Title: "Папка 6"},
					},
					},
				},
				},
			},
			},
			{
				Id: 3, Title: "папка 1", Directories: []projectApiClient.Directory{
					{Id: 4, Title: "Папка 1_2", Directories: []projectApiClient.Directory{
						{Id: 5, Title: "New directory"},
						{Id: 6, Title: "новаяя папка"},
					},
					},
				},
			},
		},
			want: `<html lang=\"ru\">
<table border=\"1\" width=\"600\">
<meta http-equiv=\"Content-Type\" content=\"text/html; charset=utf-8\">
<thead>
<tr>
<th>Directory0</th>
<th>Directory1</th>
<th>Directory2</th>
<th>Directory3</th>
</tr>
</thead>
<tbody>
<tr>
<td rowspan=2 colspan=1>Music</td>
<td rowspan=2 colspan=1>Папка 3</td>
<td rowspan=2 colspan=1>Папка 4</td>
<td rowspan=1 colspan=1>Папка 5</td>
</tr>
<tr>
<td rowspan=1 colspan=1>Папка 6</td>
</tr>
<tr><td rowspan=2 colspan=1>папка 1</td>
<td rowspan=2 colspan=1>Папка 1_2</td>
<td rowspan=1 colspan=2>New directory</td>
</tr>
<tr>
<td rowspan=1 colspan=2>новаяя папка</td>
</tr>
<tr></tbody>
</table>
</html>`,
		},
		{directories: []projectApiClient.Directory{
			{
				Id:    1,
				Title: "1",
				Directories: []projectApiClient.Directory{
					{
						Id:    2,
						Title: "2",
						Directories: []projectApiClient.Directory{
							{Id: 12, Title: "12"},
							{Id: 15, Title: "15"},
						},
					},
					{
						Id:    3,
						Title: "3",
						Directories: []projectApiClient.Directory{
							{
								Id:    4,
								Title: "4",
								Directories: []projectApiClient.Directory{
									{
										Id:    5,
										Title: "5",
										Directories: []projectApiClient.Directory{
											{
												Id:    6,
												Title: "5 столбец 1 строка всего 2"},
											{
												Id:    7,
												Title: "7",
												Directories: []projectApiClient.Directory{
													{
														Id: 8, Title: "8",
														Directories: []projectApiClient.Directory{
															{Id: 9, Title: "9", Directories: []projectApiClient.Directory{
																{
																	Title:       "Уровень 1 Директория 1",
																	Directories: nil,
																},
																{
																	Title: "Уровень 1 Директория 2",
																	Directories: []projectApiClient.Directory{
																		{
																			Title: "Уровень 2 Директория 1",
																			Directories: []projectApiClient.Directory{
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
														Directories: []projectApiClient.Directory{{Id: 11,
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
						Directories: []projectApiClient.Directory{
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
				Directories: []projectApiClient.Directory{
					{
						Title: "Уровень 2 Директория 1",
						Directories: []projectApiClient.Directory{
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
		},
			want: `html lang=\"ru\">
<table border=\"1\" width=\"600\">
<meta http-equiv=\"Content-Type\" content=\"text/html; charset=utf-8\">\n\t<thead>
<tr>
<th>Directory0</th>
<th>Directory1</th>
<th>Directory2</th>
<th>Directory3</th>
<th>Directory4</th>
<th>Directory5</th>
<th>Directory6</th>
<th>Directory7</th>
<th>Directory8</th>
<th>Directory9</th>
</tr>
</thead>
<tbody>
<tr>
<td rowspan=11 colspan=1>1</td>
<td rowspan=2 colspan=1>2</td>
<td rowspan=1 colspan=8>12</td>
</tr>
<tr>
<td rowspan=1 colspan=8>15</td>
</tr>
<tr>
<td rowspan=7 colspan=1>3</td>
<td rowspan=7 colspan=1>4</td>
<td rowspan=7 colspan=1>5</td>
<td rowspan=1 colspan=6>5 столбец 1 строка всего 2</td>
</tr>\n<tr>\n<td rowspan=6 colspan=1>7</td>
<td rowspan=5 colspan=1>8</td>
<td rowspan=5 colspan=1>9</td>
<td rowspan=1 colspan=3>Уровень 1 Директория 1</td>
</tr>
<tr>
<td rowspan=3 colspan=1>Уровень 1 Директория 2</td>
<td rowspan=2 colspan=1>Уровень 2 Директория 1</td>
<td rowspan=1 colspan=1>Уровень 3 Директория 1</td>
</tr>
<tr>
<td rowspan=1 colspan=1>Уровень 3 Директория 2</td>
</tr>
<tr>
<td rowspan=1 colspan=2>Уровень 2 Директория 2</td>
</tr>
<tr>
<td rowspan=1 colspan=3>Уровень 1 Директория 3</td>
</tr>
<tr>
<td rowspan=1 colspan=1>10</td>
<td rowspan=1 colspan=4>11</td>
</tr>
<tr>
<td rowspan=2 colspan=1>16</td>
<td rowspan=1 colspan=8>18</td>
</tr>
<tr>
<td rowspan=1 colspan=8>20</td>
</tr>
<tr><td rowspan=1 colspan=10>Уровень 1 Директория 1</td>
</tr>
<tr><td rowspan=3 colspan=1>Уровень 1 Директория 2</td>
<td rowspan=2 colspan=1>Уровень 2 Директория 1</td>
<td rowspan=1 colspan=8>Уровень 3 Директория 1</td>
</tr>
<tr>
<td rowspan=1 colspan=8>Уровень 3 Директория 2</td>
</tr>
<tr>
<td rowspan=1 colspan=9>Уровень 2 Директория 2</td>
</tr>
<tr><td rowspan=1 colspan=10>Уровень 1 Директория 3</td>
</tr>
<tr></tbody>
</table>
</html>`,
		},
	}
	for _, valueTests := range tests {
		re := regexp.MustCompile(`[[:space:]]`)
		directoriesWord := re.ReplaceAllString(valueTests.want, "")
		got := getTableResult(valueTests.directories)
		gotWord := re.ReplaceAllString(got, "")
		if gotWord != directoriesWord {
			t.Errorf("Что то пошло не так получила: %q, хочу получить %q", got, valueTests.want)
		}
	}
}
