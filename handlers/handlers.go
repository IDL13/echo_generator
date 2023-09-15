package handlers

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func Help() {
	help, err := os.ReadFile("../handlers/help.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(help))
}

func CreateNewEchoServer(n string) {
	err := os.Mkdir("cmd", 0777)
	if err != nil {
		panic(err)
	}

	err = os.Mkdir("cmd/app", 0777)
	if err != nil {
		panic(err)
	}

	err = os.Mkdir("internal", 0777)
	if err != nil {
		panic(err)
	}

	err = os.Mkdir("internal/app", 0777)
	if err != nil {
		panic(err)
	}

	err = os.Mkdir("internal/handlers", 0777)
	if err != nil {
		panic(err)
	}

	err = os.Mkdir("pkg", 0777)
	if err != nil {
		panic(err)
	}

	f, err := os.OpenFile("cmd/app/main.go", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()

	_, err = f.WriteString(fmt.Sprintf(Cmd_app, n))
	if err != nil {
		fmt.Println(err)
	}

	f, err = os.OpenFile("internal/app/app.go", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()

	_, err = f.WriteString(fmt.Sprintf(Internal_app, n))
	if err != nil {
		fmt.Println(err)
	}

	f, err = os.OpenFile("internal/handlers/handler.go", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()

	_, err = f.WriteString(Internal_handlers)
	if err != nil {
		fmt.Println(err)
	}
}

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
