package main

import (
	"codegenerator/internal/docker/docker_executor"
	"codegenerator/internal/echo/echo_executor"
	"codegenerator/internal/git/git_executor"
	"codegenerator/internal/helper"
	"codegenerator/pkg/parser"
	"flag"
	"fmt"
	"os"
	"os/exec"
)

type Options struct {
	flags0 string `name:"help" value:"false" usage:"help"`
	flags1 string `name:"new" value:"echo" usage:"new"`
	flags2 string `name:"git" value:"false" usage:"git"`
}

func main() {
	flagsB, flagsS := parser.NewParser(Options{})
	flag.Parse()
	if *flagsB[0] {
		helper.Help()
		os.Exit(1)
	}
	if *flagsS[0] != "default" {
		var path string
		fmt.Print("Enter the path to the server's echo folder:")
		fmt.Scan(&path)
		err := os.Mkdir(path+"/"+*flagsS[0], 0777)
		if err != nil {
			panic(err)
		}
		err = os.Chdir(path + "/" + *flagsS[0])
		if err != nil {
			panic(err)
		}
		ConnectEnv := func(cmd *exec.Cmd) {
			cmd.Stdout = os.Stdout
			cmd.Stdin = os.Stdin
			cmd.Stderr = os.Stderr
			cmd.Run()
		}
		cmd := exec.Command("go", "mod", "init", *flagsS[0])
		ConnectEnv(cmd)
		cmd = exec.Command("go", "get", "github.com/labstack/echo/v4")
		ConnectEnv(cmd)
		echo_executor.CreateNewEchoServer(*flagsS[0])
		docker_executor.CreateNewDockerDependencies()
		if *flagsB[1] {
			var g string
			fmt.Print("Path to your git repository:")
			fmt.Scan(&g)
			err = os.Chdir(path + "/" + *flagsS[0])
			if err != nil {
				panic(err)
			}
			git_executor.ConnectGitRepository(g)
		}
	}
}
