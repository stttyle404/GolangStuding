/*
	Дейта - аналитику поручили задачу от разработчиков , так как он часто ругался с разрабами , в связи с тем
	что они бесконтрольно использовали его кофемашину
	Ему нужно найти сломанный сервер среди n серверов , тот сервер , который нужно отправить на техобслуживание
	Cостояние , которое сообщает об техобслуживание - (это число 5) , значения состояний совершенно неважен.
	При этом ему дали древний компьютер , который имеет ОЗУ 4мб , при этом ему обязательно нужно использовать всю память.
	На каждый сервер записано состояние его ядер (от 0 до 65536) ,состояние ядер расположенны хаотично, но как они представленны?

	На вход подается n , где n - максимально допустимое количество серверов.
	На выходе подается индекс найденного сервера.

	Ограничение памяти : 4мб.
	Ограничение скорости : 1с.

*/

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

func main() {
	startTime := time.Now() //нынешнее время

	f, err := os.Open("arr.txt")
	if err != nil {
		panic(err)
	}
	arr := []int{}
	defer f.Close()
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		s, _ := strconv.Atoi(sc.Text())
		arr = append(arr, s)
	}

	go func() {
		memStats := &runtime.MemStats{}
		_ = make([]string, 1000000)
		runtime.ReadMemStats(memStats)
		fmt.Printf("\nAlloc = %v\n", memStats.Alloc/1024)
	}()

	target := 858348
	clonearr := make([]int, len(arr))
	copy(arr, clonearr)
	sort.Ints(clonearr)
	index := binarySearch(clonearr, target)
	//index := LogicalLinealSearch(arr, target)

	duration := time.Since(startTime)
	if duration > time.Second {
		fmt.Println("Превышено время выполнения")
	} else {
		if index != -1 {
			// Находим индекс в исходном массиве
			// Индекс из отсортированного массива в исходном
			for i, v := range arr {
				if v == target && index == sort.SearchInts(arr, v) {
					fmt.Printf("Элемент %d найден на индексе %d в исходном массиве\n", target, i)
					break
				}
			}
		} else {
			fmt.Println("Элемент не найден")
		}
		fmt.Printf("Время выполнения: %v\n", duration)

	}
}
