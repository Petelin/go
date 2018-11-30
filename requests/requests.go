package requests

import (
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/thedevsaddam/gojsonq"

	"github.com/kataras/golog"
)

var Client *http.Client

func init() {
	Client = http.DefaultClient
}

func Do(r *http.Request) (find func(string) interface{}, err error) {
	resp, err := Client.Do(r)
	if err != nil {
		golog.Error("return error", err)
		return
	}
	if resp.StatusCode >= 400 {
		golog.Error("status code >400")
		err = errors.New("status code >= 400")
		return
	}

	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		golog.Error("read body err", err)
		return
	}
	defer resp.Body.Close()
	find = gojsonq.New().JSONString(string(bs)).Find
	return
}
