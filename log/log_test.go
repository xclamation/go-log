package log

import (
	"io"
	"os"
	"strings"
	"testing"
)

// TestNewLogger_Default verifies that the default logger is enabled and outputs to os.Stdout.
func TestNewLogger_Default(t *testing.T) {
	logger := NewLogger()

	if !logger.enabled {
		t.Errorf("Expected  logger to be enabled by default")
	}

	if logger.output != os.Stdout {
		t.Errorf("Expected default output to be os.Stdout")
	}
}

// TestNewLogger_WithEnabled verifies that the logger can be enabled or disabled via options.
func TestNewLogger_WithEnabled(t *testing.T) {
	logger := NewLogger(WithEnabled(false))

	if logger.enabled {
		t.Errorf("Expected logger to be disabled")
	}

	logger.enabled = true

	if !logger.enabled {
		t.Error("Expected logger to be enabled")
	}
}

func TestLog(t *testing.T) {
	// We can also use buf for output for binary data such as media files, binary protocols, etc.
	// var buf bytes.Buffer
	var output strings.Builder
	var writer io.Writer = &output
	logger := NewLogger(WithOutput(writer)) //WithOutput(&buf)

	logger.Log("Test message")

	expected := "Test message\n"
	if output.String() != expected {
		t.Errorf("Expected log output to be %q, got %q", expected, output.String())
	}

}
