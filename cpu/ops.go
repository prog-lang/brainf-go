package cpu

func NOP(c *CPU) {}

func DEC(c *CPU) {
	c.tape[c.tp]--
}

func INC(c *CPU) {
	c.tape[c.tp]++
}

func PREV(c *CPU) {
	if c.tp > 0 {
		c.tp--
		return
	}
	newTape := make([]byte, len(c.tape), (len(c.tape)+1)*2)
	c.tp = uint(len(c.tape) - 1)
	c.tape = append(newTape, c.tape...)
}

func NEXT(c *CPU) {
	c.tp++
	if c.tp < uint(len(c.tape)) {
		return
	}
	newTape := make([]byte, (len(c.tape)+1)*2)
	copy(newTape, c.tape)
	c.tape = newTape
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
	if err := c.out.Flush(); err != nil {
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
