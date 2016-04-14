package main

import "fmt"

// https://www.hackerrank.com/challenges/common-child

func main() {
	var str1, str2 string

	fmt.Scan(&str1, &str2)
	// Create store 2d array to sote recursion values to gain speed
	store := make([][]int, len(str1))
	for i := 0; i < len(str1); i++ {
		store[i] = make([]int, len(str2))
		for j := 0; j < len(str2); j++ {
			store[i][j] = -1
		}
	}
	fmt.Println(findMaxSeqLen(str1, str2, len(str1)-1, len(str2)-1, store))
}

func findMaxSeqLen(str1, str2 string, i, j int, store [][]int) int {
	if i+1 == 0 || j+1 == 0 {
		return 0
	}
	if store[i][j] >= 0 {
		return store[i][j]
	}

	var storeVal int
	if str1[i] == str2[j] {
		storeVal = 1 + findMaxSeqLen(str1, str2, i-1, j-1, store)
	} else {
		storeVal = maxArg(findMaxSeqLen(str1, str2, i, j-1, store), findMaxSeqLen(str1, str2, i-1, j, store))
	}
	store[i][j] = storeVal
	return storeVal
}

func maxArg(args ...int) int {
	var maxVal int
	for i := 0; i < len(args); i++ {
		if args[i] > maxVal {
			maxVal = args[i]
		}
	}
	return maxVal
}
