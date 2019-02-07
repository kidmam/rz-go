package rz

import "time"

var (
	// DefaultTimestampFieldName is the default field name used for the timestamp field.
	DefaultTimestampFieldName = "timestamp"

	// DefaultLevelFieldName is the default field name used for the level field.
	DefaultLevelFieldName = "level"

	// DefaultMessageFieldName is the default field name used for the message field.
	DefaultMessageFieldName = "message"

	// DefaultErrorFieldName is the default field name used for error fields.
	DefaultErrorFieldName = "error"

	// DefaultCallerFieldName is the default field name used for caller field.
	DefaultCallerFieldName = "caller"

	// DefaultCallerSkipFrameCount is the default number of stack frames to skip to find the caller.
	DefaultCallerSkipFrameCount = 3

	// DefaultErrorStackFieldName is the default field name used for error stacks.
	DefaultErrorStackFieldName = "stack"

	// ErrorStackMarshaler extract the stack from err if any.
	ErrorStackMarshaler func(err error) interface{}

	// ErrorMarshalFunc allows customization of global error marshaling
	ErrorMarshalFunc = func(err error) interface{} {
		return err
	}

	// DefaultTimeFieldFormat defines the time format of the Time field type.
	// If set to an empty string, the time is formatted as an UNIX timestamp
	// as integer.
	DefaultTimeFieldFormat = time.RFC3339

	// TimestampFunc defines the function called to generate a timestamp.
	TimestampFunc = func() time.Time { return time.Now().UTC() }

	// DurationFieldUnit defines the unit for time.Duration type fields added
	// using the Dur method.
	DurationFieldUnit = time.Millisecond

	// DurationFieldInteger renders Dur fields as integer instead of float if
	// set to true.
	DurationFieldInteger = false

	// ErrorHandler is called whenever rz fails to write an event on its
	// output. If not set, an error is printed on the stderr. This handler must
	// be thread safe and non-blocking.
	ErrorHandler func(err error)
)
