package utils

import (
	"fmt"
	"os"
	"zap/07/global"
)

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func CreateDir(path string) (bool, error) {
	fmt.Println(" CreateDir path", path)
	if ok, _ := PathExists(path); !ok {
		fmt.Printf("create %v directory\n", path)
		if err := os.Mkdir(global.CONFIG.Zap.Director, os.ModePerm); err != nil {
			return false, err
		}
	}
	return true, nil
}
