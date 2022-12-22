package main

import (
	//"fmt"
	"os"
	"os/exec"
)

func main() {
	topCommand := exec.Command("top", "-d 0.5", "-b", "-n 5")
	awkCommand := exec.Command("grep", "bash")

	awkCommand.Stdin, _ = topCommand.StdoutPipe()
	awkCommand.Stdout = os.Stdout

	_ = awkCommand.Start()
	_ = topCommand.Run()
	_ = awkCommand.Wait()

	//fmt.Println(awkCommand)
}
