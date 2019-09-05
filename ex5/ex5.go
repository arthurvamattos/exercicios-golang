package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func mediaPositivos(nums ...int) int {
	media := 0
	for _, num := range nums {
		if num >= 0 {
			media += num
		}
	}
	if media > 0 { media = media/len(nums) }
	return media
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

	fmt.Println(mediaPositivos(int(n1)))
	fmt.Println(mediaPositivos(int(n1),int(n2)))
	fmt.Println(mediaPositivos(int(n1),int(n2),int(n3)))
}