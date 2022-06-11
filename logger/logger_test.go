// nolint:testpackage,paralleltest // we need access to logger var for tests.
package logger

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/sirupsen/logrus"
)

var (
	errGeneric = errors.New("MY ERROR")
	errText    = errors.New("error text")
)

func resetLoggerToDefaults(t *testing.T) {
	t.Helper()

	if err := SetLogLevel("debug"); err != nil {
		t.Error("could not set log level", err)
	}

	if err := SetFormatter("json"); err != nil {
		t.Error("could not set formatter", err)
	}
}

func TestSetFormatter(t *testing.T) {
	if err := SetFormatter("json"); err != nil {
		t.Error("could not set formatter")
	}

	if err := SetFormatter("text"); err != nil {
		t.Error("could not set formatter")
	}

	if err := SetFormatter("invalid"); err == nil {
		t.Error("expected error")
	}
}

func TestSetLogLevel(t *testing.T) {
	resetLoggerToDefaults(t)

	availableLevels := map[string]logrus.Level{
		"trace":   logrus.TraceLevel,
		"dEbug":   logrus.DebugLevel,
		"INFO":    logrus.InfoLevel,
		"warN":    logrus.WarnLevel,
		"warnIng": logrus.WarnLevel,
		"error":   logrus.ErrorLevel,
		"fatal":   logrus.FatalLevel,
		"pANIC":   logrus.PanicLevel,
	}

	if err := SetLogLevel("trce"); err == nil {
		t.Error(`expected to get error while parsing level "trce"`)
	}

	if logger.GetLevel() != logrus.DebugLevel {
		t.Error("expected logger to have default debug level. Instead got", logger.GetLevel())
	}

	for lvl, expectedLvl := range availableLevels {
		if err := SetLogLevel(lvl); err != nil {
			t.Error("could not set level", lvl, err)
		}

		if logger.GetLevel() != expectedLvl {
			t.Error("Expected level ", expectedLvl, " got ", logger.GetLevel())
		}
	}
}

func TestSetOutput(t *testing.T) {
	var buf bytes.Buffer

	resetLoggerToDefaults(t)
	SetOutput(&buf)

	ctx := context.WithValue(context.TODO(), Settings.TransactionKey, "abc")

	Info(ctx, "test", struct {
		TestStruct string
	}{TestStruct: "testStruct"})

	if buf.Len() == 0 {
		t.Error(`buffer seems empty. Should contain the log`)
	}
}

func TestSetLogTrace(t *testing.T) {
	var buf bytes.Buffer

	resetLoggerToDefaults(t)
	SetOutput(&buf)

	ctx := context.WithValue(context.TODO(), Settings.TransactionKey, "abc")

	Info(ctx, "test", struct {
		TestStruct string
	}{TestStruct: "testStruct"})

	logResult := buf.String()

	if !strings.Contains(logResult, "txID") ||
		!strings.Contains(logResult, "abc") ||
		!strings.Contains(logResult, "message") ||
		!strings.Contains(logResult, "TestStruct") ||
		strings.Contains(logResult, "trace") {
		t.Error("unexpected result. Expected key `trace` not among the log")
	}

	resetLoggerToDefaults(t)
	SetOutput(&buf)
	SetLogTrace(true)

	Warning(ctx, "test", struct {
		TestStruct string
	}{TestStruct: "testStruct"})

	logResult = buf.String()

	if !strings.Contains(logResult, "txID") ||
		!strings.Contains(logResult, "abc") ||
		!strings.Contains(logResult, "message") ||
		!strings.Contains(logResult, "TestStruct") ||
		!strings.Contains(logResult, "trace") {
		t.Error("unexpected result. Expected key `trace` not among the log")
	}

	buf.Reset()

	Info(ctx, "test", struct {
		TestStruct string
	}{TestStruct: "testStruct"})

	logResult = buf.String()

	if !strings.Contains(logResult, "txID") ||
		!strings.Contains(logResult, "abc") ||
		!strings.Contains(logResult, "message") ||
		!strings.Contains(logResult, "TestStruct") ||
		strings.Contains(logResult, "trace") {
		t.Error("unexpected result. Expected to not find key `trace` among the log for info")
	}
}

