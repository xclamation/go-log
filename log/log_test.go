package log

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
)

// TestNewLogger_Default verifies that the default logger is enabled and outputs to os.Stdout.
func TestNewLogger_Default(t *testing.T) {
	logger := NewLogger()

	if !logger.IsEnabled() {
		t.Errorf("Expected  logger to be enabled by default")
	}

	if logger.GetOutput() != os.Stdout {
		t.Errorf("Expected default output to be os.Stdout")
	}
}

// TestNewLogger_WithEnabled verifies that the logger can be enabled or disabled via options.
func TestNewLogger_WithEnabled(t *testing.T) {
	logger := NewLogger(WithEnabled(false))

	if logger.IsEnabled() {
		t.Errorf("Expected logger to be disabled")
	}

	logger.Enable()

	if !logger.IsEnabled() {
		t.Error("Expected logger to be enabled")
	}
}

func TestNewLogger_WithOutput(t *testing.T) {
	var output strings.Builder
	var writer io.Writer = &output
	logger := NewLogger(WithOutput(writer))

	if logger.GetOutput() != writer {
		t.Errorf("Expected logger to be a \"%T\"", writer)
	}

	var buf bytes.Buffer
	writer = &buf
	logger.SetOutput(writer)
	logOutput := logger.GetOutput()
	if logOutput != writer {
		t.Errorf("Expected logger to be a \"%T\", got \"%T\"", writer, logOutput)
	}
}

func TestLog(t *testing.T) {
	// We can also use buf for output for binary data such as media files, binary protocols, etc.
	// var buf bytes.Buffer
	var output strings.Builder
	var writer io.Writer = &output

	// io.StringWriter is more efficient for operations with strings
	// but WithOutput() gets io.Writer as a parameter for comprehensiveness

	//var writer io.StringWriter = &output

	logger := NewLogger(WithOutput(writer)) //WithOutput(&buf)

	logger.Log("Test message")

	expected := "Test message"
	if output.String() != expected {
		t.Errorf("Expected log output to be %q, got %q", expected, output.String())
	}
}

func TestLogf(t *testing.T) {
	var output strings.Builder
	var writer io.Writer = &output

	logger := NewLogger(WithOutput(writer))
	message := "message"

	logger.Logf("Test %s", message)

	expected := fmt.Sprintf("Test %s", message)

	if output.String() != expected {
		t.Errorf("Expected log output to be %q, got %q", expected, output.String())
	}
}

// TestBeginEnd_Verifies the Begin and End methods produce the expected output.
func TestBeginEnd(t *testing.T) {
	// Using buf for example
	var buf bytes.Buffer
	logger := NewLogger(WithOutput(&buf))

	logger.Begin()
	logger.End()

	// We can't match the exact time string, so we just check the static parts
	expectedBegin := "BEGIN\nExecution started at:"
	expectedEnd := "END\nExecution ended at:"

	if !bytes.Contains(buf.Bytes(), []byte(expectedBegin)) {
		t.Errorf("Expected log output to contain %q, got %q", expectedBegin, buf.String())
	}

	if !bytes.Contains(buf.Bytes(), []byte(expectedEnd)) {
		t.Errorf("Expected log output to contain %q, got %q", expectedEnd, buf.String())
	}
}

// TestSetOutput verifies that the logger's output can be changed dynamically.
func TestSetOutput(t *testing.T) {
	var buf1, buf2 bytes.Buffer
	var writer1, writer2 io.Writer = &buf1, &buf2
	logger := NewLogger(WithOutput(writer1))

	logger.Log("First message")
	logger.SetOutput(writer2)
	logger.Log("Second message")

	expected1 := "First message"
	expected2 := "Second message"

	if buf1.String() != expected1 {
		t.Errorf("Expected first log output to be %q, got %q", expected1, buf1.String())
	}

	if buf2.String() != expected2 {
		t.Errorf("Expected first log output to be %q, got %q", expected2, buf2.String())
	}
}

// TestLog_Disabled: Verifies that the Log method does not write output when the logger is disabled.TestLogf_Disabled: Verifies that the Logf method does not write formatted output when the logger is disabled.
// TestEnable: Tests that the Enable method correctly enables logging.
// TestDisable: Tests that the Disable method correctly disables logging.

func TestPrefix(t *testing.T) {
	prefixedLogger := NewLogger(WithPrefix("Z", "O", "V"))
	initialPrefix := "TestPrefix: Z: O: V: "
	logPrefix := prefixedLogger.GetPrefix()
	if logPrefix != initialPrefix {
		t.Errorf("Expected initial prefix to be %q, got %q", initialPrefix, logPrefix)
	}

	prefixedLogger.Prefix("apple", "banana", "cherry")
	expectedPrefix := "TestPrefix: Z: O: V: apple: banana: cherry: "
	logPrefix = prefixedLogger.GetPrefix()
	if logPrefix != expectedPrefix {
		t.Errorf("Expected initial prefix to be %q, got %q", expectedPrefix, logPrefix)
	}

	prefixedLogger.Prefix("date")
	expectedPrefix = "TestPrefix: Z: O: V: apple: banana: cherry: date: "
	logPrefix = prefixedLogger.GetPrefix()
	if logPrefix != expectedPrefix {
		t.Errorf("Expected initial prefix to be %q, got %q", expectedPrefix, logPrefix)
	}
}

func BenchmarkLogger_Log(b *testing.B) {
	var output strings.Builder
	var writer io.Writer = &output
	logger := NewLogger(WithOutput(writer))

	for i := 0; i < b.N; i++ {
		logger.Log("Test message")
	}
}
