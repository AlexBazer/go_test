package main

import "fmt"

type mapSeq map[byte]int

func (ms *mapSeq) copy() mapSeq {
	ret := make(mapSeq)
	for key, value := range *ms {
		ret[key] = value
	}
	return ret
}

func (ms *mapSeq) update(with mapSeq) mapSeq {
	ret := ms.copy()
	for key, value := range with {
		ret[key] += value
	}
	return ret
}

func (ms *mapSeq) len() int {
	var seqLen int
	for _, val := range *ms {
		seqLen += val
	}
	return seqLen
}

func (ms mapSeq) compare(with mapSeq) bool {
	for key, value := range ms {
		withValue, exists := with[key]
		if !exists || withValue != value {
			return false
		}
	}

	for key, value := range with {
		msValue, exists := ms[key]
		if !exists || msValue != value {
			return false
		}
	}

	return true
}

func main() {
	var numSimbols int
	var str string
	fmt.Scan(&numSimbols, &str)

	strMap := strToMap(str)
	baseMap := getBaseTestSeq(strMap, numSimbols)
	fmt.Println(baseMap.len(), testSeqInStr(baseMap, str))
	return
	if testSeqInStr(baseMap, str) {
		fmt.Println("0")
	} else {
		testSeqs := getSeqs(strMap, baseMap)
		var minSeqLen int
		for i := 0; i < len(testSeqs); i++ {
			if testSeqInStr(testSeqs[i], str) {
				minSeqLen = testSeqs[i].len()
				break
			}
		}
		if minSeqLen > 0 {
			fmt.Println(minSeqLen)
		} else {
			fmt.Println("BOOL Sheat!")
		}
	}
}

func strToMap(str string) mapSeq {
	retMap := make(mapSeq)
	for i := 0; i < len(str); i++ {
		retMap[str[i]]++
	}
	return retMap
}

func getBaseTestSeq(strMap mapSeq, length int) mapSeq {
	retMap := make(mapSeq)
	maxNum := length / 4

	for key, val := range strMap {
		leftNum := val - maxNum
		if leftNum > 0 {
			retMap[key] = leftNum
		}
	}
	return retMap
}

func getSeqs(strMap, baseMap mapSeq) []mapSeq {
	var ret []mapSeq
	fullSeq := "ACGT"
	for i := 1; i < len(fullSeq)+1; i++ {
		for j := 0; j < len(fullSeq)-i+1; j++ {
			if charInBase(strMap, baseMap, fullSeq[j:j+i]) {
				ret = append(ret, baseMap.update(strToMap(fullSeq[j:j+i])))
			}
		}
	}
	return ret
}

func charInBase(strMap, baseMap mapSeq, subSeq string) bool {
	subSeqMap := strToMap(subSeq)
	for key, _ := range subSeqMap {
		strValue := strMap[key]
		_, baseExists := baseMap[key]
		if strValue > 0 && baseExists {
			return false
		}
	}
	return true
}

func testSeqInStr(seq mapSeq, str string) bool {
	seqLen := seq.len()
	for i := 0; i < len(str)-seqLen+1; i++ {
		subSeq := strToMap(str[i : i+seqLen])
		if seq.compare(subSeq) {
			return true
		}
	}
	return false
}

