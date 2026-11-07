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

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	if size <= 0 {
		return []int{}
	}

	randomElements := make([]int, size)
	src := rand.NewSource(time.Now().UnixNano())
	for i := range randomElements {
		randomElements[i] = int(src.Int63())
	}
	return randomElements
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}
	if len(data) == 1 {
		return data[0]
	}
	maxNumber := data[0]
	for _, v := range data {
		if v > maxNumber {
			maxNumber = v
		}
	}
	return maxNumber
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	if len(data) < CHUNKS {
		return maximum(data)
	}
	var wg sync.WaitGroup
	maxima := make([]int, CHUNKS)
	chunkSize := len(data) / CHUNKS
	wg.Add(CHUNKS)
	for i := 0; i < CHUNKS; i++ {
		chunk := data[i*chunkSize : (i+1)*chunkSize]
		if i == CHUNKS-1 {
			chunk = data[i*chunkSize:]
		}
		go func(ch []int, idx int) {
			defer wg.Done()
			maxima[idx] = maximum(ch)
		}(chunk, i)
	}
	wg.Wait()
	return maximum(maxima)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	data := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	start := time.Now()
	max := maximum(data)
	elapsed := time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d us\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	start = time.Now()
	max = maxChunks(data)
	elapsed = time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d us\n", max, elapsed)
}
