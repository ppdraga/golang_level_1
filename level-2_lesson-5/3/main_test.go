package main

import (
	"math/rand"
	"sync"
	"testing"
)

type FloatSet struct {
	sync.Mutex
	mm map[float64]struct{}
}

func NewFloatSet() *FloatSet {
	return &FloatSet{
		mm: map[float64]struct{}{},
	}
}

func (s *FloatSet) Add(i float64) {
	s.Lock()
	s.mm[i] = struct{}{}
	s.Unlock()
}

func (s *FloatSet) Has(i float64) bool {
	s.Lock()
	defer s.Unlock()
	_, ok := s.mm[i]
	return ok
}

func BenchmarkFloatSet10w90r(b *testing.B) {
	var set = NewFloatSet()
	b.Run("", func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				if rand.Intn(10) == 1 {
					set.Add(1.55)
				} else {
					set.Has(1.55)
				}
			}
		})
	})
}

func BenchmarkFloatSet50w50r(b *testing.B) {
	var set = NewFloatSet()
	b.Run("", func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				if rand.Intn(2) == 1 {
					set.Add(1.55)
				} else {
					set.Has(1.55)
				}
			}
		})
	})
}

func BenchmarkFloatSet90w10r(b *testing.B) {
	var set = NewFloatSet()
	b.Run("", func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				if rand.Intn(10) == 1 {
					set.Has(1.55)
				} else {
					set.Add(1.55)
				}
			}
		})
	})
}

type FloatSetRW struct {
	sync.RWMutex
	mm map[float64]struct{}
}

func NewFloatSetRW() *FloatSetRW {
	return &FloatSetRW{
		mm: map[float64]struct{}{},
	}
}

func (s *FloatSetRW) Add(i float64) {
	s.Lock()
	s.mm[i] = struct{}{}
	s.Unlock()
}

func (s *FloatSetRW) Has(i float64) bool {
	s.RLock()
	defer s.RUnlock()
	_, ok := s.mm[i]
	return ok
}

func BenchmarkFloatSetRW10w90r(b *testing.B) {
	var set = NewFloatSetRW()
	b.Run("", func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				if rand.Intn(10) == 1 {
					set.Add(1.55)
				} else {
					set.Has(1.55)
				}
			}
		})
	})
}

func BenchmarkFloatSetRW50w50r(b *testing.B) {
	var set = NewFloatSetRW()
	b.Run("", func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				if rand.Intn(2) == 1 {
					set.Add(1.55)
				} else {
					set.Has(1.55)
				}
			}
		})
	})
}

func BenchmarkFloatSetRW90w10r(b *testing.B) {
	var set = NewFloatSetRW()
	b.Run("", func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				if rand.Intn(10) == 1 {
					set.Has(1.55)
				} else {
					set.Add(1.55)
				}
			}
		})
	})
}

func main() {

}
