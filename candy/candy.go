package candy

import (
	sliceint "github.com/vickz86/GoFunctions/sliceInt"
	"github.com/vickz86/peopleCandy/data"
)

// candy value
var CandyVal = [3]int{5, 10, 20}

// get total money of poeple = money + values of candys
func PeopleTotalMoney(aPeople data.People) int {
	var candyMoney int

	//get candy money
	for index, _ := range CandyVal {
		candyMoney += CandyVal[index] * aPeople.Candies[index]
	}

	//return candyMoney + money
	return aPeople.Money + candyMoney

}

// WhoMoreCandy : print and return name and nb of people who as more candy
func WhoMoreCandy(allPeople []data.People) (string, int) {
	//create a new slice of int with total of candy
	var peopleTotalCandy []int

	//loop through all People and create new slice of total candy
	for _, p := range allPeople {
		sliceC := p.Candies[:]
		total := sliceint.SumSlice(sliceC)
		peopleTotalCandy = append(peopleTotalCandy, total)

	}

	//get the index and value of max int the slice
	index, maxVal := sliceint.HighestValueIndex(peopleTotalCandy)

	//get name from index
	theName := allPeople[index].Name

	//return name and HighestValue
	return theName, maxVal

}
