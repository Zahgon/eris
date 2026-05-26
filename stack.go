package eris

// Stack is an array of stack frames stored in a human readable format.
type Stack []StackFrame

// format returns an array of formatted stack frames.
func (s Stack) format(sep string, invert bool) []string { _ = "STUB: not implemented"; return nil }

// StackFrame stores a frame's runtime information in a human readable format.
type StackFrame struct {
	Name string
	File string
	Line int
}

// format returns a formatted stack frame.
func (f *StackFrame) format(sep string) string { _ = "STUB: not implemented"; return "" }

// caller returns a single stack frame. the argument skip is the number of stack frames
// to ascend, with 0 identifying the caller of Caller.
func caller(skip int) *frame { _ = "STUB: not implemented"; return nil }

// frame is a single program counter of a stack frame.
type frame uintptr

// pc returns the program counter for a frame.
func (f frame) pc() uintptr { _ = "STUB: not implemented"; return 0 }

// get returns a human readable stack frame.
func (f frame) get() StackFrame { _ = "STUB: not implemented"; return *new(StackFrame) }

// callers returns a stack trace. the argument skip is the number of stack frames to skip before recording
// in pc, with 0 identifying the frame for Callers itself and 1 identifying the caller of Callers.
func callers(skip int) *stack { _ = "STUB: not implemented"; return nil }

// todo: change this to filtering out runtime instead of hardcoding n-2

// stack is an array of program counters.
type stack []uintptr

// insertPC inserts a wrap error program counter (pc) into the correct place of the root error stack trace.
func (s *stack) insertPC(wrapPCs stack) { _ = "STUB: not implemented"; return }

// append the pc to the end if there's only one

// break if the stack already contains the pc

// insert the first pc into the stack if the second pc is found

// get returns a human readable stack trace.
func (s *stack) get() []StackFrame { _ = "STUB: not implemented"; return nil }

// isGlobal determines if the stack trace represents a global error
func (s *stack) isGlobal() bool { _ = "STUB: not implemented"; return false }

func insert(s stack, u uintptr, at int) stack {
	_ = "STUB: not implemented"
	// this inserts the pc by breaking the stack into two slices (s[:at] and s[at:])
	return *new(stack)
}
