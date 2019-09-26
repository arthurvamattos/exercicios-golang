package main

import (
	"fmt"
    "math/rand"
    "time"
)

func main() {
	
	sorteador := megaSena()
	fmt.Println(sorteador())
	fmt.Println(sorteador())
	fmt.Println(sorteador())
	fmt.Println(sorteador())
	fmt.Println(sorteador())
	fmt.Println(sorteador())
	fmt.Println(sorteador())
	fmt.Println(sorteador())
	fmt.Println(sorteador())
	fmt.Println(sorteador())
}

func megaSena() func() int {
	nums := make([]int, 0, 6)

	return func() int {
		newNum := createNew(nums)
		if len(nums) > 5 {
			nums = make([]int, 0, 6)
		}
		nums = append(nums, newNum)
		return int(newNum)
	}
	
}

func createNew(nums []int) int {
	newNum := randInt(1, 61)
	if contains(nums, newNum) {
		newNum = createNew(nums)
	}
	return newNum
}

func contains(nums []int, newNum int) bool {
	for _, num := range nums {
		if num == newNum {
			return true
		}
	}
	return false
}
func randInt(min int, max int) int {
    rand.Seed(time.Now().UTC().UnixNano())
    return min + rand.Intn(max-min)
}