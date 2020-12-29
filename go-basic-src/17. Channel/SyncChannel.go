package main

import (
	"fmt"
	"time"
)

func main() {
	boolChan := make(chan bool)

	go func() {
		for i := 0; i < 20; i++ {
			boolChan <- true
		}
		fmt.Println("송신 루틴 완료")
	}()

	go func() {
		for i := 0; i < 20; i++ {
			fmt.Println("수신한 데이터 : ", <-boolChan)
		}
	}()

	time.Sleep(time.Second * 3)
}
