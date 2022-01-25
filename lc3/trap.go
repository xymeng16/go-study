package main

/*
#include <stdio.h>
*/
import "C"
import (
	"bufio"
)

const (
	TRAP_GETC  uint8 = 0x20 /* get character from keyboard, not echoed onto the terminal */
	TRAP_OUT         = 0x21 /* output a character */
	TRAP_PUTS        = 0x22 /* output a word string */
	TRAP_IN          = 0x23 /* get character from keyboard, echoed onto the terminal */
	TRAP_PUTSP       = 0x24 /* output a byte string */
	TRAP_HALT        = 0x25 /* halt the program */
)

var reader *bufio.Reader
var writer *bufio.Writer

func _TRAP_GETC() {

}

func _TRAP_OUT() {

}

func _TRAP_PUTS() {
	loc := reg[R_R0]
	for {
		c := byte(memory[loc])
		loc++
		if c == 0 {
			break
		}
		C.putc(C.int(c), C.stdout)
	}
	C.fflush(C.stdout)
}

func _TRAP_IN() {

}

func _TRAP_PUTSP() {

}

func _TRAP_HALT() {

}
