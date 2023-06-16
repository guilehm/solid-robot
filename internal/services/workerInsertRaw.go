package services

import "fmt"

func (service *ServiceGroup) workerInsertRaw() {
	for {
		line := <-service.rawMsgChannel
		fmt.Println("received", len(line))
	}
}
