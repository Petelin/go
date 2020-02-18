package util

import (
    "github.com/sirupsen/logrus"
)

func PanicIfErr(err error, args ...interface{}) {
    if err != nil {
        defer func() {
            panic(err)
        }()
        logrus.WithError(err).Warn(args...)
    }
}

func PanicIfErrf(err error, format string, args ...interface{}) {
    if err != nil {
        defer func() {
            panic(err)
        }()
        logrus.WithError(err).Warnf(format, args...)
    }
}


func LogIfError(err error, args ...interface{}) {
    if err != nil {
        logrus.WithError(err).Warn(args...)
    }
}

func LogIfErrorf(err error, format string, args ...interface{}) {
    if err != nil {
        logrus.WithError(err).Warnf(format, args...)
    }
}
