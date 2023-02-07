package main

import (
	"bufio"
	"encoding/json"
	"io"
	"os"
	"reflect"
	"strconv"
	"testing"
)

func TestAsteroidCollision(t *testing.T) {
	tests := []struct {
		name       string
		outputFile string
		asteroids  []int
		expected   []int
	}{
		{
			"0.1",
			"",
			[]int{5, 10, -5},
			[]int{5, 10},
		},
		{
			"0.2",
			"",
			[]int{8, -8},
			[]int{},
		},
		{
			"0.3",
			"",
			[]int{10, 2, -5},
			[]int{10},
		},
		{
			"0.4",
			"",
			[]int{-2, -1, 1, 2},
			[]int{-2, -1, 1, 2},
		},
		{
			"0.5",
			"",
			[]int{-2, 2, -1, -2},
			[]int{-2},
		},
		{
			"0.6",
			"",
			[]int{1, -1, 1, -2},
			[]int{-2},
		},
		{
			"0.6",
			"",
			[]int{1, 1, -1, -2},
			[]int{-2},
		},
		{
			"1",
			"",
			getAsteroidsFromJSONFile("./cases/case1.json"),
			[]int{5, 10},
		},
		{
			"2",
			"",
			getAsteroidsFromJSONFile("cases/case2.json"),
			[]int{},
		},
		{
			"3",
			"",
			getAsteroidsFromJSONFile("cases/case3.json"),
			[]int{10},
		},
		{
			"4",
			"",
			getAsteroidsFromTXTFile("cases/case1.txt"),
			[]int{5, 10},
		},
		{
			"5",
			"",
			getAsteroidsFromTXTFile("cases/case2.txt"),
			[]int{},
		},
		{
			"6",
			"",
			getAsteroidsFromTXTFile("cases/case3.txt"),
			[]int{10},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := asteroidCollision(tt.asteroids)
			if tt.outputFile != "" {
				saveResult(result, tt.outputFile)
			}
			if !reflect.DeepEqual(result, tt.expected) {
				t.Fatalf("Test case %s failed, expected: %v, got: %v", tt.name, tt.expected, result)
			}
		})
	}
}

func getAsteroidsFromJSONFile(file string) []int {
	f, err := os.Open(file)
	checkError(err)
	defer f.Close()

	data, err := io.ReadAll(f)
	checkError(err)

	var asteroids []int
	err = json.Unmarshal(data, &asteroids)
	checkError(err)

	return asteroids
}

func getAsteroidsFromTXTFile(file string) []int {
	f, err := os.Open(file)
	checkError(err)
	defer f.Close()

	var asteroids []int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		asteroid, err := strconv.Atoi(scanner.Text())
		checkError(err)

		asteroids = append(asteroids, asteroid)
	}

	return asteroids
}

func saveResult(result []int, file string) {
	out, err := os.Create(file)
	checkError(err)
	defer out.Close()

	data, err := json.Marshal(result)
	checkError(err)

	writer := bufio.NewWriter(out)
	writer.Write(data)
	//writer.WriteString("\n") // write new line
	//writer.Write(data) // write next line
	checkError(writer.Flush())
}
