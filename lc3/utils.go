package main

func sign_extend(x uint16, bit_count int) uint16 {
	if ((x >> (bit_count - 1)) & 1) == 1 {
		x |= 0xFFFF << bit_count
	}
	return x
}

func update_flags(r uint16) {
	if reg[r] == 0 {
		reg[R_COND] = FL_ZRO
	} else if ((reg[r] >> 15) & 1) == 1 {
		reg[R_COND] = FL_NEG
	} else {
		reg[R_COND] = FL_POS
	}
}

func mem_read(u uint16) uint16 {
	return u
}

func mem_write(addr uint16, val uint16) {

}
