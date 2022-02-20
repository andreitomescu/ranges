package ranges

import (
	"testing"
)

func assert(test *testing.T, cond bool, message ...string) {
	test.Log(cond, message)
	if cond == false {
		test.Fail()
	}
}

func done(test *testing.T) {
	// test.Fail()
}

func TestSlice(test *testing.T) {
	rng := Values(1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
	rng = Map(rng, func(x int) int { return x + 10 })
	rng = Concat(rng, Iota(10))
	rng = Sort(rng, Less[int])
	slc := ToSlice(rng)
	assert(test, Same(Slice(slc), Iota(20), Equal[int]))
	done(test)
}

func TestSame(test *testing.T) {
	rng1 := Values(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	rng2 := Iota(10)
	assert(test, Same(rng1, rng2, Equal[int]) == true)
	rng1 = Values(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	rng2 = Iota(11)
	assert(test, Same(rng1, rng2, Equal[int]) == false)
	rng1 = Values(0, 1, 2, 3, 4, 5, 6, 7, 8, 8)
	rng2 = Iota(10)
	assert(test, Same(rng1, rng2, Equal[int]) == false)
	done(test)
}

func TestRepeat(test *testing.T) {
	rng1 := Values(1, 1, 1, 1, 1)
	rng2 := Repeat(1, 5)
	assert(test, Same(rng1, rng2, Equal[int]))
	done(test)
}

func TestEmpty(test *testing.T) {
	assert(test, Same(Concat(Empty[int](), Iota(5), Empty[int]()), Iota(5), Equal[int]))
	done(test)
}

func TestSkip(test *testing.T) {
	assert(test, Same(Skip(Values(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 5), Values(5, 6, 7, 8, 9), Equal[int]))
	done(test)
}

func TestTake(test *testing.T) {
	assert(test, Same(Take(Values(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 5), Values(0, 1, 2, 3, 4), Equal[int]))
	done(test)
}

func TestGroup(test *testing.T) {
	rng := Dict(Group(Values(0, 0, 1, 1, 2, 2, 3, 3, 4, 4), Identity[int]))
	assert(test, Count(rng) == 5)
	assert(test, All(rng, func(x Pair[int, Range[int]]) bool { return Count(x.Value) == 2 }))
	assert(test, Same(Sort(Map(rng, PairKey[int, Range[int]]), Less[int]), Iota(5), Equal[int]))
	done(test)
}

func BenchmarkFilter(bench *testing.B) {
	for i := 0; i < bench.N; i++ {
		rng := Concat(Iota(bench.N*400), Iota(bench.N*400), Iota(bench.N*200))
		rng = Map(rng, func(x int) int { return x - 300 })
		rng = Filter(rng, func(x int) bool { return x < 300 })
		rng = Filter(rng, func(x int) bool { return x < 200 })
		rng = Filter(rng, func(x int) bool { return x < 100 })
		bench.Log(Count(rng), bench.N)
	}
}
