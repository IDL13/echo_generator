package docker_executor

import (
	_const "codegenerator/internal/docker/const"
	logger "codegenerator/pkg/logger"
	"fmt"
	"os"
)

func CreateNewDockerDependencies() {
	logg := logger.NewLogger("../../log.txt")

	f, err := os.OpenFile("./Dockerfile", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		logg.Write("Error from docker_executor file", err)
	}

	_, err = f.WriteString(fmt.Sprintf(_const.DockerfileBase))
	if err != nil {
		logg.Write("Error from docker_executor file", err)
	}

	f, err = os.OpenFile("./docker-compose.yaml", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		logg.Write("Error from docker_executor file", err)
	}

	_, err = f.WriteString(fmt.Sprintf(_const.DockerComposeBase))
	if err != nil {
		logg.Write("Error from docker_executor file", err)
	}

	f, err = os.OpenFile("./.dockerignore.yaml", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		logg.Write("Error from docker_executor file", err)
	}

	defer f.Close()
}
