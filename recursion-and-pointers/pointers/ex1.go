package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func tradeValues(num1, num2 *int64) {
	aux := *num1
	*num1 = *num2
	*num2 = aux 
}


func main()  {
	scanner := bufio.NewScanner(os.Stdin)

	// recebendo e convertendo input
	scanner.Scan()
    n1, _ := strconv.ParseInt(scanner.Text(), 10, 0)

	scanner.Scan()
	n2, _ := strconv.ParseInt(scanner.Text(), 10, 0)
	

	tradeValues(&n1, &n2)
	fmt.Println(n1, n2)
}