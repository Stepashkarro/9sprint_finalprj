package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

func generateRandomElements(size int) []int {
	// крайний случай: размер 0 или меньше
	if size <= 0 {
		return []int{}
	}

	data := make([]int, size)
	for i := range data {
		data[i] = rand.Int() // генераация случайного положительного числа
	}
	return data
}

func maximum(data []int) int {
	// крайний случай: пустой слайс
	if len(data) == 0 {
		return 0
	}

	max := data[0]
	for _, v := range data[1:] {
		if v > max {
			max = v
		}
	}
	return max
}

func maxChunks(data []int) int {
	// крайний случай: пустой слайс
	if len(data) == 0 {
		return 0
	}

	var wg sync.WaitGroup
	chunkSize := len(data) / CHUNKS
	maxes := make([]int, CHUNKS)

	for i := 0; i < CHUNKS; i++ {
		start := i * chunkSize
		end := start + chunkSize

		// если это последний кусок, забираем остаток элементов (если длина не делится на 8 нацело)
		if i == CHUNKS-1 {
			end = len(data)
		}

		wg.Add(1)
		// запуск горутины
		go func(idx int, slice []int) {
			defer wg.Done()
			maxes[idx] = maximum(slice)
		}(i, data[start:end])
	}

	wg.Wait() // ждём завершения всех 8 горутин

	// Нмаксимум среди 8 максимумов
	return maximum(maxes)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	data := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	start := time.Now()
	maxSingle := maximum(data)
	elapsedSingle := time.Since(start).Microseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d µs\n", maxSingle, elapsedSingle)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	start = time.Now()
	maxMulti := maxChunks(data)
	elapsedMulti := time.Since(start).Microseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d µs\n", maxMulti, elapsedMulti)
}
