package main

type vehicles interface{}

type animal struct {
	sound string
}

type dog struct {
}

type plane struct {
}

type boat struct {
	vehicle
	Length int
}

func main() {

}
