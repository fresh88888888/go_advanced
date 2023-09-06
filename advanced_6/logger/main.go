package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"sync"
)

// Define levels
type Level int

const (
	ZDebug Level = iota
	ZInfo
	ZError
)

const defaultLogLevel = ZInfo

// Type logger represent a logging object that handles log entties based on the current log level
type Logger struct {
	mu     sync.Mutex   // for serialization
	prefix string       // prefix to write at begining of every log entry
	Level  Level        // log level
	w      io.Writer    // write for output
	buffer bytes.Buffer // internel buffer
}

// New creates a new Logger
func New(w io.Writer, prefix string) *Logger {
	return &Logger{
		prefix: prefix,
		Level:  defaultLogLevel,
		w:      w,
	}
}

var Console = New(os.Stderr, "")

func (l *Logger) Debug(v ...interface{}) {
	if ZDebug < l.Level {
		return
	}
	l.WriteEntry(ZDebug, fmt.Sprintln(v...))
}

func (l *Logger) Info(v ...interface{}) {
	if ZInfo < l.Level {
		return
	}
	l.WriteEntry(ZInfo, fmt.Sprintln(v...))
}

func (l *Logger) Error(v ...interface{}) {
	if ZError < l.Level {
		return
	}
	l.WriteEntry(ZError, fmt.Sprintln(v...))
}

// SetLevel sets the output level for the logger
func (l *Logger) SetLevel(lvl Level) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.Level = lvl
}

// GetLevel gets the output level of the
func (l *Logger) GetLevel() Level {
	return l.Level
}

// WriteEntry writes the msg of specified level to the underling writer
func (l *Logger) WriteEntry(lvl Level, msg string) error {
	l.w.Write([]byte(msg))
	return nil
}

func main() {
	Console.Info("Hellos")
	Console.Debug("World", "Debug")
	Console.Error("it is Error")

	Console.SetLevel(ZError)

	Console.Info("Hello")
	Console.Debug("World", "Debug")
	Console.Error("it is Error")

}
