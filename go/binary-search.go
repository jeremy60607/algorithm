package main

import "fmt"

var searchTimes = 0

func main() {
	var ans *int
	ans = new(int)
	ans = binarySearch(13)
	fmt.Println(ans)
	fmt.Println(searchTimes)
}

var listNumbers = []int{1, 2, 5, 6, 9, 15, 25, 35}

func binarySearch(answer int) *int {

	low := 0
	high := len(listNumbers) - 1
	for low <= high {
		searchTimes++
		mid := (low + high) / 2
		guess := listNumbers[mid]
		fmt.Println("================")
		fmt.Println(high)
		fmt.Println(mid)
		fmt.Println(guess)
		fmt.Println("================")
		if guess == answer {
			return &guess
		} else if guess > answer {
			high = mid - 1
		} else if guess < answer {
			low = mid + 1
		} else {
			return nil
		}
	}
	return nil
}
