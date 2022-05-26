package main

import "fmt"

func main() {
	var i int
	j := 0
	for i = 0; i<=10; i++ {
	    for ;j<=10;j++ {
		fmt.Printf("%d x %d = %d\n", i, j, i*j)
	    }
	}
}
