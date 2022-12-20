package internal

type Printer interface {
	Error(error, string, ...interface{})
	Debug(string, ...interface{})
	Out(string, ...interface{})

	Done()
}

type NullPrinter struct{}

func (x NullPrinter) Error(err error, msg string, rest ...interface{}) { return }
func (x NullPrinter) Debug(msg string, rest ...interface{})            { return }
func (x NullPrinter) Out(msg string, rest ...interface{})              { return }
func (x NullPrinter) Done()                                            { return }
