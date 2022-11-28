package tasks

import (
	"bufio"
	"fmt"
	_ "lab3_SAKOD/Queue"
	queue "lab3_SAKOD/Queue"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func readNumberFromConsole() int64 {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err.Error())
	}
	input = strings.TrimSpace(input)
	intInput, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		log.Fatal(err.Error())
	}
	return intInput
}

func clearConsole() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func Task4() {
	var number int64 = 1
	var randQueue int

	fmt.Println("Введите количество очередей:")
	queueCnt := readNumberFromConsole()

	var listQueues = make([]queue.Queue, queueCnt)
	fmt.Println("Вводите номера клиентов для обслуживания, если хотите закончить введите 0:")
	for number != 0 {
		fmt.Println("Клиент №:")
		number = readNumberFromConsole()

		if number != 0 {
			listQueues[rand.Intn(int(queueCnt))].Enqueue(int(number))

			for i := 0; i < len(listQueues); i++ {

				fmt.Printf("Очередь №%v", i+1)
				fmt.Println(listQueues[i].Elements)
			}
		}
	}

	clearConsole()

	for len(listQueues) > 0 {
		randQueue = rand.Intn(len(listQueues))
		if listQueues[randQueue].Len == 0 {
			listQueues = append(listQueues[:randQueue], listQueues[randQueue+1:]...)
		} else {
			fmt.Printf("Обслужен клиент №%v \n", listQueues[randQueue].Dequeue())
			for i := 0; i < len(listQueues); i++ {

				fmt.Printf("Очередь №%v", i+1)
				fmt.Println(listQueues[i].Elements)
			}
		}
	}

}
