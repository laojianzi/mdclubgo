package log

// Info uses zap.SugaredLogger.Infof to log a formatd message.
func Info(format string, args ...interface{}) {
	instance.Infof(format, args...)
}

// Debug uses zap.SugaredLogger.Debugf to log a formatd message.
func Debug(format string, args ...interface{}) {
	instance.Debugf(format, args...)
}

// Warn uses zap.SugaredLogger.Warnf to log a formatd message.
func Warn(format string, args ...interface{}) {
	instance.Warnf(format, args...)
}

// Error uses zap.SugaredLogger.Errorf to log a formatd message.
func Error(format string, args ...interface{}) {
	instance.Errorf(format, args...)
}

// DPanic uses zap.SugaredLogger.DPanicf to log a formatd message. In development, the
// logger then panics. (See DPanicLevel for details.)
func DPanic(format string, args ...interface{}) {
	instance.DPanicf(format, args...)
}

// Panic uses zap.SugaredLogger.Panicf to log a formatd message, then panics.
func Panic(format string, args ...interface{}) {
	instance.Panicf(format, args...)
}

// Fatal uses zap.SugaredLogger.Fatalf to log a formatd message, then calls os.Exit.
func Fatal(format string, args ...interface{}) {
	instance.Fatalf(format, args...)
}
