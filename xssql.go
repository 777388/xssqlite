package main

import (
    "log"
    "fmt"
    "os"
    "bufio"
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
        cmnd := exec.Command("dalfox", " url " + "\"" + scanner.Text() + "\"" + " https://hahwul.xss.ht")
        cmnd2 := exec.Command("python", " sqlmap/sqlmap.py -u " + "\"" + scanner.Text() + "\"" + " --batch -o --smart ")
    	//cmnd.Run() // and wait
    	cmnd.Start()
    	fmt.Println(cmnd)
        fmt.Println(cmnd2)
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
    
}