// create a New instance of type People
package constructor

import (
	"fmt"
	"slices"

	sliceint "github.com/vickz86/GoFunctions/sliceInt"
	data "github.com/vickz86/peopleCandy/data"
)

// ValidName : check name doesnt exist in slice and return it
func ValidName(slicePeople []data.People) string {
	var newName string

	for {
		fmt.Println("type the name of new people:")
		fmt.Scan(&newName)

		//create a new slice of all existing name
		var existingName = []string{}
		for _, people := range slicePeople {
			existingName = append(existingName, people.Name)
		}

		//check if name you have typed is in the list
		if slices.Contains(existingName, newName) {
			fmt.Println(newName, "already exist , type newOne")
			continue
		}

		return newName

	}

}

// AskMoneyCandy : ask for money and candy
func AskMoneyCandy() (int, [3]int) {
	var money int
	var nbCandy [3]int

	//ask for money
	fmt.Println("how much money do you have:")
	fmt.Scan(&money)

	//ask for candie

	fmt.Println("type red,blue,greenm candy you have:")
	tempSlice := sliceint.CreateSliceFixNumber(3)
	nbCandy = [3]int(tempSlice)

	return money, nbCandy

}

// CreatePeople : create a People and add it to the slice of people
func CreatePeople(slicePeople []data.People) []data.People {
	//variable to return

	name := ValidName(slicePeople)
	money, candy := AskMoneyCandy()

	var returnPeople = data.People{
		Name:    name,
		Money:   money,
		Candies: candy,
	}

	slicePeople = append(slicePeople, returnPeople)

	return slicePeople

}
