package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"strconv"
	"sync"
)

var wg sync.WaitGroup

func enableSSH() {
	cmd := exec.Command("sudo", "ssh-keygen")

	//fmt.Println(cmd.String())
	err := cmd.Run()

	if err != nil {
		println(err.Error())
		return
	}

	for i := 1; i <= 3; i++ {
		cmd = exec.Command("sudo", "ssh-copy-id", "172.42.42.10"+strconv.Itoa(i))

		//fmt.Println(cmd.String())
		err = cmd.Run()

		if err != nil {
			println(err.Error())
			return
		}
	}

}

//Copia uma imagem para um nó
func copyFileToNode(filename, dest string) {

	defer wg.Done()

	//scp arquivo user@ip_destino:/home/user/
	//scp sourceuser@sourcehost:/path/to/source/file destinationuser@destinationhost:/path/to/destination/
	cmd := exec.Command("scp", filename, "root@"+dest+":/home/vagrant/workerInput")

	//fmt.Println(cmd.String())
	err := cmd.Run()

	if err != nil {
		println(err.Error())
		return
	}
}

func main() {

	/*var workersIps []string

	// open the file
	ips, err := os.Open("ips.txt")

	//handle errors while opening
	if err != nil {
		log.Fatalf("Error when opening file: %s", err)
	}

	fileScanner := bufio.NewScanner(ips)

	// read line by line
	for fileScanner.Scan() {
		workersIps = append(workersIps, fileScanner.Text())
	}
	// handle first encountered error while reading
	if err := fileScanner.Err(); err != nil {
		log.Fatalf("Error while reading file: %s", err)
	}

	ips.Close()*/

	//enableSSH()

	//read images
	files, err := ioutil.ReadDir("./masterInput")
	if err != nil {
		log.Fatal(err)
	}

	nodeNumber := 1
	//qtdDest := len(workersIps)

	for _, f := range files {
		wg.Add(1)
		filename := "./masterInput/" + f.Name()

		if nodeNumber == 4 {
			nodeNumber = 1
		}

		dest := "172.42.42.10" + strconv.Itoa(nodeNumber)

		nodeNumber++

		go copyFileToNode(filename, dest)
	}
	fmt.Println("Finalizado!")
	wg.Wait()

}
