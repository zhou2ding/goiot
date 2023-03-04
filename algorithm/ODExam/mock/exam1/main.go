package main

import "fmt"

func main() {
	var a int
	for {
		fmt.Scanln(&a)
		if a == 0 {
			break
		}
		fmt.Println(getMaxWater(a, 0))
	}
}

func getMaxWater(n, sum int) int {
	if n >= 3 {
		water := n / 3
		sum += water
		emptyBottle := n % 3
		return getMaxWater(water+emptyBottle, sum)
	} else if n == 2 {
		sum++
	}
	return sum
}
