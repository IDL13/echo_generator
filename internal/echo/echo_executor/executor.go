package echo_executor

import (
	_const "codegenerator/internal/echo/const"
	"fmt"
	"os"
)

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
	err = os.Mkdir("internal/internal", 0777)
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
	_, err = f.WriteString(fmt.Sprintf(_const.CmdApp, n))
	if err != nil {
		fmt.Println(err)
	}
	f, err = os.OpenFile("internal/app/app.go", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		fmt.Println(err)
	}
	_, err = f.WriteString(fmt.Sprintf(_const.InternalApp, n))
	if err != nil {
		fmt.Println(err)
	}
	f, err = os.OpenFile("internal/internal/handler.go", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		fmt.Println(err)
	}
	_, err = f.WriteString(_const.InternalHandlers)
	if err != nil {
		fmt.Println(err)
	}
}
