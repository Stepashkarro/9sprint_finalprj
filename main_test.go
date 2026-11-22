package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerateRandomElements(t *testing.T) {
	// Проверка на нулевой размер
	empty := generateRandomElements(0)
	assert.Empty(t, empty)

	// Проверка на отрицательный размер
	negative := generateRandomElements(-5)
	assert.Empty(t, negative)

	// Проверка нормального размера
	size := 100
	data := generateRandomElements(size)
	require.Len(t, data, size)

	// Проверяем, что числа неотрицательные
	for _, v := range data {
		assert.GreaterOrEqual(t, v, 0)
	}
}

func TestMaximum(t *testing.T) {
	// Проверка пустого слайса
	assert.Equal(t, 0, maximum([]int{}))

	// Проверка слайса с одним элементом
	assert.Equal(t, 42, maximum([]int{42}))

	// Проверка обычного слайса
	data := []int{1, 5, 3, 9, 2, 8}
	assert.Equal(t, 9, maximum(data))

	// Проверка слайса, где максимум в начале
	assert.Equal(t, 100, maximum([]int{100, 5, 3, 9, 2, 8}))

	// Проверка слайса, где максимум в конце
	assert.Equal(t, 100, maximum([]int{1, 5, 3, 9, 2, 100}))
}

// Тестируем и многопоточную функцию, чтобы убедиться, что она дает тот же результат
func TestMaxChunks(t *testing.T) {
	assert.Equal(t, 0, maxChunks([]int{}))
	assert.Equal(t, 42, maxChunks([]int{42}))

	// Слайс, длина которого не кратна CHUNKS (8), чтобы проверить обработку остатка
	data := []int{1, 5, 3, 9, 2, 8, 7, 6, 100, 4}
	assert.Equal(t, 100, maxChunks(data))
}
