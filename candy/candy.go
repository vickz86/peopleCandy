package candy

import "github.com/vickz86/peopleCandy/data"

//candy value
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
