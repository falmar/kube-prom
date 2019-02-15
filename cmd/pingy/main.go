package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	addr := "http://192.168.99.100:31212"

	for i := 0; i < 30; i++ {
		go func() {
			for {
				_, err := http.Get(addr + "/handler")

				if err != nil {
					log.Println("handler err:", err)
					return
				}

				<-time.After(time.Millisecond * 100)
			}
		}()

		go func() {
			for {
				_, err := http.Get(addr + "/error")

				if err != nil {
					log.Println("handler err:", err)
					return
				}

				<-time.After(time.Millisecond * 500)
			}
		}()
	}

	time.Sleep(time.Minute * 5)
}
