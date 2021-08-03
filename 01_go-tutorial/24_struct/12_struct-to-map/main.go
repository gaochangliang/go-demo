//struct to map
package main

import (
	"fmt"
	"github.com/spf13/viper"
	"reflect"
)

type Config struct {
	Person
	Animal
}

type Person struct {
	Name string
	Age  uint
}

type Animal struct {
	Age uint
}

func main() {
	var p = Person{
		Name: "kobe",
		Age:  41,
	}

	var a = Animal{
		Age: 5,
	}
	var conf Config
	conf.Person = p
	conf.Animal = a

	res := structToMap(conf)
	fmt.Println("res map", res)

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	for k, v := range res {
		viper.Set(k, v)
	}
	err := viper.WriteConfig()
	if err != nil {
		fmt.Println("WriteConfig err", err)
	}
}

func structToMap(obj interface{}) map[string]interface{} {
	obj1 := reflect.TypeOf(obj)
	obj2 := reflect.ValueOf(obj)
	data := make(map[string]interface{})
	fmt.Println("obj1.NumField() - ", obj1.NumField())
	for i := 0; i < obj1.NumField(); i++ {
		if obj1.Field(i).Tag.Get("mapstructure") != "" {
			data[obj1.Field(i).Tag.Get("mapstructure")] = obj2.Field(i).Interface()
		} else {
			data[obj1.Field(i).Name] = obj2.Field(i).Interface()
		}
	}
	return data
}
