package popcount

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of bits set) of x.
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func PopCountShifting(x uint64) int {
	i := 0
	for x != 0 {
		if x&1 == 1 {
			i++
		}
		x >>= 1
	}
	return i
}

func PopCountBitop(x uint64) int {
	i := 0
	for x != 0 {
		i++
		x &= x - 1
	}
	return i
}
