package byteop

func GetBit(b byte, n uint) int {return int(b >> n & 1)}

func GetBits(b byte, s uint, e uint) int {
	ret := 0
	for i := s; i >= e; i-- {
		ret = ret * 2 + GetBit(b, i)
	}
	return ret
}