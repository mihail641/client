package config

import (
	adapter "example.com/kate/adapter"
	"fmt"
	"gopkg.in/ini.v1"
	"os"
	"regexp"
)

type Config struct {
	ConcreteAdapterType adapter.AdapterType
	Url_add             string
}

var c Config

func init() {
	cfg, err := ini.Load("my.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	b := cfg.Section("").Key("ConcreteAdapterType").String()
	m := cfg.Section("").Key("Url_Add").String()

	c = Config{
		adapter.AdapterType(b),
		m,
	}
	fmt.Println("Из ини файла", c.Url_add)
	fmt.Println("Из ини файла ", c.ConcreteAdapterType)

}
func Get() Config {
	return c
}
func FaultAdapterTypeUrl(c Config) {

	if c.ConcreteAdapterType != adapter.File && c.ConcreteAdapterType != adapter.DB {

		fmt.Println("Задан неправильный флаг в файле my.ini", c.ConcreteAdapterType)
	}
	err, _ := regexp.MatchString("^http://./$", c.Url_add)
	if err != true {
		fmt.Println("Ошибка в URL")
	}

}
