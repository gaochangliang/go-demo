package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
)

func main() {
	//set conf name type
	viper.SetConfigName("conf") // name of config file (without extension)
	viper.SetConfigType("yaml") // REQUIRED if the config file does not have the extension in the name

	//also we can set filename
	//viper.SetConfigFile("./conf/conf.yaml")

	//search paths can be set more than one, viper will look for them in order according to the set order.
	viper.AddConfigPath("../conf") // call multiple times to add many search paths
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
	//After setting the PORT change server.httpPort will also change, serverhttpport and port is equivalent
	fmt.Println("server.httpport", viper.Get("server.httpport"))
	viper.RegisterAlias("server.httpport", "port")
	viper.Set("port", 80)
	fmt.Printf("server httpPort %v port %v", viper.Get("server.httpport"), viper.Get("port"))

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
}
