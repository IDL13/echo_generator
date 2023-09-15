package main

import (
	"codegenerator/handlers"
	"flag"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	help := flag.Bool("help", false, "help")

	new := flag.String("new", "echo", "new")

	git := flag.Bool("git", false, "git")

	flag.Parse()

	if *help {
		handlers.Help()
		os.Exit(1)
	}

	if *new != "default" {
		var path string

		fmt.Print("Enter the path to the server's echo folder:")
		fmt.Scan(&path)

		err := os.Mkdir(path+"/"+*new, 0777)
		if err != nil {
			panic(err)
		}

		err = os.Chdir(path + "/" + *new)
		if err != nil {
			panic(err)
		}

		cmdEnv := func(cmd *exec.Cmd) {
			cmd.Stdout = os.Stdout
			cmd.Stdin = os.Stdin
			cmd.Stderr = os.Stderr
			cmd.Run()
		}

		cmd := exec.Command("go", "mod", "init", *new)
		cmdEnv(cmd)

		cmd = exec.Command("go", "get", "github.com/labstack/echo/v4")
		cmdEnv(cmd)

		handlers.New(*new)

		if *git {
			var g string

			fmt.Print("Path to your git repository:")
			fmt.Scan(&g)

			err = os.Chdir(path + "/" + *new)
			if err != nil {
				panic(err)
			}

			handlers.Git(g)
		}
	}

}
