package log

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"
	"time"
)

type Logger struct {
	enabled   bool      // Unexported fields so other can't change them
	output    io.Writer // In the way that I did not intend
	prefix    strings.Builder
	funcName  string // string.Builder is not necessary because we do not need to modify funcName dinamically
	startTime time.Time
}

func NewLogger(opts ...Option) *Logger {
	var defaultOutput io.Writer = os.Stdout
	const defaultEnabled bool = true

	l := &Logger{
		enabled: defaultEnabled,
		output:  defaultOutput,
	}

	l.captureFuncName()
	l.prefix.WriteString(l.funcName)
	l.prefix.WriteString(": ")

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

func WithPrefix(newprefix ...string) Option {
	return func(l *Logger) {
		for _, s := range newprefix {
			l.prefix.WriteString(s)
			l.prefix.WriteString(": ")
		}
	}
}

func (l *Logger) Log(a ...interface{}) {
	if l.enabled {
		fmt.Fprint(l.output, l.prefix.String())
		fmt.Fprint(l.output, a...)
	}
}

func (l *Logger) Logf(format string, a ...interface{}) {
	if l.enabled {
		fmt.Fprint(l.output, l.prefix.String())
		fmt.Fprintf(l.output, format, a...)
	}
}

func (l *Logger) Begin(newprefix ...string) *Logger {
	l.captureFuncName()
	if l.funcName != "" {
		l.prefix.WriteString(l.funcName)
		l.prefix.WriteString(": ")
	}
	for _, s := range newprefix {
		l.prefix.WriteString(s)
		l.prefix.WriteString(": ")
	}

	l.startTime = time.Now()
	l.Log("BEGIN\n")
	l.Logf("Execution started at: %v\n", l.startTime.Format(time.RFC3339))

	return l
}

func (l *Logger) End() {
	endTime := time.Now()
	l.Log("END\n")
	l.Logf("Execution ended at: %v", endTime.Format(time.RFC3339))
	duration := endTime.Sub(l.startTime)
	l.Logf("Execution duration: %v\n", duration)
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

// Add prefix to logger
func (l *Logger) Prefix(newprefix ...string) {
	for _, s := range newprefix {
		l.prefix.WriteString(s)
		l.prefix.WriteString(": ")
	}
}

// captureFuncName captures the name of the function from which Begin() was called.
func (l *Logger) captureFuncName() {
	pc, _, _, ok := runtime.Caller(2) // 2 levels up to get the calling function
	if ok {
		fn := runtime.FuncForPC(pc) // pc stands for program counter
		fullFuncName := fn.Name()
		// Extract the last part after the last "/"
		parts := strings.Split(fullFuncName, ".")
		l.funcName = parts[len(parts)-1] // Instead of packageName.funcName get only funcName
		//return l.funcName
	}
	//return ""
}
