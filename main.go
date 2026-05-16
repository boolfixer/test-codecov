package main

type Calculator struct{}

func (c *Calculator) Sum(a, b int) int {
	return a + b
}

func (c *Calculator) Sub(a, b int) int {
	return a - b
}

func (c *Calculator) Mul(a, b int) int {
	return a * b
}

func main() {
	c := &Calculator{}
	print(c.Sum(1, 2))
}
