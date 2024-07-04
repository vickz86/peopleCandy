package main

import (
	"fmt"

	con "github.com/vickz86/peopleCandy/constructor"
	data "github.com/vickz86/peopleCandy/data"
)

// slice of People , create 2 default people
var Peoples = []data.People{
	{
		Name:    "victor",
		Money:   100,
		Candies: [3]int{2, 4, 6}},
	{

		Name:    "jonathan",
		Money:   200,
		Candies: [3]int{1, 3, 4}},
}

//start main

func main() {
	Peoples = con.CreatePeople(Peoples)
	fmt.Println(Peoples)

}
