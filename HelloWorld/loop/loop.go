package main

import "fmt"

func main() {
	arr := make([]int, 10)

	// First loop to fill the values in the slice using range
	for i := range arr {
		arr[i] = i +1
	}
	fmt.Println("arr:", arr)

	sum := 0
	//2nd loop : using indexes
	for i:=0; i < len(arr); i++ {
		sum = sum + arr[i]
	}
	fmt.Println("Sum", sum)

	sum = 0
	j := 0
	//3rd type of loop: without init and post statement. Similar to while loop
	for ; j < 10; {
		sum += arr[j]
		j++
	}
	fmt.Println("Sum", sum)

	sum = 0
	j = 0
	// 4: Similar to above but more close to while loop
	for j < len(arr) {
		sum += arr[j]
		j++
	}
	fmt.Println("Sum:", sum)
}