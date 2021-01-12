package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Implementation of Bloom filter")
	var falsePositiveProbability float64
	var dataCount int64
	falsePositiveProbability = .1
	dataCount = 10
	var filterArraySize = getSize(dataCount, falsePositiveProbability)
	var hashCount = getHashCount(filterArraySize, dataCount)
	filterArray := make([]bool, filterArraySize)
	fmt.Println("filterArraySize:", filterArraySize)
	fmt.Println("hashCount:", hashCount)
	dataSet := []string{"golang", "codebook", "godoc", "go", "cat", "dog", "cow", "goat", "day", "night"}
	for _, v := range dataSet {
		if Lookup(v, filterArray, hashCount, filterArraySize) == false {
			Insert(v, filterArray, hashCount, filterArraySize)
		}
	}
	fmt.Println(filterArray)
}

func Lookup(data string, array []bool, hashCount int, size int64) bool {

	for i := 0; i < hashCount; i++ {
		hash := Hash(data, size, i)
		if array[hash] == false {
			return false
		}
	}
	return true
}

func Insert(data string, array []bool, hashCount int, size int64) {
	for i := 0; i < hashCount; i++ {
		hash := Hash(data, size, i)
		array[hash] = true
	}
}

func Hash(data string, size int64, seed int) int64 {
	var hash int64
	hash = int64(seed)
	for i, s := range data {
		hash = hash + (hash*19+int64(math.Pow(7, float64(i)))*int64(s))%size
		hash = hash % size
	}
	return hash % size
}

func getSize(dataCount int64, fpp float64) int64 {
	var size float64
	size = -(math.Log(fpp) * float64(dataCount)) / (math.Pow(math.Log(2), 2.0))
	return int64(size)
}

func getHashCount(filterArraySize int64, dataCount int64) int {
	var hashCount = float64(int64(filterArraySize)/dataCount) * math.Log(2.0)
	return int(hashCount)
}
