package main

import "fmt"

// https://www.hackerrank.com/challenges/sherlock-and-valid-string

func newStrMap(str string) map[byte] int {
	// Create map sequence from string
	retMap := make(map[byte] int)
	for i := 0; i < len(str); i++ {
		retMap[str[i]]++
	}
	return retMap
}

func newIntMap(strMap map[byte]int) map[int] int {
	// Create map sequence from [byte]int map
	retMap := make(map[int] int)
    for _, value := range strMap{
        retMap[value] ++
    }

	return retMap
}

func main()  {
    var str string
    fmt.Scan(&str)
    // Gather chars in map
    strMap := newStrMap(str)
    // Collect num of char counts
    charCointsMap := make(map[int] int)
    for _, value := range strMap{
        charCointsMap[value] ++
    }

    result := "NO"
    if isValid(charCointsMap){
        result = "YES"
    }
    fmt.Println(result)
}

func isValid(charCointsMap map[int] int) bool{
    // There should be no more than 2 counts or 1 even better
    if len(charCointsMap) == 1{
        return true
    } else if len(charCointsMap) > 2{
        return false
    }
    // Assign 2 keys to variables
    var secondKey int
    oneValueKey := -1

    for key, value := range charCointsMap{
        if value == 1{
            oneValueKey = key
        } else{
            secondKey = key
        }
    }
    // There should be only count key with value one
    if oneValueKey == -1{
        return false
    }

    // In key with vith value one equal to 1 and difference between count keys equal to one
    if oneValueKey == 1 || oneValueKey - secondKey == 1{
        return true
    }

    return false

}
