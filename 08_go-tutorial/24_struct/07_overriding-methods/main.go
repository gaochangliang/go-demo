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
	first         string
}

func (p person) greeting() string {
	return "I'm just a regular person."
}

func (d doubleZero) greeting() string {
	return "Miss Moneypenny, so good to see you"
}

func main() {
	p1 := doubleZero{
		person: person{
			first:  "jordan",
			last:   "Michael",
			number: 23,
		},
		LicenseToKill: true,
		first:         "Double Zero Seven",
	}
	p2 := doubleZero{
		person: person{
			first:  "kobe",
			last:   "bryant",
			number: 24,
		},
		LicenseToKill: false,
		first:         "If looks could kill",
	}

	fmt.Println("p1 greeting - ", p1.greeting())
	fmt.Println("p1 person greeting - ", p1.person.greeting())
	fmt.Println("p2 person greeting- ", p2.person.greeting())

}