// func solve() {
// 	var str string
// 	fmt.Scan(&str)
//
// 	if isPalindrom(str) {
// 		fmt.Println("-1")
// 	} else {
// 		index := leftCreatePalindrom(str)
// 		if index < 0 {
// 			index = rightCreatePalindrom(str)
// 		}
// 		fmt.Println(index)
// 	}
// }
//
// func isPalindrom(str string) bool {
// 	var middle1, middle2 int
// 	middle := len(str) / 2
// 	if len(str)%2 == 0 {
// 		middle1 = middle
// 		middle2 = middle1
// 	} else {
// 		middle1 = middle
// 		middle2 = middle + 1
// 	}
// 	return isStringsPalindrom(str[:middle1], str[middle2:])
// }
//
// func leftCreatePalindrom(str string) int {
// 	var middle1, middle2 int
// 	middle := len(str) / 2
// 	if len(str)%2 == 0 {
// 		middle1 = middle
// 		middle2 = middle + 1
// 	} else {
// 		middle1 = middle + 1
// 		middle2 = middle1
// 	}
// 	return findSpareLeft(str[:middle1], str[middle2:])
// }
//
// func rightCreatePalindrom(str string) int {
// 	var middle1, middle2 int
// 	middle := len(str) / 2
// 	if len(str)%2 == 0 {
// 		middle1 = middle - 1
// 		middle2 = middle
// 	} else {
// 		middle1 = middle
// 		middle2 = middle
// 	}
// 	result := findSpareRight(str[:middle1], str[middle2:])
// 	if result < 0 {
// 		return -1
// 	}
// 	return result + middle2
// }
//
// func isStringsPalindrom(str1, str2 string) bool {
// 	if len(str1) != len(str2) {
// 		return false
// 	}
// 	for i := 0; i < len(str1); i++ {
// 		if str1[i] != str2[len(str2)-1-i] {
// 			return false
// 		}
// 	}
// 	return true
// }
//
// func findSpareLeft(str1, str2 string) int {
// 	for i := 1; i < len(str1)+1; i++ {
// 		if isStringsPalindrom(str1[0:i-1]+str1[i:], str2) {
// 			return i - 1
// 		}
// 	}
// 	return -1
// }
// func findSpareRight(str1, str2 string) int {
// 	for i := 1; i < len(str2)+1; i++ {
// 		if isStringsPalindrom(str1, str2[0:i-1]+str2[i:]) {
// 			return i - 1
// 		}
// 	}
// 	return -1
// }

// type sortRunes []rune
//
// func (s sortRunes) Less(i, j int) bool {
// 	return s[i] < s[j]
// }
//
// func (s sortRunes) Swap(i, j int) {
// 	s[i], s[j] = s[j], s[i]
// }
//
// func (s sortRunes) Len() int {
// 	return len(s)
// }
//
// // SortString sort string
// func SortString(s string) string {
// 	r := []rune(s)
// 	sort.Sort(sortRunes(r))
// 	return string(r)
// }

// func solve() {
// 	var str string
// 	fmt.Scan(&str)
// 	store := make(map[string]int)
// 	for i := 1; i < len(str); i++ {
// 		for j := 0; j < len(str)-i+1; j++ {
// 			subStr := SortString(str[j : j+i])
// 			store[subStr]++
// 		}
// 	}
// 	var count int
// 	for key := range store {
// 		value := store[key]
// 		if value >= 2 {
// 			count += (value - 1) * value / 2
// 		}
// 	}
// 	fmt.Println(count)
// }

// func solve() {
// 	var str1, str2 string
// 	fmt.Scan(&str1, &str2)
// 	store1 := make(map[byte]int)
// 	store2 := make(map[byte]int)
//
// 	for i := 0; i < len(str1); i++ {
// 		store1[str1[i]]++
// 	}
// 	for i := 0; i < len(str2); i++ {
// 		store2[str2[i]]++
// 	}
// 	var minSubstring bool
// 	for key := range store1 {
// 		_, exists := store2[key]
// 		if exists {
// 			minSubstring = true
// 			break
// 		}
// 	}
// 	if minSubstring {
// 		fmt.Println("YES")
// 	} else {
// 		fmt.Println("NO")
// 	}
// }

// func max(args ...int) int {
// 	max := args[0]
// 	for i := 0; i < len(args); i++ {
// 		if args[i] > max {
// 			max = args[i]
// 		}
// 	}
// 	return max
// }

