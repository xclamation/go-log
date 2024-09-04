package arg

import (
	"flag"
	"testing"
	"github/xclamation/go-log/loglevel"
)

func setFlags(args []string) {
	// Clear any existing flags
	flag.CommandLine = flag.NewFlagSet("", flag.ExitOnError)

	// Reinitialize flags
	initFlags()

	// Parse the flags with provided arguments
	flag.CommandLine.Parse(args)
}

func TestLevelFlag(t *testing.T) {
	// TableDrivenTest
	tests := []struct {
		args []string
		expected uint8
	}{
		{[]string{"-v"}, loglevel.LEVEL_1},
		{[]string{"-vv"}, loglevel.LEVEL_2},
		{[]string{"-vvv"}, loglevel.LEVEL_3},
		{[]string{"-vvvv"}, loglevel.LEVEL_4},
		{[]string{"-vvvvv"}, loglevel.LEVEL_5},
		{[]string{"-vvvvvv"}, loglevel.LEVEL_6},
		{[]string{}, loglevel.LEVEL_0},
		{[]string{"-v", "-vv"}, loglevel.LEVEL_2}, // Higher level should take precedence
	}

	for _, tt := range tests {
		setFlags(tt.args)
		if initLevel != tt.expected {
			t.Errorf("For args %v, expected level %d, but got %d", tt.args, tt.expected, initLevel)
		}
	}
}
