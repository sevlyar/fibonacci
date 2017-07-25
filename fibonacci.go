package fibonacci

import (
	"log"
	"math/big"
)

var debug = false

// Calc returns n-th Fibonacci number. It computes Fibonacci numbers in time O(log(n))
// and uses 'Rising a matrix to the power' method: (F0, F1) * P^n = (Fn, Fn+1).
// Calc panics if n is negative.
func Calc(n int64) *big.Int {
	if n < 0 {
		panic("Calc: index should be not negative")
	}

	var a, b, c, d big.Int
	var ta, tb, tc big.Int
	var rc, rd big.Int
	a.SetInt64(1)
	b.SetInt64(1)
	c.SetInt64(1)
	rd.SetInt64(1)

	for n != 0 {
		var tmp big.Int
		if (n & 1) != 0 {
			tc.Set(&rc)
			rc.Mul(&rc, &a)
			tmp.Mul(&rd, &c)
			rc.Add(&rc, &tmp)
			tmp.Mul(&rd, &d)
			rd.Mul(&tc, &b)
			rd.Add(&rd, &tmp)
			if debug {
				log.Printf("rc: %s, rd: %s", rc.String(), rd.String())
			}
		}
		ta.Set(&a)
		tb.Set(&b)
		tc.Set(&c)
		a.Add(a.Mul(&a, &a), tmp.Mul(&b, &c))
		b.Add(tmp.Mul(&ta, &b), b.Mul(&b, &d))
		c.Add(tmp.Mul(&ta, &c), c.Mul(&c, &d))
		d.Add(d.Mul(&d, &d), tmp.Mul(&tc, &tb))
		if debug {
			log.Printf("a: %s, b: %s", &a, &b)
			log.Printf("c: %s, d: %s", &c, &d)
		}
		n >>= 1
	}
	return &rc
}