func TestLogTrace(t *testing.T) {
	var buf bytes.Buffer

	resetLoggerToDefaults(t)
	SetOutput(&buf)

	ctx := context.WithValue(context.TODO(), Settings.TransactionKey, "abc")

	if err := SetLogLevel("trace"); err != nil {
		t.Error("could not set log level", err)
	}

	Trace(ctx, "test", struct {
		TestStruct string
	}{TestStruct: "testStruct"})

	logResult := buf.String()

	if !strings.Contains(logResult, "txID") ||
		!strings.Contains(logResult, "abc") ||
		!strings.Contains(logResult, "message") ||
		!strings.Contains(logResult, "TestStruct") ||
		!strings.Contains(logResult, "trace") {
		t.Error(`unexpected result. Expected something like:
{
	"level": "trace",
	"message": ["test", {
		"TestStruct": "testStruct"
	}],
	"msg": "test{testStruct}",
	"time": "2019-11-12T19:14:41+02:00",
	"txID": "abc"
}
but the result was:`, logResult)
	}

	buf.Reset()
	Trace(ctx, errGeneric, struct {
		TestStruct string
	}{TestStruct: "testStruct"})

	logResult = buf.String()

	if !strings.Contains(logResult, errGeneric.Error()) {
		t.Error(`unexpected result. Expected to see error but the result was:`, logResult)
	}

	if err := SetLogLevel("debug"); err != nil {
		t.Error("could not set log level", err)
	}

	buf.Reset()

	Trace(ctx, "test")

	if buf.Len() > 0 {
		t.Error("expected buffer to be empty, got: ", buf.String())
	}
}

func TestLogDebug(t *testing.T) {
	var buf bytes.Buffer

	resetLoggerToDefaults(t)
	SetOutput(&buf)

	ctx := context.WithValue(context.TODO(), Settings.TransactionKey, "abc")

	Debug(ctx, "test", struct {
		TestStruct string
	}{TestStruct: "testStruct"})

	logResult := buf.String()

	if !strings.Contains(logResult, "txID") ||
		!strings.Contains(logResult, "abc") ||
		!strings.Contains(logResult, "message") ||
		!strings.Contains(logResult, "TestStruct") ||
		!strings.Contains(logResult, "debug") {
		t.Error(`unexpected result. Expected something like:
{
	"level": "debug",
	"message": ["test", {
		"TestStruct": "testStruct"
	}],
	"msg": "test{testStruct}",
	"time": "2019-11-12T19:14:41+02:00",
	"txID": "abc"
}
but the result was:`, logResult)
	}

	buf.Reset()
	Debug(ctx, errGeneric, struct {
		TestStruct string
	}{TestStruct: "testStruct"})

	logResult = buf.String()

	if !strings.Contains(logResult, errGeneric.Error()) {
		t.Error(`unexpected result. Expected to see error but the result was:`, logResult)
	}

	if err := SetLogLevel("info"); err != nil {
		t.Error("could not set log level", err)
	}

	buf.Reset()
	Debug(ctx, "test")

	if buf.Len() > 0 {
		t.Error("expected buffer to be empty, got: ", buf.String())
	}
}

func TestLogInfo(t *testing.T) {
	var buf bytes.Buffer

	resetLoggerToDefaults(t)
	SetOutput(&buf)

	ctx := context.WithValue(context.TODO(), Settings.TransactionKey, "abc")

	Info(ctx, "test", struct {
		TestStruct string
	}{TestStruct: "testStruct"})

	logResult := buf.String()

	if !strings.Contains(logResult, "txID") ||
		!strings.Contains(logResult, "abc") ||
		!strings.Contains(logResult, "message") ||
		!strings.Contains(logResult, "TestStruct") ||
		!strings.Contains(logResult, "info") {
		t.Error(`unexpected result. Expected something like:
{
	"level": "info",
	"message": ["test", {
		"TestStruct": "testStruct"
	}],
	"msg": "test{testStruct}",
	"time": "2019-11-12T19:14:41+02:00",
	"txID": "abc"
}
but the result was:`, logResult)
	}

	buf.Reset()
	Info(ctx, errGeneric, struct {
		TestStruct string
	}{TestStruct: "testStruct"})

	logResult = buf.String()

	if !strings.Contains(logResult, errGeneric.Error()) {
		t.Error(`unexpected result. Expected to see error but the result was:`, logResult)
	}

	if err := SetLogLevel("warn"); err != nil {
		t.Error("could not set log level", err)
	}

	buf.Reset()
	Info(ctx, "test")

	if buf.Len() > 0 {
		t.Error("expected buffer to be empty, got: ", buf.String())
	}
}

func TestLogWarning(t *testing.T) {
	var buf bytes.Buffer

	resetLoggerToDefaults(t)
	SetOutput(&buf)

	ctx := context.WithValue(context.TODO(), Settings.TransactionKey, "abc")

	Warning(ctx, "test", struct {
		TestStruct string
	}{TestStruct: "testStruct"})

	logResult := buf.String()

	if !strings.Contains(logResult, "txID") ||
		!strings.Contains(logResult, "abc") ||
		!strings.Contains(logResult, "message") ||
		!strings.Contains(logResult, "TestStruct") ||
		!strings.Contains(logResult, "warning") {
		t.Error(`unexpected result. Expected something like:
{
	"level": "warn",
	"message": ["test", {
		"TestStruct": "testStruct"
	}],
	"msg": "test{testStruct}",
	"time": "2019-11-12T19:14:41+02:00",
	"txID": "abc"
}
but the result was:`, logResult)
	}

	buf.Reset()
	Warning(ctx, errGeneric, struct {
		TestStruct string
	}{TestStruct: "testStruct"})

	logResult = buf.String()

	if !strings.Contains(logResult, errGeneric.Error()) {
		t.Error(`unexpected result. Expected to see error but the result was:`, logResult)
	}

	if err := SetLogLevel("error"); err != nil {
		t.Error("could not set log level", err)
	}

	buf.Reset()
	Warning(ctx, "test")

	if buf.Len() > 0 {
		t.Error("expected buffer to be empty, got: ", buf.String())
	}
}

