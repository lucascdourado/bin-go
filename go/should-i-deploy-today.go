package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("bash", "-c", "curl --silent https://shouldideploy.today/api?tz=America/Sao_Paulo | jq .message")
	shouldideploytoday, err := cmd.Output()
	fmt.Println("")
	fmt.Println(string(shouldideploytoday))

	if err != nil {
		fmt.Println(err.Error())
		return
	}

}
