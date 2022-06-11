package logger

type myKey string

// Settings holds the settings for the logger.
var Settings = struct {
	TransactionKey myKey
	LogInfoKey     myKey
	MessageKey     myKey
	ErrorKey       myKey
	TraceKey       myKey
	IdentifierKey  myKey
}{
	TransactionKey: "txID",
	LogInfoKey:     "logInfo",
	MessageKey:     "message",
	ErrorKey:       "error",
	TraceKey:       "trace",
	IdentifierKey:  "identifier",
}
