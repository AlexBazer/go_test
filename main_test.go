package main

import (
	"os/exec"
	"testing"
)

func BenchmarkSolve(b *testing.B) {
	for n := 0; n < b.N; n++ {
		cmd := exec.Command("sh", "-c", "go run main.go < test.txt")
		cmd.Run()
	}
}

// func BenchmarkSolve(b *testing.B) {
// 	dat, err := ioutil.ReadFile("./test.txt")
// 	if err != nil {
// 		panic(err)
// 	}
// 	content := string(dat)
// 	lines := strings.Split(content, "\n")
// 	line1 := strings.Split(lines[0], " ")
// 	M, _ := strconv.Atoi(line1[0])
// 	N, _ := strconv.Atoi(line1[1])
// 	K, _ := strconv.Atoi(line1[2])
// 	array2d := make([][]int, M)
// 	var line []string
// 	for i := 0; i < M; i++ {
// 		line = strings.Split(lines[i+1], " ")
// 		array2d[i] = make([]int, N)
// 		for j := 0; j < N; j++ {
// 			array2d[i][j], _ = strconv.Atoi(line[j])
// 		}
// 	}
// 	for n := 0; n < b.N; n++ {
// 		solve(&array2d, K)
// 	}
// }

// TestCalcCoord test calc coord function
// func TestCalcCoord(t *testing.T) {
// 	// For 4X4 matrix
// 	depth := 0
// 	sets4 := [12]matrixPoint{
// 		matrixPoint{0, 0},
// 		matrixPoint{0, 1},
// 		matrixPoint{0, 2},
// 		matrixPoint{0, 3},
// 		matrixPoint{1, 3},
// 		matrixPoint{2, 3},
// 		matrixPoint{3, 3},
// 		matrixPoint{3, 2},
// 		matrixPoint{3, 1},
// 		matrixPoint{3, 0},
// 		matrixPoint{2, 0},
// 		matrixPoint{1, 0},
// 	}
// 	for i := 0; i < 12; i++ {
// 		if res := calcCoord(i, 4, 4, depth); res != sets4[i] {
// 			printCalcError(t, res, sets4[i])
// 		}
// 	}
//
// 	depth = 1
// 	sets2 := [4]matrixPoint{
// 		matrixPoint{1, 1},
// 		matrixPoint{1, 2},
// 		matrixPoint{2, 2},
// 		matrixPoint{2, 1},
// 	}
// 	for i := 0; i < 4; i++ {
// 		if res := calcCoord(i, 4, 4, depth); res != sets2[i] {
// 			printCalcError(t, res, sets2[i])
// 		}
// 	}
// 	// For 5X5 matrix
// 	depth = 0
// 	sets5 := [16]matrixPoint{
// 		matrixPoint{0, 0},
// 		matrixPoint{0, 1},
// 		matrixPoint{0, 2},
// 		matrixPoint{0, 3},
// 		matrixPoint{0, 4},
// 		matrixPoint{1, 4},
// 		matrixPoint{2, 4},
// 		matrixPoint{3, 4},
// 		matrixPoint{4, 4},
// 		matrixPoint{4, 3},
// 		matrixPoint{4, 2},
// 		matrixPoint{4, 1},
// 		matrixPoint{4, 0},
// 		matrixPoint{3, 0},
// 		matrixPoint{2, 0},
// 		matrixPoint{1, 0},
// 	}
// 	for i := 0; i < 16; i++ {
// 		if res := calcCoord(i, 5, 5, depth); res != sets5[i] {
// 			printCalcError(t, res, sets5[i])
// 		}
// 	}
// 	depth = 1
// 	sets54 := [8]matrixPoint{
// 		matrixPoint{1, 1},
// 		matrixPoint{1, 2},
// 		matrixPoint{1, 3},
// 		matrixPoint{2, 3},
// 		matrixPoint{3, 3},
// 		matrixPoint{3, 2},
// 		matrixPoint{3, 1},
// 		matrixPoint{2, 1},
// 	}
// 	for i := 0; i < 8; i++ {
// 		if res := calcCoord(i, 5, 5, depth); res != sets54[i] {
// 			printCalcError(t, res, sets54[i])
// 		}
// 	}
// 	// For 5X4 matrix
// 	array2d := [][]int{{1, 2, 3, 4, 5}, {1, 2, 3, 4, 5}, {1, 2, 3, 4, 5}, {1, 2, 3, 4, 5}}
// 	m := matrixNew(&array2d)
// 	depth = 0
// 	sets4x5 := [14]matrixPoint{
// 		matrixPoint{0, 0},
// 		matrixPoint{0, 1},
// 		matrixPoint{0, 2},
// 		matrixPoint{0, 3},
// 		matrixPoint{0, 4},
// 		matrixPoint{1, 4},
// 		matrixPoint{2, 4},
// 		matrixPoint{3, 4},
// 		matrixPoint{3, 3},
// 		matrixPoint{3, 2},
// 		matrixPoint{3, 1},
// 		matrixPoint{3, 0},
// 		matrixPoint{2, 0},
// 		matrixPoint{1, 0},
// 	}
// 	for i := 0; i < 14; i++ {
// 		if res := m.calcCoord(i, depth); *res != sets4x5[i] {
// 			printCalcError(t, *res, sets4x5[i])
// 		}
// 	}
// 	depth = 1
// 	sets5x44 := [6]matrixPoint{
// 		matrixPoint{1, 1},
// 		matrixPoint{1, 2},
// 		matrixPoint{1, 3},
// 		matrixPoint{2, 3},
// 		matrixPoint{2, 2},
// 		matrixPoint{2, 1},
// 	}
// 	for i := 0; i < 6; i++ {
// 		if res := calcCoord(i, 4, 5, depth); res != sets5x44[i] {
// 			printCalcError(t, res, sets5x44[i])
// 		}
// 	}
// 	// For 4x5 matrix
// 	array2d = [][]int{{1, 2, 3, 4}, {1, 2, 3, 4}, {1, 2, 3, 4}, {1, 2, 3, 4}, {1, 2, 3, 4}}
// 	m = matrixNew(&array2d)
// 	depth = 0
// 	sets5x4 := [14]matrixPoint{
// 		matrixPoint{0, 0},
// 		matrixPoint{0, 1},
// 		matrixPoint{0, 2},
// 		matrixPoint{0, 3},
// 		matrixPoint{1, 3},
// 		matrixPoint{2, 3},
// 		matrixPoint{3, 3},
// 		matrixPoint{4, 3},
// 		matrixPoint{4, 2},
// 		matrixPoint{4, 1},
// 		matrixPoint{4, 0},
// 		matrixPoint{3, 0},
// 		matrixPoint{2, 0},
// 		matrixPoint{1, 0},
// 	}
// 	for i := 0; i < 14; i++ {
// 		res := m.calcCoord(i, depth)
// 		if *res != sets5x4[i] {
// 			printCalcError(t, *res, sets5x4[i])
// 		}
// 	}
//
// }

