package git_executor

import (
	"os"
	"os/exec"
)

func ConnectGitRepository(g string) {
	ConnectEnv := func(cmd *exec.Cmd) {
		cmd.Stdout = os.Stdout
		cmd.Stdin = os.Stdin
		cmd.Stderr = os.Stderr
		cmd.Run()
	}

	cmd := exec.Command("git", "init")
	ConnectEnv(cmd)

	cmd = exec.Command("git", "add", "*")
	ConnectEnv(cmd)

	cmd = exec.Command("git", "commit", "-m", "first commit")
	ConnectEnv(cmd)

	cmd = exec.Command("git", "branch", "-M", "main")
	ConnectEnv(cmd)

	cmd = exec.Command("git", "remote", "add", "origin", "https://"+g)
	ConnectEnv(cmd)

	cmd = exec.Command("git", "push", "-u", "origin", "main")
	ConnectEnv(cmd)
}
