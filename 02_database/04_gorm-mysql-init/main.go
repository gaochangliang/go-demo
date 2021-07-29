//create initialize gorm mysql
package main

import (
	"database/04/global"
	"database/04/initalize"
)

func main() {
	global.GLOBAL_CONFIG = initalize.Gorm()
}
