package candy

import (
	"fmt"

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

// BuyCandys : buy candies with money , and return people
func BuyCandys(people data.People) data.People {
	// var
	var toBuy int
	returnMoney := people.Money
	returnCandies := people.Candies
	loopCounter := 0

	//loop to buy each kind of Candy
	for {

		// input candy to buy
		fmt.Printf("how many Candy %d you want to buy: ", loopCounter)
		fmt.Scan(&toBuy)

		//create buy value
		buyValue := toBuy * CandyVal[loopCounter]

		//check has enough money
		if buyValue > people.Money {
			fmt.Println("not enough money!")
			continue
		} else /* buy and substrac and add to candies */ {
			//substract Money
			returnMoney -= buyValue
			//add candies
			returnCandies[loopCounter] += toBuy

		}

		//check to break
		if loopCounter == (len(people.Candies) - 1) {
			break
		}

		//incrementCoutner
		loopCounter++

	}

	//create an instance of people based on new value
	var returnPeople = data.People{Name: people.Name, Money: returnMoney, Candies: returnCandies}

	//return the returnPeople
	return returnPeople

}
