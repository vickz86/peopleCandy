package candy

import (
	"fmt"
	"testing"

	"github.com/vickz86/peopleCandy/data"
)

var testSlicePeople = []data.People{
	{Name: "Bob", Money: 200, Candies: [3]int{1, 1, 2}},
	{Name: "James", Money: 100, Candies: [3]int{3, 2, 4}}}

func TestPeopleTotalMoney(t *testing.T) {
	want := 255
	got := PeopleTotalMoney(testSlicePeople[0])
	if want == got {
		fmt.Println("PeopleTotalMoney has passed")
	} else {
		fmt.Printf("error PeopleTotalMoney ,got :%v and want %v", got, want)
	}
}

func TestWhoMoreCandy(t *testing.T) {
	nameWant := "James"
	nameGot, _ := WhoMoreCandy(testSlicePeople)
	if nameWant == nameGot {
		fmt.Println("WhoMoreCandy has passed")
	} else {
		fmt.Printf("error WhoMoreCandy ,got :%v and want %v", nameGot, nameWant)
	}

}
