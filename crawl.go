package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		cmnd := exec.Command("gau", scanner.Text()+" --threads 10 --o "+os.Args[2])
		cmnd2 := exec.Command("gau", scanner.Text()+" | wsl ./massurl -r -v -n 1 >> 2"+os.Args[2])
		//cmnd.Run() // and wait
		cmnd.Start()
		fmt.Println(cmnd)
		fmt.Println(cmnd2)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
