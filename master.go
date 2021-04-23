package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"sync"
)

var wg sync.WaitGroup

//Copia uma imagem para um nó
func copyFileToNode(filename, dest string) {

	defer wg.Done()

	cmd := exec.Command("cp", filename, dest)

	fmt.Println(cmd.String())
	err := cmd.Run()

	if err != nil {
		println(err.Error())
		return
	}
}

func main() {
	files, err := ioutil.ReadDir("./masterInput")
	if err != nil {
		log.Fatal(err)
	}

	//nodeNumber := 1

	for _, f := range files {
		wg.Add(1)
		//filename := "./masterInput/" + f.Name()
		//dest := "./masterOutput"
		go copyFileToNode("./masterInput/"+f.Name(), "./masterOutput")

	}
	wg.Wait()

}
