package main

import (
	"fmt"
)

func main()  {
	a := [...]int{5,8,100,66,35,27,10}
	fmt.Println(a)
	num := len(a)

	for i := 0; i < num; i++ {
		for j := i + 1; j < num; j++ {
			if a[i] < a[j] {
				a[i],a[j] = a[j],a[i]
			}
		}
	}
	fmt.Println(a)
}
