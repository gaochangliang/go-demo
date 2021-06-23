package main

import "fmt"

type vehicles interface{}

type vehicle struct {
	Seats    int
	MaxSpeed int
	Color    string
}

type car struct {
	vehicle
	Wheels int
	Doors  int
}

type plane struct {
	vehicle
	Jet bool
}

type boat struct {
	vehicle
	Length int
}

func main() {
	prius := car{}
	tacoma := car{}
	bmw528 := car{}

	boeing747 := plane{}
	boeing757 := plane{}
	boeing767 := plane{}

	sanger := boat{}
	antique := boat{}
	malibu := boat{}

	rides := []vehicles{prius, tacoma, bmw528, boeing767, boeing757, boeing747, sanger, antique, malibu}

	for key, value := range rides {
		fmt.Println(key, " - ", value)
	}

}
