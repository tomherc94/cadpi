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
func copyFileToNode(filename, dest string, wg *sync.WaitGroup) {

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
			break
		}
		err = cmd.Wait()
	}

}

func main() {

	var wg sync.WaitGroup

	//read images
	files, err := ioutil.ReadDir("./masterInput")
	if err != nil {
		log.Fatal(err)
	}

	nodeNumber := 1

	for _, f := range files {
		wg.Add(1)

		filename := "./masterInput/" + f.Name()

		if nodeNumber == 4 {
			nodeNumber = 1
		}

		dest := "172.42.42.10" + strconv.Itoa(nodeNumber)

		nodeNumber++

		go copyFileToNode(filename, dest, &wg)

	}
	fmt.Printf("Numero de goroutines: %d\n", runtime.NumGoroutine())

	wg.Wait()

	fmt.Println("Finalizado!")
}
