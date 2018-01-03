package util

import "fmt"

func panicIfErr(err error, format string, args ...interface{})  {
	if err != nil {
		msg := fmt.Sprintf(format, args...)
		panic(msg)
	}
}
