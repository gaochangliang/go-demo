package main

import "fmt"

type animal struct {
	sound string
}

type dog struct {
	animal
	friendly bool
}

type cat struct {
	animal
	annoying bool
}

func main() {
	fido := dog{animal{sound: "woof"}, true}
	fifi := cat{animal{sound: "meow"}, true}
	shadow := dog{animal{sound: "woof"}, true}
	miao := animal{"haha"}
	critters := []interface{}{fido, fifi, shadow, miao}
	fmt.Println(critters)
}
