package demos

import (
	"fmt"
	"strings"
)

func printFormatName(firstName, lastName string) {
	getShortNameFn := func(firstName, lastName string) string {
		tmp := fmt.Sprintf("%c%c", firstName[0], lastName[0])
		return strings.ToUpper(tmp)
	}
	fmt.Printf("fname=%s, lname=%s, sname=%s\n",
		firstName, lastName, getShortNameFn(firstName, lastName))
}

// MainDemo01 : main function for demo.
func MainDemo01() {
	printFormatName("zheng", "jin")

	fmt.Println("demo 01 done.")
}
