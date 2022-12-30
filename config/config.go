package config

import (
	"errors"
	adapter "example.com/kate/adapterType"
	"fmt"
	"gopkg.in/ini.v1"
	"regexp"
)

type Config struct {
	ConcreteAdapterType adapter.AdapterType
	Url_add             string
}

var c Config

func Init() error {
	cfg, err := ini.Load("my.ini")
	if err != nil {
		fmt.Printf("Ошибка чтения файла my.ini: %v", err)
		return err
	}
	b := cfg.Section("").Key("ConcreteAdapterType").String()
	fmt.Println("Загружен флаг", b)
	m := cfg.Section("").Key("Url_Add").String()
	fmt.Println("Строка URL", m)
	c = Config{adapter.AdapterType(b), m}
	adapterTypeErr, urlErr := faultAdapterTypeUrl()
	if adapterTypeErr != nil {
		m := "Ошибка config, задан неправильный флаг в файле my.ini "
		fmt.Println(m, adapterTypeErr)
		return adapterTypeErr
	}
	if urlErr != nil {
		m := "Ошибка config, задан неправильный адрес в  my.ini "
		fmt.Println(m, urlErr)
		return urlErr
	}
	return err
}
func Get() *Config {
	return &c
}
func faultAdapterTypeUrl() (adapterTypeErr error, urlErr error) {
	if c.ConcreteAdapterType != adapter.File && c.ConcreteAdapterType != adapter.DB {
		adapterTypeErr = errors.New("Задан неправильный флаг в файле my.ini" + string(c.ConcreteAdapterType))
	}
	_, urlErr = regexp.MatchString("^http://./$", c.Url_add)
	if urlErr != nil {
		fmt.Println("Ошибка в URL")
	}
	return adapterTypeErr, urlErr
}
