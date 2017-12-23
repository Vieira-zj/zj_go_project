package demos

import (
	"fmt"
	"strings"
)

// demo 01, inner function
func printFormatName(firstName, lastName string) {
	getShortNameFn := func(firstName, lastName string) string {
		tmp := fmt.Sprintf("%c%c", firstName[0], lastName[0])
		return strings.ToUpper(tmp)
	}
	fmt.Printf("fname=%s, lname=%s, sname=%s\n",
		firstName, lastName, getShortNameFn(firstName, lastName))
}

func testPrintFormatName() {
	printFormatName("zheng", "jin")
}

// demo 02, struct init
type fullName struct {
	fName    string
	lName    string
	nickName string
}

func testPrintStructValue() {
	zjFullName := fullName{
		fName: "fname",
		lName: "lname",
	}
	fmt.Printf("full name: %v\n", zjFullName)

	zjFullName.nickName = "nick"
	fmt.Printf("full name with nick name: %v\n", zjFullName)
}

// MainDemo01 : main
func MainDemo01() {
	// testPrintFormatName()
	testPrintStructValue()

	fmt.Println("demo 01 done.")
}
