package main

import (
	"fmt"
	"os"

	"github.com/vickz86/peopleCandy/candy"
	con "github.com/vickz86/peopleCandy/constructor"
	data "github.com/vickz86/peopleCandy/data"
)

// slice of People , create 2 default people
var Peoples = []data.People{
	{
		Name:    "victor",
		Money:   100,
		Candies: [3]int{5, 6, 9}},
	{

		Name:    "jonathan",
		Money:   200,
		Candies: [3]int{1, 3, 4}},
	{

		Name:    "amelie",
		Money:   300,
		Candies: [3]int{0, 2, 1}},
}

//start main

func main() {
	//print string
	possibleToDo := "choose between\n 0 exit\n1 add person\n2 print people\n3 print money of people"
	var toDo int

	//loop action
	for {
		//print and get action
		fmt.Print(possibleToDo)
		fmt.Scan(&toDo)
		//swictch what to do
		switch toDo {
		case 0:
			os.Exit(1)
		case 1:
			Peoples = con.CreatePeople(Peoples)
		case 2:
			fmt.Println(Peoples)
		case 3:
			PrintTotalMoney(Peoples)
		}

	}

}

/* =====FUNCTION===== */

func PrintTotalMoney(slicePeople []data.People) {
	for _, people := range slicePeople {
		totalmoney := candy.PeopleTotalMoney(people)

		fmt.Printf("%s has %d in cash and %d in total\n", people.Name, people.Money, totalmoney)

	}
}
