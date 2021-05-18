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

var wg sync.WaitGroup

//Copia uma imagem para um nó
func copyFileToNode(filename, dest string) {

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
		if err == nil {
			break
		} else {
			fmt.Println(err.Error())
		}
	}

}

func main() {

	//read images
	files, err := ioutil.ReadDir("./masterInput")
	if err != nil {
		log.Fatal(err)
	}

	nodeNumber := 1
	wg.Add(len(files))
	for _, f := range files {

		filename := "./masterInput/" + f.Name()

		if nodeNumber == 4 {
			nodeNumber = 1
		}

		dest := "172.42.42.10" + strconv.Itoa(nodeNumber)

		nodeNumber++

		go copyFileToNode(filename, dest)

	}
	fmt.Printf("Numero de goroutines: %d\n", runtime.NumGoroutine())

	for {
		if runtime.NumGoroutine() == 1 {
			//time.Sleep(time.Millisecond)
			wg.Wait()
			break
		}
	}
	fmt.Println("Finalizado!")
}
