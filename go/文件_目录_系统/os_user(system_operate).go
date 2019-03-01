package main

import (
	"fmt"
	"os"
	"os/user"
	"time"
)

func main() {
	//fmt.Println("TempDir:", os.TempDir())

	pwd, err := os.Getwd()
	if err != nil {		panic(err)	}
	fmt.Println("user current pwd:",pwd)

	// get the current user
	usr, err := user.Current()
	if err != nil {		panic(err)	}
	homePath := usr.HomeDir
	fmt.Println("current user is:",usr.Name)
	fmt.Println("current user's homework is:",homePath)

	// change the directory
	if err := os.Chdir(homePath); err != nil {
		panic(err)
	}

	// get current working directory
	if twd, err := os.Getwd(); err != nil {
		panic(err)
	} else {
		fmt.Println("home pwd:", twd)
	}
	// home pwd: /home/ubuntu

	fpath := "temp.txt"

	// this errors for non-existing file.
	if err := os.Remove(fpath); err != nil {
		// panic(err)
		fmt.Println(err)
	}

	// this errors for non-existing file.
	file, err := os.Open(fpath)
	if err != nil {
		// panic(err)
		fmt.Println(err)
		// open temp.txt: no such file or directory

		// THIS DOES NOT GET CALLED
		// BECAUSE the process will be killed
		// in the next lines
		defer func() {
			fmt.Println("Deleting", fpath)
			if err := os.Remove(fpath); err != nil {
				// panic(err)
				fmt.Println(err)
			}
		}()

		fmt.Println("Creating", fpath)
		file, err = os.Create(fpath)
		if err != nil {
			panic(err)
		}
		fmt.Println(file)
	}

	fmt.Println("Deleting", fpath)
	if err := os.Remove(fpath); err != nil {
		// panic(err)
		fmt.Println(err)
	}

	// get process id
	pid := os.Getppid()
	fmt.Println("pid:", pid)

	// find the process
	p, err := os.FindProcess(pid)
	if err != nil {
		panic(err)
	}
	go func() {
		fmt.Println("goroutine: killing the process in 1 second...")
		time.Sleep(time.Second)
		if err := p.Kill(); err != nil {
			panic(err)
		}
	}()
	// fmt.Println("Sleeping 1,000 hours in main function...")
	// time.Sleep(1000 * time.Hour)
}
