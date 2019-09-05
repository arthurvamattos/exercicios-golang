package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func maiorMenor(nums ...int) (int, int) {
	maior := 0
	menor := 0
	maior = nums[0]
	menor = nums[len(nums)-1]
	for _, num := range nums {
		if num > maior { maior = num }
		if num < menor { menor = num }
	}
	return maior, menor
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

	fmt.Println(maiorMenor(int(n1)))
	fmt.Println(maiorMenor(int(n1),int(n2)))
	fmt.Println(maiorMenor(int(n1),int(n2),int(n3)))
}