// func solve() {
// 	var str string
// 	fmt.Scan(&str)
// 	if len(str)%2 != 0 {
// 		fmt.Println(-1)
// 		return
// 	}
// 	l := len(str) / 2
// 	var str1, str2 string
// 	str1, str2 = string(str[:l]), string(str[l:])
// 	store1, store2 := make(map[byte]int), make(map[byte]int)
// 	for i := 0; i < len(str1); i++ {
// 		store1[str1[i]]++
// 	}
// 	for i := 0; i < len(str2); i++ {
// 		store2[str2[i]]++
// 	}
//
// 	var count int
// 	for key := range store2 {
// 		val, exists := store1[key]
// 		if !exists {
// 			count += store2[key]
// 		} else {
// 			delta := store2[key] - val
// 			if delta > 0 {
// 				count += delta
// 			}
// 		}
// 	}
// 	fmt.Println(count)
// }

// func main() {
// 	var str1, str2 string
// 	fmt.Scan(&str1, &str2)
// 	store1, store2 := make(map[byte]int), make(map[byte]int)
// 	for i := 0; i < len(str1); i++ {
// 		store1[str1[i]]++
// 	}
// 	for i := 0; i < len(str2); i++ {
// 		store2[str2[i]]++
// 	}
//
// 	var count int
// 	for key := range store1 {
// 		val, exists := store2[key]
// 		if !exists {
// 			count += store1[key]
// 		} else {
// 			count += int(math.Abs(float64(store1[key] - val)))
// 			delete(store2, key)
// 		}
// 	}
// 	for key := range store2 {
// 		count += store2[key]
// 	}
// 	fmt.Println(count)
// }

// func main() {
// 	var numOfRocks int
// 	fmt.Scan(&numOfRocks)
// 	store := make([]map[byte]int, numOfRocks)
// 	var str string
// 	for i := 0; i < numOfRocks; i++ {
// 		fmt.Scan(&str)
// 		store[i] = make(map[byte]int)
// 		for j := 0; j < len(str); j++ {
// 			store[i][str[j]]++
// 		}
// 	}
//
// 	var count int
// 	for key := range store[0] {
// 		isGem := true
// 		for i := 1; i < numOfRocks; i++ {
// 			_, exists := store[i][key]
// 			if !exists {
// 				isGem = false
// 				break
// 			}
// 		}
// 		if isGem {
// 			count++
// 		}
// 	}
//
// 	fmt.Println(count)
// }

// func main() {
// 	var str string
// 	fmt.Scan(&str)
// 	store := make(map[byte]int)
//
// 	for i := 0; i < len(str); i++ {
// 		store[str[i]]++
// 	}
//
// 	var oddCount int
// 	for key := range store {
// 		if store[key]%2 != 0 {
// 			oddCount++
// 		}
// 		if oddCount > 1 {
// 			break
// 		}
// 	}
// 	if oddCount > 1 {
// 		fmt.Println("NO")
// 	} else {
// 		fmt.Println("YES")
// 	}
// }

// func solve() {
// 	var str string
// 	fmt.Scan(&str)
// 	var count int
// 	for i := 1; i < len(str); i++ {
// 		if str[i] == str[i-1] {
// 			count++
// 		}
// 	}
// 	fmt.Println(count)
// }

// func solve() {
// 	var str string
// 	fmt.Scan(&str)
// 	funny := true
// 	for i := 1; i < len(str); i++ {
// 		if math.Abs(float64(str[i])-float64(str[i-1])) != math.Abs(float64(str[len(str)-1-i])-float64(str[len(str)-i])) {
// 			funny = false
// 			break
// 		}
// 	}
// 	if funny {
// 		fmt.Println("Funny")
// 	} else {
// 		fmt.Println("Not Funny")
// 	}
// }

