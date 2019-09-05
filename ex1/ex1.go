package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func menor(num1, num2 int) int {
	if num1 > num2 {
		return num2
	}
	return num1
}

func main()  {
	scanner := bufio.NewScanner(os.Stdin)

	// recebendo e convertendo input
	scanner.Scan()
    n1, _ := strconv.ParseInt(scanner.Text(), 10, 0)

	scanner.Scan()
    n2, _ := strconv.ParseInt(scanner.Text(), 10, 0)

	fmt.Println(menor(int(n1),int(n2)))
}
