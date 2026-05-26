// Package eris is an error handling library with readable stack traces and flexible formatting support.
package eris

import (
	"fmt"
)

// New creates a new root error with a static message.
func New(msg string) error {
	_ = "STUB: not implemented"
	// callers(3) skips this method, stack.callers, and runtime.Callers
	return nil
}

// Errorf creates a new root error with a formatted message.
func Errorf(format string, args ...interface{}) error { _ = "STUB: not implemented"; return nil }

// Wrap adds additional context to all error types while maintaining the type of the original error.
//
// This method behaves differently for each error type. For root errors, the stack trace is reset to the current
// callers which ensures traces are correct when using global/sentinel error values. Wrapped error types are simply
// wrapped with the new context. For external types (i.e. something other than root or wrap errors), this method
// attempts to unwrap them while building a new error chain. If an external type does not implement the unwrap
// interface, it flattens the error and creates a new root error from it before wrapping with the additional
// context.
func Wrap(err error, msg string) error { _ = "STUB: not implemented"; return nil }

// Wrapf adds additional context to all error types while maintaining the type of the original error.
//
// This is a convenience method for wrapping errors with formatted messages and is otherwise the same as Wrap.
func Wrapf(err error, format string, args ...interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func wrap(err error, msg string) error { _ = "STUB: not implemented"; return nil }

// callers(4) skips runtime.Callers, stack.callers, this method, and Wrap(f)

// caller(3) skips stack.caller, this method, and Wrap(f)
// caller(skip) has a slightly different meaning which is why it's not 4 as above

// create a new root error for global values to make sure nothing interferes with the stack

// insert the frame into the stack

// insert the frame into the stack

// return a new root error that wraps the external error

// Unwrap returns the result of calling the Unwrap method on err, if err's type contains an Unwrap method
// returning error. Otherwise, Unwrap returns nil.
func Unwrap(err error) error { _ = "STUB: not implemented"; return nil }

// Is reports whether any error in err's chain matches target.
//
// The chain consists of err itself followed by the sequence of errors obtained by repeatedly calling Unwrap.
//
// An error is considered to match a target if it is equal to that target or if it implements a method
// Is(error) bool such that Is(target) returns true.
func Is(err, target error) bool { _ = "STUB: not implemented"; return false }

// As finds the first error in err's chain that matches target. If there's a match, it sets target to that error
// value and returns true. Otherwise, it returns false.
//
// The chain consists of err itself followed by the sequence of errors obtained by repeatedly calling Unwrap.
//
// An error matches target if the error's concrete value is assignable to the value pointed to by target,
// or if the error has a method As(interface{}) bool such that As(target) returns true.
func As(err error, target interface{}) bool { _ = "STUB: not implemented"; return false }

// target must be a non-nil pointer

// *target must be interface or implement error

// Cause returns the root cause of the error, which is defined as the first error in the chain. The original
// error is returned if it does not implement `Unwrap() error` and nil is returned if the error is nil.
func Cause(err error) error { _ = "STUB: not implemented"; return nil }

// StackFrames returns the trace of an error in the form of a program counter slice.
// Use this method if you want to pass the eris stack trace to some other error tracing library.
func StackFrames(err error) []uintptr { _ = "STUB: not implemented"; return nil }

type rootError struct {
	global bool   // flag indicating whether the error was declared globally
	msg    string // root error message
	ext    error  // error type for wrapping external errors
	stack  *stack // root error stack trace
}

func (e *rootError) Error() string { _ = "STUB: not implemented"; return "" }

func (e *rootError) Format(s fmt.State, verb rune) { _ = "STUB: not implemented"; return }

func (e *rootError) Is(target error) bool { _ = "STUB: not implemented"; return false }

func (e *rootError) As(target interface{}) bool { _ = "STUB: not implemented"; return false }

func (e *rootError) Unwrap() error {
	_ = "STUB: not implemented"

	// StackFrames returns the trace of a root error in the form of a program counter slice.
	// This method is currently called by an external error tracing library (Sentry).
	return nil
}

func (e *rootError) StackFrames() []uintptr { _ = "STUB: not implemented"; return nil }

type wrapError struct {
	msg   string // wrap error message
	err   error  // error type representing the next error in the chain
	frame *frame // wrap error stack frame
}

func (e *wrapError) Error() string { _ = "STUB: not implemented"; return "" }

func (e *wrapError) Format(s fmt.State, verb rune) { _ = "STUB: not implemented"; return }

func (e *wrapError) Is(target error) bool { _ = "STUB: not implemented"; return false }

func (e *wrapError) As(target interface{}) bool { _ = "STUB: not implemented"; return false }

func (e *wrapError) Unwrap() error {
	_ = "STUB: not implemented"

	// StackFrames returns the trace of a wrap error in the form of a program counter slice.
	// This method is currently called by an external error tracing library (Sentry).
	return nil
}

func (e *wrapError) StackFrames() []uintptr { _ = "STUB: not implemented"; return nil }

func printError(err error, s fmt.State, verb rune) { _ = "STUB: not implemented"; return }