// func main() {
// 	dict := map[byte]int{
// 		'a': 0,
// 		'b': 0,
// 		'c': 0,
// 		'd': 0,
// 		'e': 0,
// 		'f': 0,
// 		'g': 0,
// 		'h': 0,
// 		'i': 0,
// 		'j': 0,
// 		'k': 0,
// 		'l': 0,
// 		'm': 0,
// 		'n': 0,
// 		'o': 0,
// 		'p': 0,
// 		'q': 0,
// 		'r': 0,
// 		's': 0,
// 		't': 0,
// 		'u': 0,
// 		'v': 0,
// 		'w': 0,
// 		'x': 0,
// 		'y': 0,
// 		'z': 0,
// 	}
// 	reader := bufio.NewReader(os.Stdin)
// 	sentence, _ := reader.ReadString('\n')
// 	sentence = strings.ToLower(sentence)
// 	// count letters
// 	for i := 0; i < len(sentence); i++ {
// 		if sentence[i] != ' ' {
// 			dict[sentence[i]]++
// 		}
// 	}
// 	// Check on pangram
// 	pangram := true
// 	for key := range dict {
// 		if dict[key] == 0 {
// 			pangram = false
// 			break
// 		}
// 	}
// 	if pangram == true {
// 		fmt.Println("pangram")
// 	} else {
// 		fmt.Println("not pangram")
// 	}
// }

