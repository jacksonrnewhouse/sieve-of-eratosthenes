package main

import (
	"os"
	"fmt"
	"strconv"
	"math"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please enter exactly one argument, an integer limit",
					"to the size of the primes generated.")
		return
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	if n < 2 {
		fmt.Println("The integer limit must be greater than or equal to two.")
	}

	composites := make([]bool, n+1)
        composites[0] = true
	composites[1] = true
	for i := 2; i < int(math.Sqrt(float64(n))) + 1; i++ {
		if !composites[i] {
			for x := i * 2; x <= n; x+=i {
				composites[x] = true
			}
		}
	}

	for i := 0; i <= n; i++ {
		if !composites[i] {
			fmt.Println(i)
		}
	}
}
