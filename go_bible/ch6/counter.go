package ch6

// Counter 封装n
type Counter struct {
	n int
}

// N count n的个数
func (c *Counter) N() int { return c.n }