// // martix
// type matrixPoint struct {
// 	i int
// 	j int
// }
//
// type matrix struct {
// 	data    [][]int
// 	width   int
// 	height  int
// 	minSide int
// }
//
// func matrixNew(array2d *[][]int) *matrix {
// 	width := len((*array2d)[0])
// 	height := len(*array2d)
// 	minSide := width
// 	if minSide > height {
// 		minSide = height
// 	}
// 	return &matrix{*array2d, width, height, minSide}
// }
//
// func (m *matrix) calcCoord(k, depth int) *matrixPoint {
// 	var coord matrixPoint
// 	height, width := m.height-2*depth, m.width-2*depth
// 	if k < width {
// 		coord.i = depth
// 		coord.j = depth + k
// 	} else if k >= width && k < width+height-1 {
// 		coord.i = depth + k - width + 1
// 		coord.j = depth + width - 1
// 	} else if k >= width+height-1 && k < 2*width+height-2 {
// 		coord.i = depth + height - 1
// 		coord.j = depth + width - (k - (height + width - 2)) - 1
//
// 	} else if k >= 2*width+height-2 && k < 2*width+2*height-3 {
// 		coord.i = depth + height - (k - (2*width + height - 3)) - 1
// 		coord.j = depth
// 	}
//
// 	return &coord
// }
//
// func (m *matrix) String() string {
// 	var buffer bytes.Buffer
// 	for i := 0; i < m.height; i++ {
// 		for j := 0; j < m.width; j++ {
// 			if j != 0 {
// 				buffer.WriteString(" ")
// 			}
// 			buffer.WriteString(strconv.Itoa(m.data[i][j]))
// 		}
// 		buffer.WriteString("\n")
// 	}
// 	return buffer.String()
// }
//
// // Spiral
// type spiralsStorage map[int][]int
//
// func (spirals *spiralsStorage) shift(shiftIndex int) {
// 	for key := range *spirals {
// 		spirals.shiftSpiral(key, shiftIndex)
// 	}
// }
//
// func (spirals *spiralsStorage) shiftSpiral(depth, shiftIndex int) {
// 	spiral := (*spirals)[depth]
// 	shiftIndex = shiftIndex % len(spiral)
// 	buff := append(spiral[shiftIndex:], spiral[:shiftIndex]...)
// 	for i := 0; i < len(spiral); i++ {
// 		spiral[i] = buff[i]
// 	}
// }
//
// func (spirals *spiralsStorage) readFromMatrix(m *matrix) {
// 	numOfSpirals := m.minSide / 2
// 	var numOfElements int
// 	var coord *matrixPoint
// 	for i := 0; i < numOfSpirals; i++ {
// 		numOfElements = 2*(m.width-2*i+m.height-2*i) - 4
// 		(*spirals)[i] = make([]int, numOfElements)
// 		for j := 0; j < numOfElements; j++ {
// 			coord = m.calcCoord(j, i)
// 			(*spirals)[i][j] = m.data[coord.i][coord.j]
// 		}
// 	}
// }
//
// func (spirals *spiralsStorage) writeToMatrix(m *matrix) {
// 	var spiral []int
// 	var coord *matrixPoint
// 	for key := range *spirals {
// 		spiral = (*spirals)[key]
// 		for i := 0; i < len(spiral); i++ {
// 			coord = m.calcCoord(i, key)
// 			m.data[coord.i][coord.j] = spiral[i]
// 		}
// 	}
// }
//
// func main() {
// 	var M, N, R int
// 	fmt.Scan(&M, &N, &R)
// 	array2d := make([][]int, M)
// 	for i := 0; i < int(M); i++ {
// 		array2d[i] = make([]int, N)
// 		for j := 0; j < int(N); j++ {
// 			fmt.Scan(&array2d[i][j])
// 		}
// 	}
// 	fmt.Println(solve(&array2d, R))
// }
//
// func solve(array2d *[][]int, shiftIndex int) *matrix {
// 	m := matrixNew(array2d)
// 	spirals := make(spiralsStorage)
// 	spirals.readFromMatrix(m)
// 	spirals.shift(shiftIndex)
// 	spirals.writeToMatrix(m)
// 	return m
// }
//
// func convertSpirals(matrix [][]int) spiralsStorage {
// 	maxSide := len(matrix)
// 	if maxSide < len(matrix[0]) {
// 		maxSide = len(matrix[0])
// 	}
// 	numOfSpirals := maxSide / 2
// 	spirals := make(spiralsStorage)
// 	for i := 0; i < numOfSpirals; i++ {
// 		numOfElements := 2*(len(matrix[0])-2*i+len(matrix)-2*i) - 4
// 		spirals[i] = make([]int, numOfElements)
// 		for j := 0; j < numOfElements; j++ {
// 			coord := calcCoord(j, len(matrix), len(matrix[0]), i)
// 			spirals[i][j] = matrix[coord.i][coord.j]
// 		}
// 	}
//
// 	return spirals
// }
//
// func shift(array []int, shiftIndex int) {
// 	shiftIndex = shiftIndex % len(array)
// 	buff := append(array[shiftIndex:], array[:shiftIndex]...)
// 	for i := 0; i < len(array); i++ {
// 		array[i] = buff[i]
// 	}
// }
//
// func calcCoord(k, M, N int, depth int) matrixPoint {
// 	var coord matrixPoint
// 	m, n := M-2*depth, N-2*depth
// 	if k < n {
// 		coord.i = depth
// 		coord.j = depth + k
// 	} else if k >= n && k < n+m-1 {
// 		coord.i = depth + k - n + 1
// 		coord.j = depth + n - 1
// 	} else if k >= n+m-1 && k < n+2*m-2 {
// 		coord.i = depth + m - 1
// 		coord.j = depth + n - (k - (m + n - 2)) - 1
//
// 	} else if k >= n+2*m-2 && k < 2*n+2*m-3 {
// 		coord.i = depth + m - (k - (2*n + m - 3)) - 1
// 		coord.j = depth
// 	}
//
// 	return coord
// }

// GridSize - Grig size structure
// type GridSize struct {
// 	rows int
// 	cols int
// }

