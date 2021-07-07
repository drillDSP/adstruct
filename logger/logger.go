package ad_struct_logger

import "log"

type AdStructLogger interface {
	Info(args ...interface{})
}

func SetAdStructLogger(l AdStructLogger) {
	DefaultLogger = l
}

var DefaultLogger AdStructLogger = &DefaultLog{}

type DefaultLog struct {
}

func (l *DefaultLog) Info(args ...interface{}) {
	log.Println(args)
}
