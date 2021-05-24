package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"runtime"
	"strconv"
	"sync"
	"time"
)

//Copia uma imagem para um nó
func copyFileToNode(filename, dest string, wg *sync.WaitGroup, channel chan int, nodeNumber int) {

	defer wg.Done()

	//sshpass -p '123' scp -o StrictHostKeyChecking=no ./masterInput/image_3.jpg vagrant@172.42.42.103:/home/vagrant/workerInput

	arg0 := "sshpass"
	arg1 := "-p"
	arg2 := "123"
	arg3 := "scp"
	arg4 := "-o"
	arg5 := "StrictHostKeyChecking=no"
	arg6 := filename
	arg7 := "root@" + dest + ":/home/vagrant/workerInput"

	cmd := exec.Command(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)

	//do-while in golang :'(
	for {
		err := cmd.Start()
		if err != nil {
			log.Fatal(err)
		} else {
			err = cmd.Wait()
			channel <- nodeNumber
			break
		}
		err = cmd.Wait()
	}

}

func main() {

	channel := make(chan int, 6)

	channel <- 1
	channel <- 2
	channel <- 3

	var wg sync.WaitGroup

	//read images
	files, err := ioutil.ReadDir("./masterInput")
	if err != nil {
		log.Fatal(err)
	}

	var nodeNumber int

	//exercutar as goroutines de acordo com o buffer channel
	for _, f := range files {
		wg.Add(1)

		filename := "./masterInput/" + f.Name()

		nodeNumber = <-channel

		dest := "172.42.42.10" + strconv.Itoa(nodeNumber)

		fmt.Println(filename + " -> " + dest)
		go copyFileToNode(filename, dest, &wg, channel, nodeNumber)

	}

	wg.Wait()
	fmt.Printf("Numero de goroutines: %d\n", runtime.NumGoroutine())

	//executar remotamente aplicativo JAVA em cada Worker
	for i := 1; i <= 3; i++ {
		dest := "172.42.42.10" + strconv.Itoa(i)

		arg0 := "sshpass"
		arg1 := "-p"
		arg2 := "123"

		arg3 := "ssh"
		arg4 := "-o"
		arg5 := "StrictHostKeyChecking=no"
		arg6 := "root@" + dest
		arg7 := "'java -jar -Xmx1024m workerApp.jar'"

		cmd := exec.Command(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)

		err := cmd.Start()
		if err != nil {
			log.Fatal(err)
		}
		err = cmd.Wait()
		time.Sleep(3 * time.Second)

	}

	fmt.Println("Finalizado!")
}