// func main() {
// 	var text string
// 	fmt.Scan(&text)
// 	textgridsize := calcGridSize(len(text))
//
// 	// Create store array
// 	store := make([][]byte, textgridsize.rows)
// 	for i := 0; i < textgridsize.rows; i++ {
// 		store[i] = make([]byte, textgridsize.cols)
// 	}
// 	// Populate store array
// 	var textCarriage int
// 	for i := 0; i < textgridsize.rows; i++ {
// 		for j := 0; j < textgridsize.cols; j++ {
// 			if textCarriage < len(text) {
// 				store[i][j] = text[textCarriage]
// 				textCarriage++
// 			}
// 		}
// 	}
// 	// Print it
// 	print(store)
// }
//
// func calcGridSize(textLen int) GridSize {
// 	var size GridSize
// 	// Calc rows and check if it's enought
// 	size.rows = int(math.Sqrt(float64(textLen)))
// 	if size.rows*size.rows == textLen {
// 		size.cols = size.rows
// 		return size
// 	}
// 	// Or assume cols size bigger on one unit
// 	size.cols = size.rows + 1
// 	// And if in's not enought add one unit to rows size
// 	if size.rows*size.cols < textLen {
// 		size.rows = size.rows + 1
// 	}
// 	return size
// }
//
// func print(array [][]byte) {
// 	for j := 0; j < len(array[0]); j++ {
// 		if j != 0 {
// 			fmt.Print(" ")
// 		}
// 		for i := 0; i < len(array); i++ {
// 			if array[i][j] != 0 {
// 				fmt.Printf("%c", array[i][j])
// 			}
// 		}
// 	}
// 	fmt.Print("\n")
// }

// func main() {
// 	var low, high int
// 	fmt.Scan(&low, &high)
// 	var kaprekarExists bool
// 	for i := low; i <= high; i++ {
// 		if isKaprekar(i) {
// 			fmt.Print(i, " ")
// 			kaprekarExists = true
// 		}
// 	}
// 	if !kaprekarExists {
// 		fmt.Println("INVALID RANGE")
// 	} else {
// 		fmt.Print("\n")
// 	}
//
// }
//
// func isKaprekar(number int) bool {
// 	quad := number * number
// 	quadStr := strconv.Itoa(quad)
// 	d := len(quadStr) / 2
// 	p1, _ := strconv.Atoi(string(quadStr[:d]))
// 	p2, _ := strconv.Atoi(string(quadStr[d:]))
// 	if p1+p2 == number {
// 		return true
// 	}
// 	return false
// }

// func main() {
// 	var hours, minutes int
// 	var res string
// 	fmt.Scan(&hours, &minutes)
// 	minutesStr := "minutes"
// 	if minutes%10 == 1 {
// 		minutesStr = "minute"
// 	}
//
// 	if minutes == 0 {
// 		res = getDigitStr(hours) + " o' clock"
// 	} else if minutes == 15 {
// 		res = "quarter past " + getDigitStr(hours)
// 	} else if minutes < 30 {
// 		res = getDigitStr(minutes) + " " + minutesStr + " past " + getDigitStr(hours)
// 	} else if minutes == 30 {
// 		res = "half past " + getDigitStr(hours)
// 	} else if minutes == 45 {
// 		res = "quarter to " + getDigitStr(hours+1)
// 	} else if minutes > 30 && minutes < 60 {
// 		res = getDigitStr(60-minutes) + " " + minutesStr + " to " + getDigitStr(hours+1)
// 	}
//
// 	fmt.Println(res)
// }

// func getDigitStr(digit int) string {
// 	digits := map[int]string{
// 		1:  "one",
// 		2:  "two",
// 		3:  "three",
// 		4:  "four",
// 		5:  "five",
// 		6:  "six",
// 		7:  "seven",
// 		8:  "eigth",
// 		9:  "nine",
// 		10: "ten",
// 		11: "eleven",
// 		12: "twelve",
// 		13: "thirteen",
// 		14: "fourteen",
// 		15: "fifteen",
// 		20: "twenty",
// 		30: "thirty",
// 		50: "fifty",
// 	}
// 	if digit <= 15 {
// 		return digits[digit]
// 	} else if digit < 20 {
// 		return digits[digit%10] + "teen"
// 	} else if digit/10 == 2 || digit/10 == 3 || digit/10 == 5 {
// 		if digit%10 == 0 {
// 			return digits[digit]
// 		}
// 		return digits[digit-digit%10] + " " + digits[digit%10]
// 	} else if digit%10 == 0 {
// 		return digits[digit/10] + "ty"
// 	}
//
// 	return digits[digit/10] + "ty" + " " + digits[digit%10]
// }

