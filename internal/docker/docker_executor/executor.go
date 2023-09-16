package docker_executor

import (
	__const "codegenerator/internal/docker/const"
	"fmt"
	"os"
)

func CreateNewDockerDependencies() {
	f, err := os.OpenFile("./Dockerfile", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		fmt.Println(err)
	}
	_, err = f.WriteString(fmt.Sprintf(__const.DockerfileBase))
	if err != nil {
		fmt.Println(err)
	}
	f, err = os.OpenFile("./docker-compose.yaml", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		fmt.Println(err)
	}
	_, err = f.WriteString(fmt.Sprintf(__const.DockerComposeBase))
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
}
