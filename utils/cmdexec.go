package utils

import (
	"fmt"
	"os/exec"
)

func CmdExecutor(prefix string, args ...string) (string, error) {
	cmd := exec.Command(prefix, args...)
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	return string(stdout), nil
}
