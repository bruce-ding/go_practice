package test

// go test -bench='.' lock_test.go
import (
	"sync"
	"testing"
)

func Foo1() {
	i := 0
	i++
}

func Foo2() {
	lock := new(sync.Mutex)
	i := 0
	lock.Lock()
	i++
	lock.Unlock()
}

func BenchmarkFoo1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Foo1()
	}
}

func BenchmarkFoo2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Foo2()
	}
}

func BenchmarkFoo3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Foo2()
	}
}
