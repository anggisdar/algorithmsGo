package main

import "fmt"

func engineerTeam(num int) {
	if num%3 == 0 && num%5 == 0 {
		fmt.Println("EngineerTeam")
	} else if num%3 == 0 {
		fmt.Println("Engineer")
	} else if num%5 == 0 {
		fmt.Println("Team")
	} else {
		fmt.Println(num)
	}
}

func main() {
	for i := 1; i <= 100; i++ {
		engineerTeam(i)
	}
}
