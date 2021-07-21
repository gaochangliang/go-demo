package main

import (
	"encoding/json"
	"fmt"
	"github.com/mitchellh/mapstructure"
)

type Person struct {
	Name string
}

/*

为了正确解码需要嵌套person格式的数据
map[string]interface{} {
  "person": map[string]interface{}{"name": "dj"},
}

*/
type Friend1 struct {
	Person
}

/*设置mapstructure:",squash"将该结构体的字段提到父结构中

map[string]interface{}{
 "name": "dj",
}
不需要嵌套json
*/
type Friend2 struct {
	Person `mapstructure:",squash"`
}

func main() {

	datas := []string{
		`{ 
			  "type": "friend1",
			  "person": {
				"name":"dj"
			  }
         }`,
		`{
			  "type": "friend2",
			  "name": "dj2"
    	  }`,
	}

	for _, data := range datas {
		var m map[string]interface{}
		err := json.Unmarshal([]byte(data), &m)
		if err != nil {
			fmt.Println("json unmarshal err", err)
		}
		switch m["type"].(string) {
		case "friend1":
			var f1 Friend1
			mapstructure.Decode(m, &f1)
			fmt.Println("friend1", f1)

		case "friend2":
			var f2 Friend2
			mapstructure.Decode(m, &f2)
			fmt.Println("friend2", f2)

		}
	}

}

/*

type Friend1 struct {
	Person
}

type Friend1 struct {
	Person Person
}

These two ways of defining Friend are the same for mapstructure
*/
