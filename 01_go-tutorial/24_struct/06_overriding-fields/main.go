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

	//fields and methods of the inner-type are promoted to the outer-type
	//内部嵌入类型的字段和方法被提升到外部类型
	fmt.Println("p1 first-", p1.first, "- p1 person first", p1.person.first)
	fmt.Println("p2 first-", p2.first, "- p2 person first", p2.person.first)

}
