package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	n1, _ := strconv.ParseInt(scanner.Text(), 10, 0)
	
	counter := generator(int(n1))
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
}

func generator(num int) func() int {
	n := num
	return func() int {
		n = n + n
		return n
	}
}