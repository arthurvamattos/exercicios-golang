package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func sumAll(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}


func main()  {
	scanner := bufio.NewScanner(os.Stdin)

	// recebendo e convertendo input
	scanner.Scan()
    n1, _ := strconv.ParseInt(scanner.Text(), 10, 0)

	scanner.Scan()
	n2, _ := strconv.ParseInt(scanner.Text(), 10, 0)
	
	scanner.Scan()
    n3, _ := strconv.ParseInt(scanner.Text(), 10, 0)

	fmt.Println(sumAll(int(n1)))
	fmt.Println(sumAll(int(n1),int(n2)))
	fmt.Println(sumAll(int(n1),int(n2),int(n3)))
}