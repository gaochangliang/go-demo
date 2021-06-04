package main

import "fmt"

type person struct {
	first  string
	last   string
	number int
}

type doubleZero struct {
	person
	LicenseToKill bool
}

func (p person) fullName() string {
	return p.first + p.last
}

func main() {
	p1 := doubleZero{
		person: person{
			first:  "jordan",
			last:   "Michael",
			number: 23,
		},
		LicenseToKill: true,
	}
	p2 := doubleZero{
		person: person{
			first:  "kobe",
			last:   "bryant",
			number: 24,
		},
		LicenseToKill: false,
	}

	fmt.Println("p1 info", p1.first, p1.last, p1.number, p1.LicenseToKill)
	fmt.Println("p2 info", p2.first, p2.last, p2.number, p2.LicenseToKill)

}
