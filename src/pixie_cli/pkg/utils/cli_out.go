package utils

import (
	"fmt"
	"io"
	"os"

	"github.com/fatih/color"
)

// CLIOutputEntry represents an output log entry.
type CLIOutputEntry struct {
	textColor *color.Color
	err       error
}

var defaultCLIOutput = &CLIOutputEntry{
	textColor: nil,
	err:       nil,
}

// WithColor returns a struct that can be used to log text to the CLI
// in a specific color.
func WithColor(c *color.Color) *CLIOutputEntry {
	return &CLIOutputEntry{
		textColor: c,
	}
}

// WithError returns a struct that can be used to log text to the CLI
// with a specific error.
func WithError(err error) *CLIOutputEntry {
	return &CLIOutputEntry{
		err: err,
	}
}

// Infof prints the input string to stdout formatted with the input args.
func Infof(format string, args ...interface{}) {
	defaultCLIOutput.Infof(format, args...)
}

// Info prints the input string to stdout.
func Info(str string) {
	defaultCLIOutput.Info(str)
}

// Errorf prints the input string to stderr formatted with the input args.
func Errorf(format string, args ...interface{}) {
	defaultCLIOutput.Errorf(format, args...)
}

// Error prints the input string to stderr.
func Error(str string) {
	defaultCLIOutput.Error(str)
}

// WithColor returns a struct that can be used to log text to the CLI
// in a specific color.
func (c *CLIOutputEntry) WithColor(textColor *color.Color) *CLIOutputEntry {
	return &CLIOutputEntry{
		err:       c.err,
		textColor: textColor,
	}
}

// WithError returns a struct that can be used to log text to the CLI
// with a specific error.
func (c *CLIOutputEntry) WithError(err error) *CLIOutputEntry {
	return &CLIOutputEntry{
		err:       err,
		textColor: c.textColor,
	}
}

func (c *CLIOutputEntry) write(w io.Writer, format string, args ...interface{}) {
	text := fmt.Sprintf(format, args...)
	if c.err != nil {
		text = text + fmt.Sprintf(" error=%s", c.err.Error())
	}
	text = text + "\n"
	if c.textColor == nil {
		fmt.Fprint(w, text)
	} else {
		c.textColor.Fprint(w, text)
	}
}

// Infof prints the input string to stdout formatted with the input args.
func (c *CLIOutputEntry) Infof(format string, args ...interface{}) {
	c.write(os.Stderr, format, args...)
}

// Info prints the input string to stdout.
func (c *CLIOutputEntry) Info(str string) {
	c.Infof(str)
}

// Errorf prints the input string to stderr formatted with the input args.
func (c *CLIOutputEntry) Errorf(format string, args ...interface{}) {
	c.write(os.Stderr, format, args...)
}

// Error prints the input string to stderr.
func (c *CLIOutputEntry) Error(str string) {
	c.write(os.Stderr, str)
}
