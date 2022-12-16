package internal

type Printer interface {
	Error(error, string, ...interface{})
	Debug(string, ...interface{})
	Out(string, ...interface{})
}
