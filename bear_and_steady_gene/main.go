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
	minLen := int(^uint(0) >> 1)

	diffSeq := strMap.generateSparseSeq()
	if diffSeq.len() == 0 {
		// The sequence already is steady
		fmt.Println("0")
		return
	}
	// Minimum possible seq length
	minPossibleLen := diffSeq.len()
	// Create firsth boat sequence
	// This sequence with minPossibleLen will be moving along seting sequence
	boatSeq := newMapSeq(str[0:minPossibleLen])
	for i := 0; i < len(str)-minPossibleLen; i++ {
		if i > 0 && i < len(str)-minPossibleLen-1 {
			boatSeq.complementSymbol(str[i-1])
			boatSeq.unionSymbol(str[i+minPossibleLen-1])
		}
		// Copy boat to trailSeq test against checkSeq
		trailSeq := boatSeq.copy()
		trailSeq.complement(diffSeq)
		if trailSeq.checkSeq() {
			// Minimal possible sequence is steady
			fmt.Println(minPossibleLen)
			return
		}
		// Widen trailSeq by one to len(str) or to possible len(j-i+1) lower minLen
		for j := i + diffSeq.len(); j < len(str) && j-i+1 < minLen; j++ {
			trailSeq.complementSymbol(str[j])
			if trailSeq.checkSeq() {
				curLen := j - i + 1
				if curLen < minLen {
					minLen = curLen
				}
				break
			}
		}
	}
	fmt.Println(minLen)
}