// func solve() {
// 	var Nb, Nw, Pb, Pw, Pch, res int64
// 	fmt.Scan(&Nb, &Nw, &Pb, &Pw, &Pch)
// 	if Pb > Pw+Pch {
// 		res = Pw*Nw + (Pw+Pch)*Nb
// 	} else if Pw > Pb+Pch {
// 		res = Pb*Nb + (Pb+Pch)*Nw
// 	} else {
// 		res = Pw*Nw + Pb*Nb
// 	}
//
// 	fmt.Println(res)
//
// }

// func main() {
// 	var numOfPeople, numOfSubjects int
// 	fmt.Scan(&numOfPeople, &numOfSubjects)
// 	// Parse input
// 	subjects := make([][]bool, numOfPeople)
// 	for i := 0; i < numOfPeople; i++ {
// 		var line string
// 		fmt.Scan(&line)
// 		subjects[i] = make([]bool, numOfSubjects)
// 		for j := 0; j < numOfSubjects; j++ {
// 			if line[j] == '0' {
// 				subjects[i][j] = false
// 			} else {
// 				subjects[i][j] = true
// 			}
// 		}
// 	}
//
// 	commStore := make(map[int]int)
//
// 	for i := 0; i < numOfPeople; i++ {
// 		for j := i + 1; j < numOfPeople; j++ {
// 			// count number of success subjects for group
// 			var count int
// 			for k := 0; k < numOfSubjects; k++ {
// 				if subjects[i][k] || subjects[j][k] {
// 					count++
// 				}
// 			}
// 			// store them in map
// 			_, inCommStore := commStore[count]
// 			if inCommStore {
// 				commStore[count]++
// 			} else {
// 				commStore[count] = 1
// 			}
// 		}
// 	}
//
// 	maxSubjects := findMaxKeyInMap(commStore)
// 	fmt.Println(maxSubjects)
// 	fmt.Println(commStore[maxSubjects])
// }
//
// func findMaxKeyInMap(mapToFind map[int]int) int {
// 	var max int
// 	for key := range mapToFind {
// 		if key > max {
// 			max = key
// 		}
// 	}
// 	return max
// }

// func calcFine(dateExpected, dateReal time.Time) int64 {
// 	if dateReal.Before(dateExpected) || dateReal.Equal(dateExpected) {
// 		return 0
// 	}
// 	nextMonth := dateExpected.AddDate(0, 1, -dateExpected.Day()+1)
// 	if dateReal.Before(nextMonth) {
// 		duration := dateReal.Day() - dateExpected.Day()
// 		return 15 * int64(duration)
// 	}
//
// 	nextYear := dateExpected.AddDate(1, -int(dateExpected.Month())+1, -dateExpected.Day()+1)
// 	if dateReal.Before(nextYear) {
// 		duration := int(dateReal.Month()) - int(dateExpected.Month())
// 		return 500 * int64(duration)
// 	}
// 	return 10000
// }

// func solve() {
// 	var n, a, b int
// 	fmt.Scan(&n, &a, &b)
// 	store := make([]int, n)
// 	for i := n - 1; i >= 0; i-- {
// 		store[i] = a*i + b*(n-i-1)
// 	}
//
// 	sort.Ints(store)
// 	for i := 0; i < len(store); i++ {
// 		fmt.Print(store[i], " ")
// 	}
// 	fmt.Println("")
// }

