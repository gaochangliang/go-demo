package main

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

type appConfig struct {
	Server   server   `mapstructure:"server" json:"server" yaml:"server"`
	Database database `mapstructure:"database" json:"database" yaml:"database"`
}

type server struct {
	Httpport     int `mapstructure:"httpport" json:"httpport" yaml:"httpport"`
	Readtimeout  int `mapstructure:"readtimeout" json:"readtimeout" yaml:"readtimeout"`
	Writetimeout int `mapstructure:"writetimeout" json:"writetimeout" yaml:"writetimeout"`
}

type database struct {
	User     string `mapstructure:"user" json:"user" yaml:"user"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
	Host     string `mapstructure:"host" json:"host" yaml:"host"`
	Name     string `mapstructure:"name" json:"name" yaml:"name"`
}

func main() {
	viper.SetConfigName("conf") // name of config file (without extension)
	viper.SetConfigType("yaml") // REQUIRED if the config file does not have the extension in the name

	//also we can set filename
	//viper.SetConfigFile("./conf/conf.yaml")

	//search paths can be set more than one, viper will look for them in order according to the set order.
	viper.AddConfigPath("../conf") // call multiple times to add many search paths

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatal("Config file not found; ignore error if desired", err)
		} else {
			log.Fatal("Config file was found but another error was produced", err)
		}
	}

	var appConf appConfig
	_ = viper.Unmarshal(&appConf)
	fmt.Printf("%+v", appConf)

}
