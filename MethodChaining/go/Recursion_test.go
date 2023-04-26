package test

import "testing"

func BenchmarkNORecursion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		target := 100
		sum := 0
		for i := 0; i < target; i++ {
			sum += i
		}
	}
}
func SUM(target int) int {
	if target == 1 {
		return 1
	} else {
		return target + SUM(target-1)
	}
}

func BenchmarkRecursion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SUM(100)
	}
}

func BenchmarkTailRecursion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TailSUM(100, 0)
	}
}

func TailSUM(target int, sum int) int {
	if target == 1 {
		return sum + 1
	} else {
		return TailSUM(target-1, sum+target)
	}
}

// 线性递归
func LinerSum(a int, b int, target int) int {
	if a == target {
		return b
	}
	return LinerSum(a+1, b+a, target)
}

func BenchmarkLinerRecursion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LinerSum(1, 0, 100)
	}
}
