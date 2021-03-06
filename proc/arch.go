package proc

import "runtime"

type Arch interface {
	PtrSize() int
	BreakpointInstruction() []byte
	BreakpointSize() int
	CurgInstructions() []byte
	HardwareBreakpointUsage() []bool
	SetHardwareBreakpointUsage(int, bool)
}

type AMD64 struct {
	ptrSize                 int
	breakInstruction        []byte
	breakInstructionLen     int
	curgInstructions        []byte
	hardwareBreakpointUsage []bool
}

func AMD64Arch() *AMD64 {
	var (
		curg       []byte
		breakInstr = []byte{0xCC}
	)

	switch runtime.GOOS {
	case "darwin":
		curg = []byte{
			0x65, 0x48, 0x8b, 0x0C, 0x25, 0xA0, 0x08, // mov %gs:0x8a0,%rcx
			0x0, 0x0,
		}
	case "linux":
		curg = []byte{
			0x64, 0x48, 0x8b, 0x0c, 0x25, 0xf0, 0xff, 0xff, 0xff, // mov %fs:0xfffffffffffffff0,%rcx
		}
	}
	curg = append(curg, breakInstr[0])

	return &AMD64{
		ptrSize:                 8,
		breakInstruction:        breakInstr,
		breakInstructionLen:     1,
		curgInstructions:        curg,
		hardwareBreakpointUsage: make([]bool, 4),
	}
}

func (a *AMD64) PtrSize() int {
	return a.ptrSize
}

func (a *AMD64) BreakpointInstruction() []byte {
	return a.breakInstruction
}

func (a *AMD64) BreakpointSize() int {
	return a.breakInstructionLen
}

func (a *AMD64) CurgInstructions() []byte {
	return a.curgInstructions
}

func (a *AMD64) HardwareBreakpointUsage() []bool {
	return a.hardwareBreakpointUsage
}

func (a *AMD64) SetHardwareBreakpointUsage(reg int, set bool) {
	a.hardwareBreakpointUsage[reg] = set
}
