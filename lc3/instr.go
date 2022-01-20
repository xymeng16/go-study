package main

func sign_extend(x uint16, bit_count int) uint16 {
	if (x >> (bit_count - 1)) & 1 {
		x |= 0xFFFF << bit_count
	}
	return x
}
