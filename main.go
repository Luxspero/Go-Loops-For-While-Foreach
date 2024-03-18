package main

import "fmt"

func main() {
	hallo := "Hallo World"

	// *** FOR ***
	//single variable loops
	for i := 0; i < 10; i++ {
		fmt.Println(i, hallo)
	}
	//Multiple varible loops
	fmt.Println("Multiple")
	for i, j := 0, 1; i < 10 && j < 10; i, j = i+1, j+1 {
		fmt.Println(i, j, hallo)
	}

	// *** FOREACH ***
	for i := 0; i < len(hallo); i++ {
		fmt.Printf("%d %c", i, hallo[i])
		fmt.Println()
	}
	fmt.Println()
	//Other Structure
	for i, c := range hallo {
		fmt.Printf("%d %c", i, c)
		fmt.Println()
	}

	//break
	for _, c := range hallo {
		//If there is a space in C then coding will stop
		if c == ' ' {
			break
		}
		fmt.Printf("%c", c)
	}
	fmt.Println()

	//continue
	for _, c := range hallo {
		//If there is a space in C then the coding will continue and delete the space
		if c == ' ' {
			continue
		}
		fmt.Printf("%c", c)
	}
	fmt.Println()

	//label
ifForloop:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if j == 3 {
				break ifForloop
			}
			fmt.Printf("%d%d ", i, j)
		}
		fmt.Println()
	}
}
