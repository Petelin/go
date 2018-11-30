package util

import "fmt"
import "github.com/kataras/golog"

func PanicIfErr(err error, args ...interface{})  {
	if err != nil {
		golog.Fatal(args...)
	}
}


func PanicIfErrf(err error, args ...interface{})  {
	if err != nil {
		golog.Fatalf(args[0].(string), args[1:]...)
	}
}

func LogIfError(err error, args ...interface{})  {
	if err != nil{
		golog.Error(fmt.Sprintf("err=%s,", err) +fmt.Sprint( args...))
	}
}

func LogIfErrorf(err error, args... interface{})  {
	if err != nil{
		if err != nil{
			golog.Error(fmt.Sprintf("err=%s,", err) + fmt.Sprintf(args[0].(string), args[1:]...))
		}
	}
}