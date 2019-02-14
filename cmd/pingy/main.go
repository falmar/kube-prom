package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	for i := 0; i < 30; i++ {
		go func() {
			for {
				_, err := http.Get("http://192.168.99.100:32152/handler")

				if err != nil {
					log.Println("handler err:", err)
					return
				}

				<-time.After(time.Millisecond * 100)
			}
		}()

		go func() {
			for {
				_, err := http.Get("http://192.168.99.100:32152/error")

				if err != nil {
					log.Println("handler err:", err)
					return
				}

				<-time.After(time.Millisecond * 100)
			}
		}()
	}

	time.Sleep(time.Minute * 5)
}
