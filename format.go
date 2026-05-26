package eris

// FormatOptions defines output options like omitting stack traces and inverting the error or stack order.
type FormatOptions struct {
	InvertOutput bool // Flag that inverts the error output (wrap errors shown first).
	WithTrace    bool // Flag that enables stack trace output.
	InvertTrace  bool // Flag that inverts the stack trace output (top of call stack shown first).
	WithExternal bool // Flag that enables external error output.
	// todo: maybe allow users to hide wrap frames if desired
}

// StringFormat defines a string error format.
type StringFormat struct {
	Options      FormatOptions // Format options (e.g. omitting stack trace or inverting the output order).
	MsgStackSep  string        // Separator between error messages and stack frame data.
	PreStackSep  string        // Separator at the beginning of each stack frame.
	StackElemSep string        // Separator between elements of each stack frame.
	ErrorSep     string        // Separator between each error in the chain.
}

// NewDefaultStringFormat returns a default string output format.
func NewDefaultStringFormat(options FormatOptions) StringFormat {
	_ = "STUB: not implemented"
	return *new(StringFormat)
}

// ToString returns a default formatted string for a given error.
//
// An error without trace will be formatted as follows:
//
//	<Wrap error msg>: <Root error msg>
//
// An error with trace will be formatted as follows:
//
//	<Wrap error msg>
//	  <Method2>:<File2>:<Line2>
//	<Root error msg>
//	  <Method2>:<File2>:<Line2>
//	  <Method1>:<File1>:<Line1>
func ToString(err error, withTrace bool) string { _ = "STUB: not implemented"; return "" }

// ToCustomString returns a custom formatted string for a given error.
//
// To declare custom format, the Format object has to be passed as an argument.
// An error without trace will be formatted as follows:
//
//	<Wrap error msg>[Format.ErrorSep]<Root error msg>
//
// An error with trace will be formatted as follows:
//
//	<Wrap error msg>[Format.MsgStackSep]
//	[Format.PreStackSep]<Method2>[Format.StackElemSep]<File2>[Format.StackElemSep]<Line2>[Format.ErrorSep]
//	<Root error msg>[Format.MsgStackSep]
//	[Format.PreStackSep]<Method2>[Format.StackElemSep]<File2>[Format.StackElemSep]<Line2>[Format.ErrorSep]
//	[Format.PreStackSep]<Method1>[Format.StackElemSep]<File1>[Format.StackElemSep]<Line1>[Format.ErrorSep]
func ToCustomString(err error, format StringFormat) string { _ = "STUB: not implemented"; return "" }

// JSONFormat defines a JSON error format.
type JSONFormat struct {
	Options FormatOptions // Format options (e.g. omitting stack trace or inverting the output order).
	// todo: maybe allow setting of wrap/root keys in the output map as well
	StackElemSep string // Separator between elements of each stack frame.
}

// NewDefaultJSONFormat returns a default JSON output format.
func NewDefaultJSONFormat(options FormatOptions) JSONFormat {
	_ = "STUB: not implemented"
	return *new(JSONFormat)
}

// ToJSON returns a JSON formatted map for a given error.
//
// An error without trace will be formatted as follows:
//
//	{
//	  "root": {
//	      "message": "Root error msg"
//	  },
//	  "wrap": [
//	    {
//	      "message": "Wrap error msg"
//	    }
//	  ]
//	}
//
// An error with trace will be formatted as follows:
//
//	{
//	  "root": {
//	    "message": "Root error msg",
//	    "stack": [
//	      "<Method2>:<File2>:<Line2>",
//	      "<Method1>:<File1>:<Line1>"
//	    ]
//	  },
//	  "wrap": [
//	    {
//	      "message": "Wrap error msg",
//	      "stack": "<Method2>:<File2>:<Line2>"
//	    }
//	  ]
//	}
func ToJSON(err error, withTrace bool) map[string]interface{} {
	_ = "STUB: not implemented"
	return nil
}

// ToCustomJSON returns a JSON formatted map for a given error.
//
// To declare custom format, the Format object has to be passed as an argument.
// An error without trace will be formatted as follows:
//
//	{
//	  "root": {
//	    "message": "Root error msg",
//	  },
//	  "wrap": [
//	    {
//	      "message": "Wrap error msg'",
//	    }
//	  ]
//	}
//
// An error with trace will be formatted as follows:
//
//	{
//	  "root": {
//	    "message": "Root error msg",
//	    "stack": [
//	      "<Method2>[Format.StackElemSep]<File2>[Format.StackElemSep]<Line2>",
//	      "<Method1>[Format.StackElemSep]<File1>[Format.StackElemSep]<Line1>"
//	    ]
//	  }
//	  "wrap": [
//	    {
//	      "message": "Wrap error msg",
//	      "stack": "<Method2>[Format.StackElemSep]<File2>[Format.StackElemSep]<Line2>"
//	    }
//	  ]
//	}
func ToCustomJSON(err error, format JSONFormat) map[string]interface{} {
	_ = "STUB: not implemented"
	return nil
}

// Unpack returns a human-readable UnpackedError type for a given error.
func Unpack(err error) UnpackedError { _ = "STUB: not implemented"; return *new(UnpackedError) }

// prepend links in stack trace order

// UnpackedError represents complete information about an error.
//
// This type can be used for custom error logging and parsing. Use `eris.Unpack` to build an UnpackedError
// from any error type. The ErrChain and ErrRoot fields correspond to `wrapError` and `rootError` types,
// respectively. If any other error type is unpacked, it will appear in the ExternalErr field.
type UnpackedError struct {
	ErrExternal error
	ErrRoot     ErrRoot
	ErrChain    []ErrLink
}

// String formatter for external errors.
func formatExternalStr(err error, withTrace bool) string { _ = "STUB: not implemented"; return "" }

// ErrRoot represents an error stack and the accompanying message.
type ErrRoot struct {
	Msg   string
	Stack Stack
}

// String formatter for root errors.
func (err *ErrRoot) formatStr(format StringFormat) string { _ = "STUB: not implemented"; return "" }

// JSON formatter for root errors.
func (err *ErrRoot) formatJSON(format JSONFormat) map[string]interface{} {
	_ = "STUB: not implemented"
	return nil
}

// ErrLink represents a single error frame and the accompanying message.
type ErrLink struct {
	Msg   string
	Frame StackFrame
}

// String formatter for wrap errors chains.
func (eLink *ErrLink) formatStr(format StringFormat) string { _ = "STUB: not implemented"; return "" }

// JSON formatter for wrap error chains.
func (eLink *ErrLink) formatJSON(format JSONFormat) map[string]interface{} {
	_ = "STUB: not implemented"
	return nil
}
