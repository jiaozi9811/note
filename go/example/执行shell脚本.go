package main

import (
	"fmt"
	"os/exec"
	"sync"
	"time"
)

func execsh(s string) (result string, err error) {
	wg := new(sync.WaitGroup)

	if len(s) == 0 {
		return "", fmt.Errorf("shell is nil")
	}

	comm := exec.Command("/bin/bash", "-c", s)

	wg.Add(1)
	outp, err := comm.Output()
	if err != nil {
		fmt.Println("commend run fail")
		return "", err
	}
	wg.Done()

	result = string(outp)
	wg.Wait()
	return
}

func main() {
	commands := []string{"echo newline1", "echo newline2", "echo newline4"}
	for _, str := range commands {
		go fmt.Println(execsh(str))
	}
	time.Sleep(time.Second)
}
