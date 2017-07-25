package fibonacci

import (
	"math/big"
	"testing"
)

func TestCalc(t *testing.T) {
	type testCase struct {
		N   int64
		Fib int64
	}
	var fixture = []testCase{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{7, 13},
		{8, 21},
		{9, 34},
		{10, 55},
	}
	for _, tc := range fixture {
		res := Calc(tc.N)
		if res.Cmp(big.NewInt(tc.Fib)) != 0 {
			t.Errorf("Invalid Fibonacci number for N = %d: %s (expected %d)", tc.N, res, tc.Fib)
		}
	}
}

func BenchmarkCalc_Speed(b *testing.B) {
	var res *big.Int
	res = Calc(int64(b.N))
	res.Sign()
}

func BenchmarkCalc_Mem(b *testing.B) {
	var res *big.Int
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		res = Calc(1000000)
	}
	res.Sign()
}
