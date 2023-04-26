package test

import "testing"

type NumArray struct {
	nums []int
}

// 构造函数
func Constructor(nums []int) *NumArray {
	return &NumArray{nums: nums}
}

func (n *NumArray) SumRange(i int, j int) int {
	sum := 0
	for ; i <= j; i++ {
		sum += n.nums[i]
	}
	return sum
}

func (n *NumArray) Sort() *NumArray {
	for i := 0; i < len(n.nums); i++ {
		for j := i + 1; j < len(n.nums); j++ {
			if n.nums[i] > n.nums[j] {
				n.nums[i], n.nums[j] = n.nums[j], n.nums[i]
			}
		}
	}
	return n
}
func (n *NumArray) Insert(value int) *NumArray {
	n.nums = append(n.nums, value)
	return n
}

// 链式调用性能分析
func BenchmarkMethodChain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// 创建一个int
		Constructor([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}).Insert(0).Sort().Insert(1).SumRange(1, 5)
	}
}
func BenchmarkNoMethodChain(b *testing.B) {

	for i := 0; i < b.N; i++ {
		nums1 := Constructor([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
		nums2 := nums1.Insert(0)
		nums3 := nums2.Sort()
		nums4 := nums3.Insert(1)
		nums4.SumRange(1, 5)
		// 编译器更加容易优化不是吗? 也就是递归的性能更差不是吗?
	}
}
