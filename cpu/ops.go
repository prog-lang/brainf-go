package cpu

func NOP(c *CPU) {}

func INC(c *CPU) {
	c.tape[c.tp]++
}

func DEC(c *CPU) {
	c.tape[c.tp]--
}

func NEXT(c *CPU) {
	c.tp++
}

func PREV(c *CPU) {
	c.tp--
}

func IN(c *CPU) {
	char, err := c.in.ReadByte()
	if err != nil {
		panic(err)
	}
	c.tape[c.tp] = char
}

func OUT(c *CPU) {
	if err := c.out.WriteByte(c.tape[c.tp]); err != nil {
		panic(err)
	}
}

func FWD(ip uint) Op {
	return func(c *CPU) {
		if c.tape[c.tp] == 0 {
			c.ip = ip
		}
	}
}

func BACK(ip uint) Op {
	return func(c *CPU) {
		if c.tape[c.tp] != 0 {
			c.ip = ip
		}
	}
}
