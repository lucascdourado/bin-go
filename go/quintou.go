package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("curl", "--silent", "--output", "/dev/null", "-o-", "https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.1/install.sh", "|", "bash")
	stdout, err := cmd.Output()
	fmt.Println(string(stdout))

	cmd = exec.Command("source", "$HOME/.nvm/nvm.sh")

	cmd = exec.Command("nvm", "install", "node")

	cmd = exec.Command("nvm", "install", "--lts")

	cmd = exec.Command("npm", "install", "--silent", "-g", "daledale")

	cmd = exec.Command("echo", "QUINTOU!")
	stdout, err = cmd.Output()
	fmt.Println(string(stdout))

	cmd = exec.Command("daledale")
	stdout, err = cmd.Output()
	fmt.Println(string(stdout))

	if err != nil {
		fmt.Println(err.Error())
		return
	}

}
