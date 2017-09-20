package main

import (
	"fmt"
)

func main() {
	a := 0
	b := true

	fmt.Println("Hello Esme!")

	for b {
		fmt.Print("Which times table would you like: ")
		fmt.Scan(&a)

		doTimesTable(a)

		var c string
		fmt.Printf("\n\nWould you like another? (y/n) ")
		fmt.Scan(&c)
		b = (c == "y")
	}
}

func doTimesTable(whichTable int) {
	var upper = 12

	//Can only handle up to 9999
	if whichTable*upper > 9999 {
		fmt.Println("Sorry, I can only work with tables where the largest product is less than ten thousand.")
		return
	}
	fmt.Printf("Okay. This is your %d times table \n\n", whichTable)
	aL := len(fmt.Sprintf("%d", upper))
	aP := len(fmt.Sprintf("%d", whichTable*upper))
	sF := fmt.Sprintf("%%d x %%%dd = %%%dd", aL, aP)

	for j := 0; j <= upper; j++ {
		fmt.Printf(sF, whichTable, j, whichTable*j)

		fmt.Println("")
	}
}

func productBreakdown(product int, maxProdLen int) {
	dhtu :=[]
	sP := fmt.Sprintf("%d", product)
	iP := len(sP)

}
