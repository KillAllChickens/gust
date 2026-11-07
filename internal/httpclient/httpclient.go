package httpclient

import (
	"sync"
	"time"

	"resty.dev/v3"
)



var (
	client *resty.Client
	once sync.Once
)


func Get() *resty.Client {
	once.Do(func () {
		client = resty.New().
			SetTimeout(10 * time.Second).
			SetHeader("User-Agent", "gust/1.0 (admin@killallchickens.org)").
			SetHeader("Accept-Encoding", "gzip")
	})
	return client
}
