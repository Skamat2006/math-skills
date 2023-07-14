package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"sort"
)

func readFile(filePath string) (numbers []int) {
	fd, err := os.Open(filePath)
	if err != nil {
		panic(fmt.Sprintf("open %s: %v", filePath, err))
	}
	var line int
	for {
		_, err := fmt.Fscanf(fd, "%d\n", &line)

		if err != nil {
			if err == io.EOF {
				return
			}
			panic(fmt.Sprintf("Scan Failed %s: %v", filePath, err))

		}
		numbers = append(numbers, line)
	}
	return
}

func average(numbers []int) float64 {
	array := numbers
	n := len(array)
	sum := 0
	for i := 0; i < n; i++ {
		sum += (array[i])
	}
	avg := (float64(sum)) / (float64(n))
	return avg
}

func median(numbers []int) float64 {
	sort.Ints(numbers)
	Median := 0.0
	size := len(numbers)
	M := size / 2

	// if size is odd, take the middle number
	if size%2 == 1 {
		Median = float64(numbers[M])
	} else {
		// if size is even, take the average of the middle two numbers
		Median = (float64(numbers[M-1]) + float64(numbers[M])) / 2
	}
	return Median
}

func variance(numbers []int) float64 {
	sumSqErrors := 0.0
	size := len(numbers)
	for _, number := range numbers {
		sumSqErrors += math.Pow(float64(number)-float64(average(numbers)), 2)
	}
	variance := sumSqErrors / float64(size)
	return variance
}

func main() {
	numbers := readFile("data.txt")
	fmt.Println("Average :", math.Round(average(numbers)))
	fmt.Println("Median :", math.Round(median(numbers)))
	fmt.Println("Variance:", math.Round(variance(numbers)))
	sd := math.Sqrt(variance(numbers))
	fmt.Println("Standard Deviation:", math.Round(sd))
}