// func TestArrayShift(t *testing.T) {
// 	spirals := spiralsStorage{
// 		0: {1, 2, 3, 4, 5, 6, 7},
// 		1: {1, 2, 3, 4, 5, 6, 7},
// 	}
// 	resSpirals := spiralsStorage{
// 		0: {4, 5, 6, 7, 1, 2, 3},
// 		1: {1, 2, 3, 4, 5, 6, 7},
// 	}
// 	spirals.shiftSpiral(0, 3)
//
// 	if !reflect.DeepEqual(spirals, resSpirals) {
// 		t.Errorf("res %v != set %v", resSpirals, spirals)
// 	}
//
// 	spirals = spiralsStorage{
// 		0: {1, 2, 3, 4, 5, 6, 7},
// 		1: {1, 2, 3, 4, 5, 6, 7},
// 	}
// 	resSpirals = spiralsStorage{
// 		0: {1, 2, 3, 4, 5, 6, 7},
// 		1: {4, 5, 6, 7, 1, 2, 3},
// 	}
// 	spirals.shiftSpiral(1, 10)
//
// 	if !reflect.DeepEqual(spirals, resSpirals) {
// 		t.Errorf("res %v != set %v", resSpirals, spirals)
// 	}
//
// 	spirals = spiralsStorage{
// 		0: {1, 2, 3, 4, 5, 6, 7},
// 		1: {10, 20, 30, 40, 50},
// 	}
// 	resSpirals = spiralsStorage{
// 		0: {4, 5, 6, 7, 1, 2, 3},
// 		1: {40, 50, 10, 20, 30},
// 	}
// 	spirals.shift(3)
// 	if !reflect.DeepEqual(spirals, resSpirals) {
// 		t.Errorf("res %v != set %v", resSpirals, spirals)
// 	}
//
// 	spirals = spiralsStorage{
// 		0: {1, 2, 3, 4, 5, 6, 7},
// 		1: {10, 20, 30, 40, 50},
// 	}
// 	resSpirals = spiralsStorage{
// 		0: {7, 1, 2, 3, 4, 5, 6},
// 		1: {20, 30, 40, 50, 10},
// 	}
// 	spirals.shift(6)
// 	if !reflect.DeepEqual(spirals, resSpirals) {
// 		t.Errorf("res %v != set %v", resSpirals, spirals)
// 	}

// }

func printCalcError(t *testing.T, res, set matrixPoint) {
	t.Errorf("res %v  !=   set %v", res, set)
}
