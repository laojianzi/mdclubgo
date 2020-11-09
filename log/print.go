package log

// print func
var (
	Debug  func(string, ...interface{})
	Info   func(string, ...interface{})
	Warn   func(string, ...interface{})
	Error  func(string, ...interface{})
	DPanic func(string, ...interface{})
	Panic  func(string, ...interface{})
	Fatal  func(string, ...interface{})
)

// Printer print func interface list
type Printer interface {
	Debugf(string, ...interface{})
	Infof(string, ...interface{})
	Warnf(string, ...interface{})
	Errorf(string, ...interface{})
	DPanicf(string, ...interface{})
	Panicf(string, ...interface{})
	Fatalf(string, ...interface{})
}

func initPrinter(p Printer) {
	Debug = p.Debugf
	Info = p.Infof
	Warn = p.Warnf
	Error = p.Errorf
	DPanic = p.DPanicf
	Panic = p.Panicf
	Fatal = p.Fatalf
}
