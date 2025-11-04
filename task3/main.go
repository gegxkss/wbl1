package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		os.Exit(1)
	}

	numWorkers, err := strconv.Atoi(os.Args[1])

	if err != nil || numWorkers < 2 {
		os.Exit(1)
	}

	dataChannel := make(chan string)
	// запуск воркеров
	for i := 0; i < numWorkers; i++ {
		go work(i, dataChannel)
	}
	//реализация постоянной записи данных, но с ограничением в 100 сообщений
	for count := 1; count < 100; count++ {
		message := fmt.Sprintf("Coобщение %d", count)
		dataChannel <- message
		count++
	}
	close(dataChannel)

}

func work(id int, dataChannel <-chan string) {
	for data := range dataChannel {
		fmt.Printf("Воркер номер %d: %s\n", id, data)
	}
}
