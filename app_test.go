package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 파일명이 _test.go 로 이뤄져야 한다
// testing 패키지를 임포트해야 한다
// 테스트 코드는 func TestXxxxx(t *testing T) 형태이다

func TestSquare(t *testing.T) {
	assertions := assert.New(t)
	assertions.Equal(81, Square(9), "square(9) should be 81")
}

func TestSquare_second(t *testing.T) {
	assertions := assert.New(t)
	assertions.Equal(36, Square(6), "square(6) should be 36")
}

func TestSquare_third(t *testing.T) {
	result := Square(3)
	if result != 9 {
		t.Errorf("square(9) should be 9 but returns %d", result)
	}
}

func BenchmarkFibo1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fibo(5)
	}
}

func BenchmarkFibo2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FiboRecursive(5)
	}
}

// 깊이 우선 탐색
const (
	n = 6
)

var graph = [][]int{
	{},
	{2, 6, 5},
	{1},
	{4, 5, 6},
	{3},
	{3},
	{1, 3},
}

func BenchmarkDFS(b *testing.B) {
	for i := 0; i < b.N; i++ {
		visited := make([]bool, n+1)
		Dfs1(1, visited, graph)
	}
}

func BenchmarkDFSRecursive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		visited := make([]bool, n+1)
		Dfs2(1, visited, graph)
	}
}

// test table
func TestSquare_fourth(t *testing.T) {

	type test struct {
		data   int
		answer int
	}

	tests := []test{
		{5, 25},
		{7, 49},
		{8, 64},
		{10, 100},
		{2, 4},
	}

	assertions := assert.New(t)
	for _, test := range tests {
		assertions.Equalf(test.answer, Square(test.data),
			"square(%d) should be %d", test.data, test.answer)
	}
}
