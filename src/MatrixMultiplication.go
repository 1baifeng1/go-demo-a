package main

import "fmt"

type Complex struct {
	r float32 // 实部
	i float32 // 虚部
}

type Triple struct {
	a0 Complex
	a1 Complex
	a2 Complex
}

type Mat struct {
	m, n int
	data [][]Triple
}

func (c Complex) CAdd(e Complex) Complex {
	r := Complex{}
	r.r = c.r + e.r
	r.i = c.i + e.i

	return r
}

func (c Complex) CMul(e Complex) Complex {
	r := Complex{}
	r.r = c.r*e.r - c.i*e.i
	r.i = c.i*e.r + c.r*e.i

	return r
}

func (t Triple) TAdd(e Triple) Triple {
	r := Triple{}
	r.a0 = t.a0.CAdd(e.a0)
	r.a1 = t.a1.CAdd(e.a1)
	r.a2 = t.a2.CAdd(e.a2)

	return r
}

func (t Triple) TMul(e Triple) Triple {
	r := Triple{}
	r.a0 = t.a0.CMul(e.a0)
	r.a1 = t.a1.CMul(e.a1)
	r.a2 = t.a2.CMul(e.a2)
	return r
}

// m的列数和e的行数相等
func (m Mat) MMul(e Mat) [][]Triple {
	res := [][]Triple{}
	for i := 0; i < m.n; i++ {
		t := []Triple{}
		for j := 0; j < e.m; j++ {
			r := Triple{}
			for k := 0; k < m.m; k++ {
				r = r.TAdd(m.data[i][k].TMul(e.data[k][j]))
			}
			t = append(t, r)
		}
		res = append(res, t)
	}
	return res
}

func main() {
	c1 := Complex{1, 1}
	c2 := Complex{1, 1}
	c3 := Complex{1, 1}
	c4 := Complex{1, 1}
	c5 := Complex{1, 1}
	c6 := Complex{1, 1}

	t1 := Triple{c1, c1, c1}
	t2 := Triple{c2, c2, c2}
	t3 := Triple{c3, c3, c3}
	t4 := Triple{c4, c4, c4}
	t5 := Triple{c5, c5, c5}
	t6 := Triple{c6, c6, c6}

	m1 := [][]Triple{
		{t1, t2},
		{t3, t4},
		{t5, t6},
	}
	m2 := [][]Triple{
		{t1, t2, t3},
		{t4, t5, t6},
	}

	A := Mat{
		2,
		3,
		m1,
	}
	B := Mat{
		3,
		2,
		m2,
	}

	res := A.MMul(B)
	fmt.Println(res)
}
