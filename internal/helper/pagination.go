package helper

import (
	"math"
	"strconv"
)

// AtoiOrDefault mengubah string menjadi int, jika gagal atau kosong maka return defaultValue
func AtoiOrDefault(s string, defaultValue int) int {
	if s == "" {
		return defaultValue
	}
	v, err := strconv.Atoi(s)
	if err != nil {
		return defaultValue
	}
	return v
}

// Max mengembalikan nilai terbesar dari dua int
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// CalculateTotalPages menghitung total halaman berdasarkan total data dan limit
func CalculateTotalPages(total int64, limit int) int {
	if total == 0 || limit == 0 {
		return 0
	}
	return int(math.Ceil(float64(total) / float64(limit)))
}
