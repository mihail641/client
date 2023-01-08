module example.com/kate

go 1.18

require (
	github.com/gorilla/mux v1.8.0
	gopkg.in/ini.v1 v1.67.0
	example.com/projectApiClient v0.0.0
)
replace (
	example.com/projectApiClient => ../projectApiClient
)

