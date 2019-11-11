package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func exponentiation(x, n int64) int64 {
	if n == 0 {
		return 1
	} 
	return x * exponentiation(x, n - 1)
}


func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// recebendo e convertendo input
	scanner.Scan()
    n1, _ := strconv.ParseInt(scanner.Text(), 10, 0)

	scanner.Scan()
	n2, _ := strconv.ParseInt(scanner.Text(), 10, 0)
	
	fmt.Println(exponentiation(n1, n2))
}