package config

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

func LoadConfig(filenames string) (*Config, []string, error) {
	files, ignored := getFiles(filenames)
	if files != nil {
		if err := godotenv.Load(files...); err != nil {
			return nil, ignored, err
		}
	}

	var c Config
	walk(&c)

	return &c, ignored, nil
}

func getFiles(filenames string) (files []string, ignored []string) {
	list := strings.Split(filenames, ";")

	for _, v := range list {
		if v != "" {
			if _, err := os.Open(v); err != nil {
				ignored = append(ignored, v)
			} else {
				files = append(files, v)
			}
		}
	}

	return
}

func walk(e any) {
	v := reflect.ValueOf(e)
	t := reflect.TypeOf(e)

	if v.Kind() == reflect.Pointer && t.Kind() == reflect.Pointer {
		v = v.Elem()
		t = t.Elem()
	}

	for i := 0; i < v.NumField(); i++ {
		value := v.Field(i)

		if value.Kind() == reflect.Struct {
			walk(value.Addr().Interface())
		} else {
			set(t.Field(i), value, value.Kind())
		}
	}
}

func set(t reflect.StructField, v reflect.Value, kind reflect.Kind) {
	tag, ok := t.Tag.Lookup("env_config")
	if !ok {
		panic("not all Config fields have \"env_config\" tag")
	}

	value, ok := os.LookupEnv(tag)
	if !ok {
		panic(fmt.Sprintf("env variable %q does not exist", tag))
	}

	switch kind {
	case reflect.String:
		v.SetString(value)
	case reflect.Bool:
		if val, err := strconv.ParseBool(value); err == nil {
			v.SetBool(val)
		} else {
			panic(err)
		}
	default:
		panic("not supported type of Config struct field")
	}
}
