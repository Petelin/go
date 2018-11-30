package util

import (
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/kataras/golog"
)

var notiChan chan string
var initOnce = new(sync.Once)

func Noti(body string) {
	select {
	case notiChan <- body:
	default:
		golog.Warn("Noti chan is full")
	}
}
func InitNoti(url string) {
	initOnce.Do(func() {
		notiChan = make(chan string, 10)
		go func() {
			for {
				select {
				case body := <-notiChan:
					beijing := time.FixedZone("Beijing Time", int((8 * time.Hour).Seconds()))
					n := time.Now().In(beijing).Format(time.StampMilli)
					//.Format(time.StampMilli)
					data := fmt.Sprintf(`{"text": "%s:\n%s"}`, n, body)
					for i := 0; i < 3; i++ {
						_, err := http.Post(url, "application/json", strings.NewReader(data))
						if err == nil {
							break
						} else {
							golog.Error("post url err: ", err)
						}
					}
				}
			}
		}()
	})
}
