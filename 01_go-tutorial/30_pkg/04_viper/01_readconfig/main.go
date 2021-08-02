package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
)

func main() {
	//set config name type
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name

	//also we can set filename
	//viper.SetConfigFile("./config/config.yaml")

	//search paths can be set more than one, viper will look for them in order according to the set order.
	viper.AddConfigPath("../config") // call multiple times to add many search paths
	//viper.AddConfigPath(".")        // optionally look for config in the working directory

	//simple variable set default value
	//设置优先级
	/*
		explicit call to Set
		flag
		env
		config
		key/value store
		default
	*/

	// Find and read the config file
	viper.SetDefault("runmode", "dev")
	viper.SetDefault("app.page_size", 1000)
	viper.Set("app.page_size", 10000)

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatal("Config file not found; ignore error if desired", err)
		} else {
			log.Fatal("Config file was found but another error was produced", err)
		}
	}

	//Viper configuration keys are case insensitive 不区分大小写
	fmt.Println("runmode - ", viper.Get("runmode"))

	//Viper can access a nested field by passing a . delimited path of keys:
	fmt.Println("app.page_size - ", viper.Get("app.page_size"))

	//Internal format conversion
	fmt.Println("app.page_size - ", viper.GetInt64("app.page_size"))

	//get object  	map[jwt_secret:23347$040412 page_size:10]
	fmt.Println("app - ", viper.Get("app"))

	//get arr 		[kobe jordan wade]
	fmt.Println("developer - ", viper.Get("developer"))

	//Registering and Using Aliases
	//Aliases permit a single value to be referenced by multiple keys
	//After setting the PORT change api.httpPort will also change, serverhttpport and port is equivalent
	fmt.Println("api.httpport", viper.Get("api.httpport"))
	viper.RegisterAlias("api.httpport", "port")
	viper.Set("port", 80)
	fmt.Printf("api httpPort %v port %v", viper.Get("api.httpport"), viper.Get("port"))

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
}