func TestLogError(t *testing.T) {
	var buf bytes.Buffer

	resetLoggerToDefaults(t)
	SetOutput(&buf)

	ctx := context.WithValue(context.TODO(), Settings.TransactionKey, "abc")

	Error(ctx, errText, "test", struct {
		TestStruct string
	}{TestStruct: "testStruct"})

	logResult := buf.String()

	if !strings.Contains(logResult, "txID") ||
		!strings.Contains(logResult, "abc") ||
		!strings.Contains(logResult, "message") ||
		!strings.Contains(logResult, "TestStruct") ||
		!strings.Contains(logResult, "error") ||
		!strings.Contains(logResult, "error text") {
		t.Error(`unexpected result. Expected something like:
{
	"level": "warn",
	"message": ["test", {
		"TestStruct": "testStruct"
	}],
	"msg": "test{testStruct}",
	"time": "2019-11-12T19:14:41+02:00",
	"txID": "abc"
}
but the result was:`, logResult)
	}

	buf.Reset()
	Error(ctx, errGeneric, struct {
		TestStruct string
	}{TestStruct: "testStruct"})

	logResult = buf.String()

	if !strings.Contains(logResult, errGeneric.Error()) {
		t.Error(`unexpected result. Expected to see error but the result was:`, logResult)
	}

	buf.Reset()
	Error(ctx, errText, "d")

	logResult = buf.String()

	if !strings.Contains(logResult, "error text") {
		t.Error(`unexpectd result. Expected to see error when no messages but the result was:`, logResult)
	}

	if err := SetLogLevel("fatal"); err != nil {
		t.Error("could not set log level", err)
	}

	buf.Reset()
	Error(ctx, errText, "test")

	if buf.Len() > 0 {
		t.Error("expected buffer to be empty, got: ", buf.String())
	}
}

func TestLogFatal(t *testing.T) {
	// In order to test fatal we need to pass a fake exit function
	fakeExit := func(int) {}

	logger.ExitFunc = fakeExit

	var buf bytes.Buffer

	resetLoggerToDefaults(t)
	SetOutput(&buf)

	ctx := context.WithValue(context.TODO(), Settings.TransactionKey, "abc")

	Fatal(ctx, errText, "test", struct {
		TestStruct string
	}{TestStruct: "testStruct"})

	logResult := buf.String()

	if !strings.Contains(logResult, "txID") ||
		!strings.Contains(logResult, "abc") ||
		!strings.Contains(logResult, "message") ||
		!strings.Contains(logResult, "TestStruct") ||
		!strings.Contains(logResult, "fatal") ||
		!strings.Contains(logResult, "error text") {
		t.Error(`unexpected result. Expected something like:
{
	"level": "fatal",
	"message": ["test", {
		"TestStruct": "testStruct"
	}],
	"msg": "test{testStruct}",
	"time": "2019-11-12T19:14:41+02:00",
	"txID": "abc"
}
but the result was:`, logResult)
	}

	buf.Reset()
	Fatal(ctx, errGeneric, struct {
		TestStruct string
	}{TestStruct: "testStruct"})

	logResult = buf.String()

	if !strings.Contains(logResult, errGeneric.Error()) {
		t.Error(`unexpected result. Expected to see error but the result was:`, logResult)
	}

	if err := SetLogLevel("panic"); err != nil {
		t.Error("could not set log level", err)
	}

	buf.Reset()
	Fatal(ctx, errText, "test")

	if buf.Len() > 0 {
		t.Error("expected buffer to be empty, got: ", buf.String())
	}
}

func TestLogPanic(t *testing.T) {
	var buf bytes.Buffer

	defer func() {
		if err := recover(); err != nil {
			logEntry, _ := err.(*logrus.Entry)

			if logEntry.Data[string(Settings.TransactionKey)] != "abc" ||
				!strings.Contains(fmt.Sprint(logEntry.Data[string(Settings.MessageKey)]), "testStruct") ||
				logEntry.Message != "error text" {
				t.Error("did not get the correct recover message")
			}
		}
	}()

	resetLoggerToDefaults(t)
	SetOutput(&buf)

	ctx := context.WithValue(context.TODO(), Settings.TransactionKey, "abc")

	Panic(ctx, errText, "test", struct {
		TestStruct string
	}{TestStruct: "testStruct"})

	t.Error("should not reach this line")
}
