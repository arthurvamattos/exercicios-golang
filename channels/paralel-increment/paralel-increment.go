package main

import (
	"fmt"
)

func main() {
	i := 0
    for i <= 10000 {	
		go func() { i++ }()
		go func() { i++ }()
		fmt.Println(i)
    }
}