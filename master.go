package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"runtime"
	"strconv"
	"sync"
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
		err := cmd.Run()
		if err != nil {
			fmt.Println(err)
		} else {
			//err = cmd.Wait()
			channel <- nodeNumber
			break
		}
		//err = cmd.Wait()
	}

}

func workerApp(dest string, wgJava *sync.WaitGroup, channelJava chan int, nodeNumber int) {
	defer wgJava.Done()

	arg0 := "sshpass"
	arg1 := "-p"
	arg2 := "123"

	arg3 := "/usr/bin/ssh"
	//arg4 := "-o"
	//arg5 := "StrictHostKeyChecking=no"
	arg6 := "root@" + dest
	arg7 := "/home/vagrant/executeWorkerApp.sh"

	cmd := exec.Command(arg0, arg1, arg2, arg3, arg6, arg7)
	//fmt.Println(cmd.String())

	errCmd := cmd.Run()

	if errCmd != nil {
		fmt.Println(errCmd)
	}
	channelJava <- nodeNumber
	fmt.Printf("WorkerApp.jar do worker%d finalizado!\n\n", nodeNumber)

	//err = cmd.Wait()

}

func workerCopy(dest string, wgCopy *sync.WaitGroup, channelCopy chan int) {
	defer wgCopy.Done()

	arg0 := "sshpass"
	arg1 := "-p"
	arg2 := "123"

	arg3 := "/usr/bin/ssh"
	//arg4 := "-o"
	//arg5 := "StrictHostKeyChecking=no"
	arg6 := "root@" + dest
	arg7 := "/home/vagrant/executeWorkerCopy.sh"

	cmd := exec.Command(arg0, arg1, arg2, arg3, arg6, arg7)
	//fmt.Println(cmd.String())

	errCmd := cmd.Run()

	if errCmd != nil {
		fmt.Println(errCmd)
	}
	channelCopy <- 1

	//err = cmd.Wait()

}

func main() {

	channel := make(chan int, 6)

	channel <- 1
	channel <- 2
	channel <- 3

	var wg, wgJava, wgCopy sync.WaitGroup

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

	channelJava := make(chan int, 6)

	channelJava <- 1
	channelJava <- 2
	channelJava <- 3

	//executar workerApp.jar em cada Worker
	for {
		wgJava.Add(1)
		i := <-channelJava
		dest := "172.42.42.10" + strconv.Itoa(i)

		fmt.Printf("Executando workerApp.jar no worker%d\n\n", i)
		go workerApp(dest, &wgJava, channelJava, i)
		fmt.Printf("Numero de goroutines: %d\n", runtime.NumGoroutine())

		if i == 3 {
			break
		}
	}
	wgJava.Wait()

	channelCopy := make(chan int)
	channelCopy <- 1

	//executar workerCopy.jar em cada Worker
	for j := 1; j <= 3; j++ {
		wgCopy.Add(1)
		fmt.Println("Aguardando término de processamento ...")
		i := <-channelJava
		copyOk := <-channelCopy
		dest := "172.42.42.10" + strconv.Itoa(i)

		fmt.Printf("Copiando arquivos de worker%d\n\n", i)
		if copyOk == 1 {
			go workerCopy(dest, &wgCopy, channelCopy)
		}

	}
	wgCopy.Wait()

	fmt.Println("Finalizado!")
}
