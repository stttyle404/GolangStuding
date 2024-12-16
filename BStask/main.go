//в Неевклидовом коридоре располагается 10^6 серверов
//При этом у каждого есть свой серийный номер, но он не указан на самом сервере
//Вас попросили найти сломанный сервер с номером 651953 и починить его .
//Поэтому вам выдали ноутбук с файлом , где вертикальное расположение серийных номеров серверов (они также располагаются в коридоре, только горизонтально)
//
// Ввод : номер нужного сервера target (1<=target<=1000000)
// Вывод : нужно вывести расположение сервера
//
// Ограничение памяти : 4mb
// Ограничение времени : 1 секунда
//
//
//

package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"sort"
	"strconv"
	"time"
)

// Функция поисков
func binarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
func LogicalLinealSearch(arr []int, target int) int {
	for index, value := range arr {
		if value == target {
			return index
		}
	}
	return -1
}

// Основная функция
func main() {
	originalArr := []int{}
	var target int
	fmt.Scan(&target)
	if target > 1000000 || target < 1 {
		panic("Target должен быть от 1 до 10^6 включительно")
	}
	startTime := time.Now() //нынешнее время
	f, err := os.Open("arr.txt")
	if err != nil {
		panic(err)
	}
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		s, _ := strconv.Atoi(sc.Text())
		originalArr = append(originalArr, s)
	}
	defer f.Close()

	go func() {
		memStats := &runtime.MemStats{}
		_ = make([]string, 1000000)
		runtime.ReadMemStats(memStats)
		fmt.Printf("\nАлоцированно = %v\n", memStats.Alloc/1024)
	}()

	arr := make([]int, len(originalArr))
	copy(arr, originalArr)
	sort.Ints(arr)

	duration := time.Since(startTime)
	indexInSorted := binarySearch(arr, target)
	if duration > time.Second {
		fmt.Println("Превышено время выполнения")
	} else {
		if indexInSorted != -1 {
			for i, v := range originalArr {
				if v == target && indexInSorted == sort.SearchInts(arr, v) {
					fmt.Printf("Сломанный сервер %d является %d ящик\n", target, i+1)
					break
				}
			}
		} else {
			fmt.Println("Сервер не найден")
		}
		fmt.Printf("Время выполнения: %v\n", duration)
	}
}
