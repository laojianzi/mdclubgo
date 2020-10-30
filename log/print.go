package log

// print func
var (
	Info   func(string, ...interface{})
	Debug  func(string, ...interface{})
	Warn   func(string, ...interface{})
	Error  func(string, ...interface{})
	DPanic func(string, ...interface{})
	Panic  func(string, ...interface{})
	Fatal  func(string, ...interface{})
)

// Printer print func interface list
type Printer interface {
	Infof(string, ...interface{})
	Debugf(string, ...interface{})
	Warnf(string, ...interface{})
	Errorf(string, ...interface{})
	DPanicf(string, ...interface{})
	Panicf(string, ...interface{})
	Fatalf(string, ...interface{})
}

func initPrinter(p Printer) {
	Info = p.Infof
	Debug = p.Debugf
	Warn = p.Warnf
	Error = p.Errorf
	DPanic = p.DPanicf
	Panic = p.Panicf
	Fatal = p.Fatalf
}
