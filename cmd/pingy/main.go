package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	for i := 0; i < 4; i++ {
		go func() {
			for {
				_, err := http.Get("http://192.168.99.100:30249/handler")

				if err != nil {
					log.Println("handler err:", err)
					return
				}

				<-time.After(time.Millisecond * 200)
			}
		}()

		go func() {
			for {
				_, err := http.Get("http://192.168.99.100:30249/error")

				if err != nil {
					log.Println("handler err:", err)
					return
				}

				<-time.After(time.Millisecond * 600)
			}
		}()
	}

	time.Sleep(time.Minute * 5)
}
