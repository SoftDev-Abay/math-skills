package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]

	file, err := os.Open(args[0])

	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example

	var numbers []float64

	for scanner.Scan() {
		numberStr := scanner.Text()
		number, err := strconv.ParseFloat(numberStr, 32)
		if err != nil {
			fmt.Println(err)
			return
		}
		numbers = append(numbers, number)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return
	}

	average := math.Round(getAverage(numbers))
	median := math.Round(getMedian(numbers))
	variance := math.Round(getVarience(numbers))
	stdDeviation := math.Round(getStandDeviation(numbers))

	fmt.Println("Average:", average)
	fmt.Println("Median:", median)
	fmt.Println("Variance:", variance)
	fmt.Println("Standard Deviation:", stdDeviation)
}

func getAverage(nums []float64) float64 {
	total := 0.0
	for _, n := range nums {
		total += n
	}
	average := total / float64(len(nums))
	return average
}

func getMedian(nums []float64) float64 {
	nums = mergeSort(nums)

	lenNums := len(nums)
	if lenNums%2 == 1 {
		return nums[lenNums/2]
	}
	sumMiddleNums := nums[lenNums/2-1] + nums[lenNums/2]
	return sumMiddleNums / 2
}

func getVarience(nums []float64) float64 {
	avg := getAverage(nums)
	temp := 0.0

	for _, n := range nums {
		diffNAndAverage := n - avg
		temp += diffNAndAverage * diffNAndAverage
	}
	varience := temp / float64(len(nums))

	return varience
}

func getStandDeviation(nums []float64) float64 {
	varience := getVarience(nums)
	return math.Sqrt(varience)
}

func mergeSort(nums []float64) []float64 {
	lenNums := len(nums)
	if lenNums == 1 {
		return nums
	}
	first := mergeSort(nums[:lenNums/2])
	second := mergeSort(nums[lenNums/2:])

	return merge(first, second)
}

func merge(nums1 []float64, nums2 []float64) []float64 {
	i, j := 0, 0
	var result []float64
	lenNums1 := len(nums1)
	lenNums2 := len(nums2)
	for i < lenNums1 && j < lenNums2 {
		if nums1[i] <= nums2[j] {
			result = append(result, nums1[i])
			i++
		} else {
			result = append(result, nums2[j])
			j++
		}
	}
	for i < lenNums1 {
		result = append(result, nums1[i])
		i++
	}
	for j < lenNums2 {
		result = append(result, nums2[j])
		j++
	}
	return result
}
