// Package logger is the library that will be used for logging
package logger

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
)

type logLevels struct {
	PANIC   logLevelType
	FATAL   logLevelType
	ERROR   logLevelType
	WARNING logLevelType
	INFO    logLevelType
	DEBUG   logLevelType
	TRACE   logLevelType
}

var (
	showTraces = false
	// ErrInvalidFormatter error for invalid format.
	ErrInvalidFormatter = errors.New("invalid formatter")
	logger              = &logrus.Logger{
		Out:          os.Stderr,
		Formatter:    &logrus.JSONFormatter{},
		Hooks:        make(logrus.LevelHooks),
		Level:        logrus.DebugLevel,
		ExitFunc:     os.Exit,
		ReportCaller: false,
	}
	// LogLevels the available log levels.
	LogLevels = logLevels{
		PANIC:   "PANIC",
		FATAL:   "FATAL",
		ERROR:   "ERROR",
		WARNING: "WARNING",
		INFO:    "INFO",
		DEBUG:   "DEBUG",
		TRACE:   "TRACE",
	}
)

type exportedLogger struct {
	identifier string
	level      logLevelType
}

type logLevelType string

func NewLogger(identifier string, levelType logLevelType) *exportedLogger {
	return &exportedLogger{
		identifier: identifier,
		level:      levelType,
	}
}

// Printf
// nolint:goerr113 // Cannot declare a static error with a dynamic message.
func (expLogger *exportedLogger) Printf(format string, v ...interface{}) {
	msgStr := fmt.Sprintf(format, v...)
	ctx := context.Background()

	if expLogger.identifier != "" {
		ctx = context.WithValue(context.Background(), Settings.IdentifierKey, expLogger.identifier)
	}

	switch expLogger.level {
	case LogLevels.PANIC:
		Panic(ctx, errors.New(msgStr))
	case LogLevels.FATAL:
		Fatal(ctx, errors.New(msgStr))
	case LogLevels.ERROR:
		Error(ctx, errors.New(msgStr))
	case LogLevels.WARNING:
		Warning(ctx, msgStr)
	case LogLevels.INFO:
		Info(ctx, msgStr)
	case LogLevels.DEBUG:
		Debug(ctx, msgStr)
	case LogLevels.TRACE:
		Trace(ctx, msgStr)
	}
}

type StackTrace struct {
	File     string
	Line     int
	Function string
}

// SetOutput changes the output of the logger.
func SetOutput(out io.Writer) {
	logger.SetOutput(out)
}

// SetFormatter changes the formatter for the logs.
// Valid values are: "json", "text".
func SetFormatter(formatter string) error {
	switch strings.ToLower(formatter) {
	case "json":
		logger.SetFormatter(&logrus.JSONFormatter{})
	case "text":
		logger.SetFormatter(&logrus.TextFormatter{
			FullTimestamp: true,
		})
	default:
		return fmt.Errorf("%w : %v", ErrInvalidFormatter, formatter)
	}

	return nil
}

// SetLogLevel changes the level of the logging util. Acceptable strings are based on logrus.
func SetLogLevel(level string) error {
	logLevel, err := logrus.ParseLevel(level)
	if err != nil {
		return fmt.Errorf("could not parse error level %v : %w", level, err)
	}

	logger.SetLevel(logLevel)

	return nil
}

// SetLogTrace changes the showTrace option.
func SetLogTrace(show bool) {
	showTraces = show
}

// Panic gets the transaction from context.
func Panic(ctx context.Context, err error, messages ...interface{}) {
	if err != nil {
		l := parseMessages(ctx, err, showTraces, messages...)
		l.Panic(err)
	}
}

// Fatal uses ctx to extract information about the transaction in order to log it.
// Messages are logged in Settings.MessageKey (default is "message").
func Fatal(ctx context.Context, err error, messages ...interface{}) {
	l := parseMessages(ctx, err, showTraces, messages...)
	l.Fatal(err)
}

// Error uses ctx to extract information about the transaction in order to log it.
// Messages are logged in Settings.MessageKey (default is "message").
func Error(ctx context.Context, err error, messages ...interface{}) {
	l := parseMessages(ctx, err, showTraces, messages...)
	l.Error(err)
}

// Warning uses ctx to extract information about the transaction in order to log it.
// Messages are logged in Settings.MessageKey (default is "message").
func Warning(ctx context.Context, messages ...interface{}) {
	l := parseMessages(ctx, nil, showTraces, messages...)
	l.Warn()
}

// Info uses ctx to extract information about the transaction in order to log it.
// Messages are logged in Settings.MessageKey (default is "message").
func Info(ctx context.Context, messages ...interface{}) {
	l := parseMessages(ctx, nil, false, messages...)
	l.Info()
}

// Debug uses ctx to extract information about the transaction in order to log it.
// Messages are logged in Settings.MessageKey (default is "message").
func Debug(ctx context.Context, messages ...interface{}) {
	l := parseMessages(ctx, nil, false, messages...)
	l.Debug()
}

// Trace uses ctx to extract information about the transaction in order to log it.
// Messages are logged in Settings.MessageKey (default is "message").
func Trace(ctx context.Context, messages ...interface{}) {
	l := parseMessages(ctx, nil, false, messages...)
	l.Trace()
}

func findTrace() []StackTrace {
	traces := make([]StackTrace, 0)

	for level := 1; level < 100; level++ {
		fpcs := make([]uintptr, 1)

		// Skip levels to find the trace route
		n := runtime.Callers(level, fpcs)
		if n == 0 {
			break
		}

		caller := runtime.FuncForPC(fpcs[0] - 1)
		if caller == nil {
			break
		}

		file, line := caller.FileLine(fpcs[0] - 1)
		function := caller.Name()

		traces = append(traces, StackTrace{File: file, Line: line, Function: function})
	}

	return traces
}

func parseMessages(ctx context.Context, err error, showStackTrace bool, messages ...interface{}) *logrus.Entry {
	if ctx == nil {
		ctx = context.Background()
	}

	transactionID := ctx.Value(Settings.TransactionKey)
	logInfo := ctx.Value(Settings.LogInfoKey)
	identifier := ctx.Value(Settings.IdentifierKey)

	l := logger.WithFields(
		logrus.Fields{
			string(Settings.TransactionKey): transactionID,
			string(Settings.LogInfoKey):     logInfo,
			string(Settings.IdentifierKey):  identifier,
		},
	)

	for i, m := range messages {
		if e, ok := m.(error); ok {
			messages[i] = e.Error()
		}
	}

	if len(messages) > 0 {
		l = l.WithField(string(Settings.MessageKey), messages)

		if showStackTrace {
			l = l.WithField(string(Settings.TraceKey), findTrace())
		}
	}

	if err != nil {
		l = l.WithField(string(Settings.ErrorKey), err.Error())
	}

	return l
}