// func solve() {
// 	var R, C, r, c int
// 	fmt.Scan(&R, &C)
// 	big := scan2DArray(R, C)
// 	fmt.Scan(&r, &c)
// 	small := scan2DArray(r, c)
// 	res := find2DArray(big, small)
// 	if res == true {
// 		fmt.Println("YES")
// 	} else {
// 		fmt.Println("NO")
// 	}
// }
//
// func scan2DArray(r, c int) [][]int {
// 	array := make([][]int, r)
// 	for i := 0; i < r; i++ {
// 		array[i] = make([]int, c)
// 	}
//
// 	for i := 0; i < r; i++ {
// 		var line string
// 		fmt.Scan(&line)
// 		for j := 0; j < c; j++ {
// 			array[i][j] = int(line[j]) - 48
// 		}
//
// 	}
// 	return array
// }
//
// func find2DArray(big, small [][]int) bool {
// 	for i := 0; i < len(big); i++ {
// 		for j := 0; j < len(big[0]); j++ {
// 			if big[i][j] == small[0][0] && search(big, small, i, j) {
// 				return true
// 			}
// 		}
// 	}
// 	return false
// }
//
// func search(big, small [][]int, foundI, foundJ int) bool {
// 	if foundI+len(small) > len(big) || foundJ+len(small[0]) > len(big[0]) {
// 		return false
// 	}
//
// 	for i := 0; i < len(small); i++ {
// 		for j := 0; j < len(small[0]); j++ {
// 			if small[i][j] != big[i+foundI][j+foundJ] {
// 				return false
// 			}
// 		}
// 	}
// 	return true
// }

// func getCipherSymbol(symbol byte, shiftTo byte) byte {
// 	shiftTo = shiftTo % 26
// 	if symbol >= 'a' && symbol <= 'z' {
// 		if symbol+shiftTo <= 'z' {
// 			return symbol + shiftTo
// 		}
// 		return 'a' + (symbol + shiftTo - 'z' - 1)
// 	} else if symbol >= 'A' && symbol <= 'Z' {
// 		if symbol+shiftTo <= 'Z' {
// 			return symbol + shiftTo
// 		}
// 		return 'A' + (symbol + shiftTo - 'Z' - 1)
// 	}
// 	return symbol
// }

// func feast() {
// 	var N, C, M, total, totalExtra int
// 	fmt.Scan(&N, &C, &M)
// 	total = N / C
// 	for {
// 		extra := total / M
// 		if extra-totalExtra <= 0 {
// 			break
// 		}
// 		total += extra - totalExtra
// 		totalExtra = extra
// 	}
// 	fmt.Println(total)
// }

// func serviceLane(freeway []int) {
// 	var start, end int
// 	fmt.Scan(&start, &end)
// 	fmt.Println(minInArray(freeway, start, end))
// }
//
// func minInArray(array []int, start, end int) int {
// 	min := array[start]
// 	for i := start; i <= end; i++ {
// 		if min > array[i] {
// 			min = array[i]
// 		}
// 	}
// 	return min
// }

// func sherlockRoot() {
// 	var N, M int
// 	fmt.Scan(&N, &M)
//
// 	var count, start int
// 	for i := N; i < M+1; {
// 		digSqrt := math.Sqrt(float64(i))
// 		if digSqrt-math.Floor(digSqrt) == 0.0 {
// 			count++
// 			start = int(digSqrt) + 1
// 			break
// 		} else {
// 			i++
// 		}
//
// 	}
// 	if start > 0 {
// 		for i := start; i*i <= M; i++ {
// 			count++
// 		}
// 	}
// 	fmt.Println(count)
// }
//
// func findDigits() {
// 	var numStr string
// 	fmt.Scan(&numStr)
// 	num, _ := strconv.Atoi(numStr)
//
// 	var count int
// 	for j := 0; j < len(numStr); j++ {
// 		var dig = int(numStr[j] - 48)
// 		if dig > 0 && num%dig == 0 {
// 			count++
// 		}
// 	}
// 	fmt.Println(count)
// }
