package restaurants

type Logger interface {
	Debug(args ...interface{})
	Info(args ...interface{})
	Infow(msg string, KeysAndValues ...interface{})
	Error(args ...interface{})
	Errorw(msg string, KeysAndValues ...interface{})
}

var logger Logger

func SetLogger(cfgLogger Logger) {
	logger = cfgLogger
}
