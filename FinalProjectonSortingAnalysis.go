package main

import (
	"fmt"
	"math/rand" //random data
	"os"        //Dipakai untuk berhentiin program
	"time"      //Waktu
)

func main() {
	var (
		angkaRandom, n, i, j int
	)
	dataSet := [8]int{500, 1000, 10000, 20000, 100000, 200000, 1000000, 2000000}

	fmt.Println("--------------------------------------------------------------------------------------------------------------------")
	fmt.Printf("%8s%21s%21s%21s%21s%21s\n", "Ukuran Data", "Quick Sort(ms)", "Merge Sort(ms)", "Insertion Sort(ms)", "Selection Sort(ms)", "Bubble Sort(ms)")
	fmt.Println("--------------------------------------------------------------------------------------------------------------------")

	// Generate Data Random Start
	i = 0
	for i < len(dataSet) {
		n = dataSet[i]
		fmt.Printf("%8d", n)
		arrBubble := make([]int, n)
		arrInsertion := make([]int, n)
		arrMerge := make([]int, n)
		arrQuick := make([]int, n)
		arrSelection := make([]int, n)
		j = 0
		for j < n {
			angkaRandom = rand.Intn(n)
			arrBubble[j] = angkaRandom
			arrInsertion[j] = angkaRandom
			arrMerge[j] = angkaRandom
			arrQuick[j] = angkaRandom
			arrSelection[j] = angkaRandom
			j++
		}
		// Generate Data Random End

		hitungWaktu(arrQuick, "Quick Sort")
		hitungWaktu(arrMerge, "Merge Sort")
		hitungWaktu(arrInsertion, "Insertion Sort")
		hitungWaktu(arrSelection, "Selection Sort")
		hitungWaktu(arrBubble, "Bubble Sort")

		fmt.Println()
		i++
	}
}

func hitungWaktu(arr []int, namaSort string) {
	var (
		waktuMulai, waktuSelesai int64
	)
	waktuMulai = 0
	waktuSelesai = 0
	if namaSort == "Bubble Sort" {
		waktuMulai = time.Now().UnixNano() / 1000000
		bubbleSort(arr)
		waktuSelesai = time.Now().UnixNano() / 1000000
	} else if namaSort == "Insertion Sort" {
		waktuMulai = time.Now().UnixNano() / 1000000
		insertionSort(arr)
		waktuSelesai = time.Now().UnixNano() / 1000000
	} else if namaSort == "Merge Sort" {
		waktuMulai = time.Now().UnixNano() / 1000000
		mergeSort(arr)
		waktuSelesai = time.Now().UnixNano() / 1000000
	} else if namaSort == "Quick Sort" {
		waktuMulai = time.Now().UnixNano() / 1000000
		quickSort(arr)
		waktuSelesai = time.Now().UnixNano() / 1000000
	} else if namaSort == "Selection Sort" {
		waktuMulai = time.Now().UnixNano() / 1000000
		selectionSort(arr)
		waktuSelesai = time.Now().UnixNano() / 1000000
	}

	elapsed := waktuSelesai - waktuMulai
	fmt.Printf("%21d", elapsed)

	//10 menit = 600000 milisecond
	if elapsed > 600000 {
		fmt.Println("\n--------------------------------------------------------------------------------------------------------------------")
		fmt.Println("Ukuran data", len(arr), "untuk algoritma", namaSort, "mulai memakan waktu terlalu lama.")
		os.Exit(0)
	}
}

//Bubble Sort
func bubbleSort(data []int) {
	for i := 0; i < len(data); i++ {
		for j := 1; j < len(data)-i; j++ {
			if data[j] < data[j-1] {
				data[j], data[j-1] = data[j-1], data[j]
			}
		}
	}
}

//Insertion Sort
func insertionSort(data []int) {
	for i := 1; i < len(data); i++ {
		if data[i] < data[i-1] {
			j := i - 1
			temp := data[i]
			for j >= 0 && data[j] > temp {
				data[j+1] = data[j]
				j--
			}
			data[j+1] = temp
		}
	}
}

//Merge Sort
func mergeSort(data []int) []int {
	if len(data) <= 1 {
		return data
	}
	middle := len(data) / 2
	left := mergeSort(data[:middle])
	right := mergeSort(data[middle:])

	return merge(left, right)
}
func merge(left, right []int) []int {
	result := make([]int, len(left)+len(right))
	for i := 0; len(left) > 0 || len(right) > 0; i++ {
		if len(left) > 0 && len(right) > 0 {
			if left[0] < right[0] {
				result[i] = left[0]
				left = left[1:]
			} else {
				result[i] = right[0]
				right = right[1:]
			}
		} else if len(left) > 0 {
			result[i] = left[0]
			left = left[1:]
		} else if len(right) > 0 {
			result[i] = right[0]
			right = right[1:]
		}
	}
	return result
}

//Quick Sort
func quickSort(nums []int) {
	recursionSort(nums, 0, len(nums)-1)
}

func recursionSort(data []int, left int, right int) {
	if left < right {
		pivot := partition(data, left, right)
		recursionSort(data, left, pivot-1)
		recursionSort(data, pivot+1, right)
	}
}

func partition(data []int, left int, right int) int {
	for left < right {
		for left < right && data[left] <= data[right] {
			right--
		}
		if left < right {
			data[left], data[right] = data[right], data[left]
			left++
		}

		for left < right && data[left] <= data[right] {
			left++
		}
		if left < right {
			data[left], data[right] = data[right], data[left]
			right--
		}
	}
	return left
}

//Selection Sort
func selectionSort(data []int) {
	length := len(data)
	for i := 0; i < length; i++ {
		maxIndex := 0
		for j := 1; j < length-i; j++ {
			if data[j] > data[maxIndex] {
				maxIndex = j
			}
		}
		data[length-i-1], data[maxIndex] = data[maxIndex], data[length-i-1]
	}
}
