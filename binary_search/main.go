package main

import "fmt"

// https://www.hackerrank.com/challenges/tutorial-intro
func main() {
	var searched, numOfElem int
	fmt.Scan(&searched, &numOfElem)
	ar := make([]int, numOfElem)
	for i := 0; i < numOfElem; i++ {
		fmt.Scan(&ar[i])
	}
	fmt.Println(search(ar, len(ar)/2-1, searched))
}

func search(ar []int, i, searched int) int {
	if ar[i] == searched {
		return i
	} else if ar[i] > searched {
		return search(ar, i/2, searched)
	} else if ar[i] < searched {
		return search(ar, i+(len(ar)-i)/2, searched)
	}
	return -1

}
