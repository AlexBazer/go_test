package main

import "fmt"
func main() {
    var str1, str2 string

    fmt.Scan(&str1, &str2)
    fmt.Println(findMaxSeqLen(str1, str2))
}

func findMaxSeqLen(str1, str2 string) int{
    if len(str1) == 0 || len(str2) == 0 {
        return 0
    }

    if str1[len(str1)-1] == str2[len(str2)-1]{
        return 1 + findMaxSeqLen(str1[:len(str1)-1], str2[:len(str2)-1])
    }
    return maxArg(findMaxSeqLen(str1, str2[:len(str2)-1]), findMaxSeqLen(str1[:len(str1)-1], str2))
}

func maxArg(args ...int) int{
    var maxVal int
    for i:=0;i<len(args);i++{
        if args[i]>maxVal{
            maxVal = args[i]
        }
    }
    return maxVal
}
