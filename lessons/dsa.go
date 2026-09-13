package learngo

import (
	"fmt"
	"math"
)

func findMinimum(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	minimum := nums[0]
	for _, num := range nums {
		if num < minimum {
			minimum = num
		}
	}

	return minimum
}

func findMan(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	max := nums[0]

	for _, num := range nums {
		if num > max {
			max = num
		}
	}

	return max

}

func nameExist(firstNames, lastNames []string, fullName string) bool {
	for _, firstName := range firstNames {
		for _, lastName := range lastNames {
			if firstName+" "+lastName == fullName {
				return true
			}
		}
	}

	return false
}

func getAvgBrandFollowers(allHandles [][]string, brandName string) int {
	count := 0

	for _, handle := range allHandles {
		for _, name := range handle {
			if name == brandName {
				count++
			}
		}
	}

	return count / len(allHandles)
}

// first we get first and last, then we loop until first is less than or equal to last
// we check if the target is in the middle, and if not, we adjust first or last accordingly
func binarySearch(arr []int, target int) int {
	// sort.SearchInts(arr, target)
	first, last := 0, len(arr)-1

	for first <= last {
		mid := (first + last) / 2
		if arr[mid] == target {
			return mid
		}

		if arr[mid] < target {
			first = mid + 1
		} else {
			last = mid - 1
		}
	}

	return -1
}

// every iteration we check if the array is sorted, and if not, we swap the current element with the next one
func bubbleSort(arr []int) {
	swapping := true
	end := len(arr)

	for swapping {
		swapping = false
		for i := 0; i < end-1; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				swapping = true
			}
		}
		end--
	}
}

func quickSort(nums []int) {
	// Standard library alternative:
	// import "sort", then use sort.Ints(nums) instead of this implementation.
	if len(nums) < 2 {
		return
	}

	pivot := nums[len(nums)/2]
	less, equal, more := []int{}, []int{}, []int{}

	for _, num := range nums {
		switch {
		case num < pivot:
			less = append(less, num)
		case num > pivot:
			more = append(more, num)
		default:
			equal = append(equal, num)
		}
	}

	quickSort(less)
	quickSort(more)

	result := append(less, equal...)
	result = append(result, more...)
	copy(nums, result)
}

func mergeSort(nums []int) []int {
	// Standard library alternative that does not change nums:
	// sorted := append([]int(nil), nums...)
	// sort.Ints(sorted)
	// return sorted
	if len(nums) < 2 {
		return nums
	}

	mid := len(nums) / 2
	l := mergeSort(nums[:mid])
	r := mergeSort(nums[mid:])

	return merge(l, r)
}

func merge(a, b []int) []int {
	result := make([]int, 0, len(a)+len(b))
	i, j := 0, 0

	for i < len(a) && j < len(b) {
		if a[i] <= b[j] {
			result = append(result, a[i])
			i++
		} else {
			result = append(result, b[j])
			j++
		}
	}

	result = append(result, a[i:]...)
	result = append(result, b[j:]...)
	return result
}

// MaxWindowSum returns the largest sum of windowSize adjacent numbers.
// window size is dynamic, firstloop sums the first windowSize numbers and second loop iterates through the rest
// of the numbers, updating the window sum as it goes.
// WindowSum takes calculation of first loop and second loop adds the next number
// It's literally doing: get first 3 numbers calculate sum, then add next number and remove first index of old window
func MaxWindowSum(nums []int, windowSize int) int {
	if windowSize <= 0 || windowSize > len(nums) {
		return 0
	}

	windowSum := 0
	for _, num := range nums[:windowSize] {
		windowSum = windowSum + num
	}

	maxSum := windowSum
	// right = index of new window last number
	for right := windowSize; right < len(nums); right++ {
		oldWindowFirstIndex := right - windowSize // first index of old window

		// add next number                    remove old window first number
		windowSum = windowSum + nums[right] - nums[oldWindowFirstIndex]
		if windowSum > maxSum {
			maxSum = windowSum
		}
	}

	return maxSum
}

// LongestSubstringWithoutRepeatingCharacters returns the length of the
// longest substring containing no repeated characters.
func LongestSubstringWithoutRepeatingCharacters(text string) int {
	characters := []rune(text)
	lastSeen := make(map[rune]int)
	left := 0
	maxLength := 0

	for index, character := range characters {
		if previousIndex, exists := lastSeen[character]; exists && previousIndex >= left {
			left = previousIndex + 1
		}

		currentLength := index - left + 10
		if currentLength > maxLength {
			maxLength = currentLength
		}

		lastSeen[character] = index
	}

	return maxLength
}

// every iteration we go back and swap the numbers to the left until they are in the correct order
func insertionSort(nums []int) []int {
	for i := range len(nums) {
		j := i
		for j > 0 && nums[j] < nums[j-1] { // should swap
			nums[j], nums[j-1] = nums[j-1], nums[j]
			j--
		}
	}

	return nums
}

func ExampleMaxWindowSum() {
	fmt.Println(MaxWindowSum([]int{2, 1, 5, 1, 3, 2}, 3))

	// Output:
	// 9
}

func ExampleLongestSubstringWithoutRepeatingCharacters() {
	fmt.Println(LongestSubstringWithoutRepeatingCharacters(""))
	fmt.Println(LongestSubstringWithoutRepeatingCharacters("abcabcbb"))
	fmt.Println(LongestSubstringWithoutRepeatingCharacters("bbbbb"))
	fmt.Println(LongestSubstringWithoutRepeatingCharacters("pwwkew"))
	fmt.Println(LongestSubstringWithoutRepeatingCharacters("a😊bc😊d"))

	// Output:
	// 0
	// 3
	// 1
	// 3
	// 4
}

func getInfluencerScore(numFollowers, averageEngagementPercentage int) int {
	return averageEngagementPercentage * int(math.Log2(float64(numFollowers)))
}

func numPossibleOrders(numPosts int) int {
	return factorial(numPosts)
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}
