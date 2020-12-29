package main

import (
	"fmt"
	"time"
)

func main() {
	strChan := make(chan string)
	var str string

	go sendMessage(strChan)

	for i := 10; i > 0; i-- {
		select {
		case str = <-strChan:
			fmt.Printf("%s 메시지가 발송되었습니다.\n", str)
			break
		default:
			fmt.Printf("%d초 안에 메시지를 입력하시오.\n", i)
			break
		}
		if str != "" {
			break
		}
		time.Sleep(time.Second)
	}
}

func sendMessage(c chan string) {
	var message string
	fmt.Scanln(&message)

	c <- message
}
