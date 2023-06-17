package services

import "fmt"

func (service *ServiceGroup) parser() {
	for {
		line := <-service.rawMsgChannel
		fmt.Println("received", len(line))
	}
}
