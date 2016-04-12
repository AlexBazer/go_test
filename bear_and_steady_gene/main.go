package main

// https://www.hackerrank.com/challenges/bear-and-steady-gene
import (
	"fmt"
	"math"
)

type mapSeq map[byte]int

func newMapSeq(str string) mapSeq {
	// Create map sequence from string
	retMap := make(mapSeq)
	for i := 0; i < len(str); i++ {
		retMap[str[i]]++
	}
	return retMap
}

func (ms mapSeq) generateSparseSeq() mapSeq {
	// Generates sequence with sparce genome codes
	genomeCodes := []byte{'A', 'C', 'G', 'T'}
	retMap := make(mapSeq)
	maxNum := ms.len() / 4

	for i := 0; i < len(genomeCodes); i++ {
		value, exists := ms[genomeCodes[i]]
		if exists && value-maxNum > 0 {
			retMap[genomeCodes[i]] = value - maxNum
		} else {
			retMap[genomeCodes[i]] = 0
		}
	}
	return retMap
}

func (ms *mapSeq) copy() mapSeq {
	// Make a copy of sequence
	ret := make(mapSeq)
	for key, value := range *ms {
		ret[key] = value
	}
	return ret
}

func (ms mapSeq) copyTo(to mapSeq) {
	// Make a copy of sequence
	for key, value := range ms {
		to[key] = value
	}
}

func (ms *mapSeq) len() int {
	// Get sequence len
	var seqLen int
	for _, val := range *ms {
		seqLen += int(math.Abs(float64(val)))
	}
	return seqLen
}

func (ms mapSeq) checkSeq() bool {
	// Check if sequence does't have sparce genome codes
	for _, value := range ms {
		if value > 0 {
			return false
		}
	}
	return true
}

func (ms mapSeq) complement(with mapSeq) {
	// Complement sequence with enother sequence
	for key := range with {
		withValue, exists := with[key]
		if exists {
			ms[key] = withValue - ms[key]
		}
	}
}

func (ms mapSeq) complementSymbol(with byte) {
	// Complement sequence with symbol
	ms[with]--
}

func (ms mapSeq) unionSymbol(with byte) {
	// Union sequence with symbol
	ms[with]++
}

func main() {
	var numSimbols int
	var str string
	fmt.Scan(&numSimbols, &str)

	strMap := newMapSeq(str)
	// Find number of sparce symbols
	diffSeq := strMap.generateSparseSeq()
	if diffSeq.len() == 0 {
		// The sequence already is steady
		fmt.Println("0")
		return
	}
	// Create boat seq that will flow from start to the end of genome
	boatSeq := diffSeq.copy()
	var minLen int
	// find min from the start
	for i := 0; i < len(str)-1; i++ {
		boatSeq.complementSymbol(str[i])
		if boatSeq.checkSeq() {
			minLen = i + 1
			break
		}
	}

	var curI int
	// While knowing the minLen for now, we can:
	// 1.Delete one symbol from the tail of the boat and if seq is correct, decrement minLen
	// 2.Add one symbol to the head of the boat
	for curI+minLen < len(str) {
		boatSeq.unionSymbol(str[curI])
		if boatSeq.checkSeq() {
			minLen--
		} else {
			boatSeq.complementSymbol(str[curI+minLen])
		}
		curI++
	}
	fmt.Println(minLen)
}
