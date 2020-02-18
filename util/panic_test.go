package util

import (
    "errors"
    "testing"
)

func TestLogErr(t *testing.T) {
    err := errors.New("new error")
    LogIfError(err)
    LogIfError(err, "haha")
    LogIfErrorf(err, "info:%s", "haha")
}
