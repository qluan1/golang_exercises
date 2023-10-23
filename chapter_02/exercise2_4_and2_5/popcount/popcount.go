package popcount

func PopCount1(x uint64) int {
	count := 0
	for i:= 0; i < 64; i++ {
		count += int(byte(x>>(i)&1))
	}
	return count
}

func PopCount2(x uint64) int {
	count := 0
	for x != 0 {
		x = x & (x-1)
		count++
	}
	return count
}