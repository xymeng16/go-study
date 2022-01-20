package main

import "C"
import (
	"fmt"
	"github.com/pkg/term/termios"
	"golang.org/x/sys/unix"
	"os"
	"os/signal"
)

/* 65536 memory locations */
var memory []uint16

/* 10 registers */
const (
	R_R0 uint8 = iota
	R_R1
	R_R2
	R_R3
	R_R4
	R_R5
	R_R6
	R_R7
	R_PC /* program counter */
	R_COND
	R_COUNT
)

var reg [R_COUNT]uint16

/* 16 opcodes */
const (
	OP_BR   uint8 = iota /* branch */
	OP_ADD               /* add  */
	OP_LD                /* load */
	OP_ST                /* store */
	OP_JSR               /* jump register */
	OP_AND               /* bitwise and */
	OP_LDR               /* load register */
	OP_STR               /* store register */
	OP_RTI               /* unused */
	OP_NOT               /* bitwise not */
	OP_LDI               /* load indirect */
	OP_STI               /* store indirect */
	OP_JMP               /* jump */
	OP_RES               /* reserved (unused) */
	OP_LEA               /* load effective address */
	OP_TRAP              /* execute trap */
)

/* condition flags */
const (
	FL_POS uint8 = 1 << 0 /* positive */
	FL_ZRO       = 1 << 1 /* zero */
	FL_NEG       = 1 << 2 /* negative */
)

func read_image(f string) bool {
	return false
}

func check_key() uint16 {
	var readfds unix.FdSet

	var timeout unix.Timeval
	timeout.Sec = 0
	timeout.Usec = 0

	readfds.Zero()
	readfds.Set(unix.Stdin)

	n, err := unix.Select(1, &readfds, nil, nil, &timeout)
	if err != nil {
		return uint16(n)
	} else {
		return 0
	}
}

var original_tio unix.Termios

func disable_input_buffering() {
	tmp_tio, _ := termios.Tcgetattr(uintptr(unix.Stdin))
	original_tio = *tmp_tio
	new_tio := original_tio
	//fmt.Printf("%x\n", new_tio.Lflag & uint64(-265))
	//new_tio.Lflag &= ^termios.ICANON & ^termios.ECHO // -101 & -9
	new_tio.Lflag = 0x4c3

	termios.Tcsetattr(uintptr(unix.Stdin), termios.TCSANOW, &new_tio)
}

func restore_input_buffering() {
	termios.Tcsetattr(uintptr(unix.Stdin), termios.TCSANOW, &original_tio)
}

func handle_interrupt(signal int) {
	restore_input_buffering()
	fmt.Println()
	os.Exit(-2)
}

func setup_sigint_handler() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		fmt.Println("\r- Ctrl+C pressed in Terminal")
		os.Exit(-2)
	}()
}

func main() {
	/* load arguments */
	args := os.Args
	if len(args) < 2 {
		/* show usage string */
		fmt.Println("usage: lc3 [image-file1] ...")
		os.Exit(2)
	}

	for j := 1; j < len(args); j++ {
		if !read_image(args[j]) {
			fmt.Printf("failed to load image %s\n...", args[j])
			os.Exit(1)
		}
	}

	/* setup terminal */
	setup_sigint_handler()
	disable_input_buffering()

	/* since exactly one condition flag should be set at any given time, set the Z flag */
	reg[R_COND] = FL_ZRO

	/* set the PC to starting position */
	/* 0x3000 is the default */
	const PC_START = 0x3000
	reg[R_PC] = PC_START

	running := true
	for running {
		/* FETCH */
		instr := mem_read(reg[R_PC]++)
		op := instr >> 12
		reg[R_PC]++


	}
}

func mem_read(u uint16) uint16 {
	return u
}

func exec_op(op uint16) error {

}