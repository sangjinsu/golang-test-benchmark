package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 파일명이 _test.go 로 이뤄져야 한다
// testing 패키지를 임포트해야 한다
// 테스트 코드는 func TestXxxxx(t *testing T) 형태이다

func TestSquare1(t *testing.T) {
	assertions := assert.New(t)
	assertions.Equal(81, square(9), "square(9) should be 81")
}

func TestSquare2(t *testing.T) {
	assertions := assert.New(t)
	assertions.Equal(36, square(6), "square(6) should be 36")
}

func TestSquare3(t *testing.T) {
	result := square(3)
	if result != 9 {
		t.Errorf("square(9) should be 9 but returns %d", result)
	}
}

func BenchmarkFibo1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibo(5)
	}
}

func BenchmarkFibo2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fiboRecursive(5)
	}
}
