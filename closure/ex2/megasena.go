package main

import (
	"fmt"
    "math/rand"
    "time"
)

func main() {
	
	counter := megaSena()
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
}
// 2. Faça um programa que utilize closures para gerar
// sequências de 6 números pseudo-aleatórios não repetidos
// (a cada 6 chamadas a sequência é reinicializada).
func megaSena() func() int {
	nums := make([]int, 0, 6)

	return func() int {
		if len(nums) > 5 {
			nums = make([]int, 0, 6)
		}
		newNum := createNew(nums)
		nums = append(nums, newNum)
		return int(newNum)
	}
	
}

func createNew(nums []int) int {
	newNum := randInt(1, 61)
	if contains(nums, newNum) {
		createNew(nums)
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