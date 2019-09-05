package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)


func menorMaior(num1, num2 int) (int, int) {
	if num1 > num2 {
		return num2, num1
	}
	return num1, num2
}

func main()  {
	scanner := bufio.NewScanner(os.Stdin)

	// recebendo e convertendo input
	scanner.Scan()
	n1, _ := strconv.ParseInt(scanner.Text(), 10, 0)

	scanner.Scan()
	n2, _ := strconv.ParseInt(scanner.Text(), 10, 0)

	fmt.Println(menorMaior(int(n1),int(n2)))
}