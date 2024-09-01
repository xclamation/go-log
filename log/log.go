package log

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"
	"time"
)

// Logger interface with different logging levels
type Logger interface {
	Enable()
	Disable()

	IsEnabled() bool
	SetOutput(output io.Writer)
	GetOutput() io.Writer

	Alert(a ...interface{})
	Alertf(format string, a ...interface{})

	Error(a ...interface{})
	Errorf(format string, a ...interface{})

	Highlight(a ...interface{})
	Highlightf(format string, a ...interface{})

	Inform(a ...interface{})
	Informf(format string, a ...interface{})

	Log(a ...interface{})
	Logf(format string, a ...interface{})

	Trace(a ...interface{})
	Tracef(format string, a ...interface{})

	Warn(a ...interface{})
	Warnf(format string, a ...interface{})

	Prefix(...string) Logger
	GetPrefix() string

	Begin(newprefix ...string) Logger

	End()

	Level(uint8) Logger
}

// logger struct implementing the Logger interface
type logger struct {
	enabled   bool      // Unexported fields so other can't change them
	output    io.Writer // In the way that I did not intend
	prefix    strings.Builder
	funcName  string // string.Builder is not necessary because we do not need to modify funcName dinamically
	startTime time.Time
	level     uint8
}

func NewLogger(opts ...Option) Logger {
	var defaultOutput io.Writer = os.Stdout
	const defaultEnabled bool = true
	const defaultLevel uint8 = 6

	l := &logger{
		enabled: defaultEnabled,
		output:  defaultOutput,
		level:   defaultLevel,
	}

	// If can uncomma that if we want to have funcName when NewLogger() is used
	//l.captureFuncName()
	// l.prefix.WriteString(l.funcName)
	// l.prefix.WriteString(": ")

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
type Option func(*logger)

// Returns Option = func(l *Logger) {...} so opt(l) = func(l) {...}
func WithEnabled(enabled bool) Option {
	return func(l *logger) {
		l.enabled = enabled
	}
}

// WithOutput sets the output field
func WithOutput(output io.Writer) Option {
	return func(l *logger) {
		l.output = output
	}
}

func WithPrefix(newprefix ...string) Option {
	return func(l *logger) {
		for _, s := range newprefix {
			l.prefix.WriteString(s)
			l.prefix.WriteString(": ")
		}
	}
}

func (l *logger) Begin(newprefix ...string) Logger {
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
	l.logMessage("BEGIN\n")
	l.logfMessage("Execution started at: %v\n", l.startTime.Format(time.RFC3339))

	return l
}

func (l *logger) End() {
	endTime := time.Now()
	l.logfMessage("END\n")
	l.logfMessage("Execution ended at: %v\n", endTime.Format(time.RFC3339))
	duration := endTime.Sub(l.startTime)
	l.logfMessage("Execution duration: %v\n", duration)
}

func (l *logger) Enable() {
	l.enabled = true
}

func (l *logger) Disable() {
	l.enabled = false
}

func (l *logger) IsEnabled() bool {
	return l.enabled
}

func (l *logger) SetOutput(output io.Writer) {
	l.output = output
}

func (l *logger) GetOutput() io.Writer {
	return l.output
}

func (l *logger) Level(level uint8) Logger {
	l.level = level
	return l
}

func (l *logger) logMessage(a ...interface{}) {
	if l.enabled {
		fmt.Fprint(l.output, l.prefix.String())
		fmt.Fprint(l.output, a...)
	}
}

func (l *logger) logfMessage(format string, a ...interface{}) {
	if l.enabled {
		fmt.Fprint(l.output, l.prefix.String())
		fmt.Fprintf(l.output, format, a...)
	}
}

// Alert logs an alert-level message
func (l *logger) Alert(a ...interface{}) {
	if l.level >= 1 {
		l.logMessage(append([]interface{}{"ALERT: "}, a...)...)
	}
}

func (l *logger) Alertf(format string, a ...interface{}) {
	if l.level >= 1 {
		l.logfMessage(format, append([]interface{}{"ALERT: "}, a...)...)
	}
}

// Error logs an error-level message.
func (l *logger) Error(a ...interface{}) {
	if l.level >= 1 {
		l.logMessage(append([]interface{}{"ERROR: "}, a...)...)
	}
}

func (l *logger) Errorf(format string, a ...interface{}) {
	if l.level >= 1 {
		l.logfMessage(format, append([]interface{}{"ERROR: "}, a...)...)
	}
}

// Warn logs a warning-level message.
func (l *logger) Warn(a ...interface{}) {
	if l.level >= 2 {
		l.logMessage(append([]interface{}{"WARN: "}, a...)...)
	}
}

func (l *logger) Warnf(format string, a ...interface{}) {
	if l.level >= 2 {
		l.logfMessage(format, append([]interface{}{"WARN: "}, a...)...)
	}
}

// Highlight logs a highlight-level message.
func (l *logger) Highlight(a ...interface{}) {
	if l.level >= 3 {
		l.logMessage(append([]interface{}{"HIGHLIGHT: "}, a...)...)
	}
}

func (l *logger) Highlightf(format string, a ...interface{}) {
	if l.level >= 3 {
		l.logfMessage(format, append([]interface{}{"HIGHLIGHT: "}, a...)...)
	}
}

// Inform logs an information-level message.
func (l *logger) Inform(a ...interface{}) {
	if l.level >= 4 {
		l.logMessage(append([]interface{}{"INFORM: "}, a...)...)
	}
}

func (l *logger) Informf(format string, a ...interface{}) {
	if l.level >= 4 {
		l.logfMessage(format, append([]interface{}{"INFORM: "}, a...)...)
	}
}

// Log logs a general-level message.
func (l *logger) Log(a ...interface{}) {
	if l.level >= 5 {
		l.logMessage(append([]interface{}{"LOG: "}, a...)...)
	}
}

func (l *logger) Logf(format string, a ...interface{}) {
	if l.level >= 5 {
		l.logfMessage(format, append([]interface{}{"LOG: "}, a...)...)
	}
}

// Trace logs a trace-level message.
func (l *logger) Trace(a ...interface{}) {
	if l.level >= 6 {
		l.logMessage(append([]interface{}{"TRACE: "}, a...)...)
	}
}

func (l *logger) Tracef(format string, a ...interface{}) {
	if l.level >= 6 {
		l.logfMessage(format, append([]interface{}{"TRACE: "}, a...)...)
	}
}

// Add prefix to logger.
func (l *logger) Prefix(newprefix ...string) Logger {
	for _, s := range newprefix {
		l.prefix.WriteString(s)
		l.prefix.WriteString(": ")
	}

	return l
}

func (l *logger) GetPrefix() string {
	return l.prefix.String()
}

// captureFuncName captures the name of the function from which Begin() was called.
func (l *logger) captureFuncName() {
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
