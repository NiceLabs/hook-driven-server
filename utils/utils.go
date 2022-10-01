package utils

import (
	"fmt"
	"os/exec"
)

func AddEnv(cmd *exec.Cmd, envs map[string]string) {
	for name, value := range envs {
		cmd.Env = append(cmd.Env, fmt.Sprintf("%s=%s", name, value))
	}
	return
}
