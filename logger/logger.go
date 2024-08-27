package logger

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Logger struct {
	enabled bool      // Unexported fields so other can't change them
	output  io.Writer // in the way that I did not intend
}

func NewLogger(opts ...Option) *Logger {
	var defaultOutput io.Writer = os.Stdout
	const defaultEnabled bool = true

	l := &Logger{
		enabled: defaultEnabled,
		output:  defaultOutput,
	}

	// Apply options
	// It's more convenient way to initialize fields of the instance
	// You can pass options in different orders
	// You don't need to add more input variables into constructor,
	// just a new function with required Option in return
	for _, opt := range opts {
		opt(l) // example: opt(l) = (WithEnabled(true))(l) = func(l) {l.enabled = true}
	}

	return l
}

// Option is a function type that modifies Logger settings
// It's used to write "Option" instead of "func(*Logger)" every time
type Option func(*Logger)

// Returns Option = func(l *Logger) {...} so opt(l) = func(l) {...}
func WithEnabled(enabled bool) Option {
	return func(l *Logger) {
		l.enabled = enabled
	}
}

// WithOutput sets the output field
func WithOutput(output io.Writer) Option {
	return func(l *Logger) {
		l.output = output
	}
}

func (l *Logger) Log(a ...interface{}) {
	if l.enabled {
		fmt.Fprintln(l.output, a...)
	}
}

func (l *Logger) Logf(format string, a ...interface{}) {
	if l.enabled {
		fmt.Fprintf(l.output, format, a...)
	}
}

func (l *Logger) Begin() {
	l.Log("BEGIN")
	l.Log("Execution started at:", time.Now().Format(time.RFC3339))
}

func (l *Logger) End() {
	l.Log("END")
	l.Log("Execution ended at:", time.Now().Format(time.RFC3339))
}

func (l *Logger) Enable() {
	l.enabled = true
}

func (l *Logger) Disable() {
	l.enabled = false
}

func (l *Logger) SetOutput(output io.Writer) {
	l.output = output
}
