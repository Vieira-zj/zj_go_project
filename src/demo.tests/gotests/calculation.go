package gotests

func calAdd(a, b int) int {
	return a + b
}

// MyCal struct for unit test.
type MyCal struct {
	base int
}

func (c *MyCal) addAndGet(num int) int {
	c.base += num
	return c.base
}

func (c *MyCal) selfAdd(num int) *MyCal {
	c.base += num
	return c
}

func (c *MyCal) selfDivide(num int) *MyCal {
	c.base -= num
	return c
}

func (c MyCal) getValue() int {
	return c.base
}

// NewMyCal creates new MyCal.
func NewMyCal(initNum int) *MyCal {
	return &MyCal{
		base: initNum,
	}
}
