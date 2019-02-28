package main

import (
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
)

func main() {
	addr := os.Getenv("REQUESTY_HOST")

	if addr == "" {
		// minikube
		addr = "http://192.168.99.100:31212"
	}

	foo := func() error {
		req := &http.Request{
			URL: &url.URL{
				Host: addr,
				Path: "/handler",
				Scheme: "http",
			},
		}
		clt := &http.Client{}

		_, err := clt.Do(req)

		if err != nil {
			return err
		}

		return nil
	}

	for {
		err := foo()

		if err != nil {
			log.Println("handler err:", err)

			return
		}

		<-time.After(time.Millisecond * 100)
	}
}
