package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Println("pwd:", pwd)

	lsCmd := exec.Command("ls", "-a")

	lsOutput1, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(lsOutput1))

	if err := exec.Command("touch", "temp.txt").Run(); err != nil {
		panic(err)
	}

	// lsOutput2, err := lsCmd.Output()
	// if err != nil {
	// 	panic(err)
	// }
	// panic: exec: Stdout already set
	// SHOULD NOT REUSE IT

	lsCmd2 := exec.Command("ls", "-a")
	lsOutput2, err := lsCmd2.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(lsOutput2))

	if err := os.Remove("temp.txt"); err != nil {
		panic(err)
	}

	lsCmd3 := exec.Command("ls", "-a")
	lsOutput3, err := lsCmd3.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(lsOutput3))

	// http://golang.org/pkg/os/exec/#example_Cmd_Start
	slCmd := exec.Command("sleep", "1")
	if err := slCmd.Start(); err != nil {
		panic(err)
	}
	fmt.Println("Waiting for command to finish...")
	err = slCmd.Wait()
	fmt.Println("Command finished with error:", err)

	// http://golang.org/pkg/os/exec/#Cmd.Output
	cmd := exec.Command("echo", "-n", `{"Name": "Bob", "Age": 32}`)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}
	if err := cmd.Start(); err != nil {
		panic(err)
	}
	var person struct {
		Name string
		Age  int
	}
	if err := json.NewDecoder(stdout).Decode(&person); err != nil {
		panic(err)
	}
	if err := cmd.Wait(); err != nil {
		panic(err)
	}
	fmt.Printf("%s is %d years old\n", person.Name, person.Age)
	// Bob is 32 years old
}
