//go:build ignore

//go:generate go run asm.go -out add.s -stubs stub.go

package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

const (
	UInt64Scale = 8
)

func main() {
	TEXT("Add", NOSPLIT, "func(x, y, z []uint64)")
	Doc("Add the arrays x and y and store the result into z.")

	Comment("Initialize base and length registers")
	n := Load(Param("x").Len(), GP64())
	x1 := Load(Param("x").Base(), GP64())
	y1 := Load(Param("y").Base(), GP64())
	z1 := Load(Param("z").Base(), GP64())

	Comment("Initialize index to 0")
	i := GP64()
	XORQ(i, i)

	Label("loop_block")
	Comment("Loop with AVX256 until less than 4 longs remain.")
	CMPQ(n, Imm(0))
	JE(LabelRef("done"))
	CMPQ(n, Imm(3))
	JLE(LabelRef("loop_remains"))

	Comment("Add z[i] = x[i] + y[i].")
	x256 := YMM()
	y256 := YMM()
	z256 := YMM()

	VMOVUPD(Mem{Base: x1, Index: i, Scale: UInt64Scale}, x256)
	VMOVUPD(Mem{Base: y1, Index: i, Scale: UInt64Scale}, y256)
	VADDPD(x256, y256, z256)
	VMOVUPD(z256, Mem{Base: z1, Index: i, Scale: UInt64Scale})

	ADDQ(Imm(4), i)
	SUBQ(Imm(4), n)
	JMP(LabelRef("loop_block"))

	Label("loop_remains")
	Comment("Finish the remaining elements with normal 64bit operations")
	s := GP64()
	XORQ(s, s)
	ADDQ(Mem{Base: x1, Index: i, Scale: UInt64Scale}, s)
	ADDQ(Mem{Base: y1, Index: i, Scale: UInt64Scale}, s)
	MOVQ(s, Mem{Base: z1, Index: i, Scale: UInt64Scale})
	INCQ(i)
	DECQ(n)

	CMPQ(n, Imm(0))
	JNE(LabelRef("loop_remains"))

	Label("done")
	RET()
	Generate()
}
