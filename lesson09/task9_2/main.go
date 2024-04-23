package main

import "fmt"

func main() {
	s := []int{1, 2, 3}
	for _, v := range s {
		fmt.Println("v1:", v)
	out:
		for _, v := range s {
			fmt.Println("\tv2:", v)
			for _, v := range s {
				fmt.Println("\t\tv3:", v)
				for i, v := range s {
					fmt.Println("\t\t\tv4:", v)
					if i == 1 {
						break out
					}
				}
			}
		}

	}
}
