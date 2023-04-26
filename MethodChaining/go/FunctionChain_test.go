package test

import (
	"github.com/KM911/oslib"
	"testing"
)

func BenchmarkChaining(b *testing.B) {
	for i := 0; i < b.N; i++ {
		oslib.IsSameArray(oslib.Dir(oslib.CmdPath()), oslib.Dir(oslib.ExecutePath()))
	}
}

func BenchmarkNoChaining(b *testing.B) {
	for i := 0; i < b.N; i++ {
		path1 := oslib.CmdPath()
		list1 := oslib.Dir(path1)
		path2 := oslib.ExecutePath()
		list2 := oslib.Dir(path2)
		oslib.IsSameArray(list1, list2)
	}
}
