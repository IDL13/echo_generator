package logger

import (
	"fmt"
	"log"
	"os"
	"sync"
)

type Logger struct {
	mutex *sync.Mutex
	wg    sync.WaitGroup
	path  string
	file  *os.File
}

func NewLogger(p string) *Logger {
	f, err := os.OpenFile(p, os.O_CREATE, 0777)
	if err != nil {
		log.Fatal("Error from NewLogger constructor")
	}

	defer f.Close()

	return &Logger{
		mutex: &sync.Mutex{},
		file:  f,
		path:  p,
	}
}

func (l *Logger) Write(str string, inputError error) {
	file, err := os.OpenFile(l.path, os.O_APPEND|os.O_WRONLY, 0777)
	if err != nil {
		panic(err)
	}

	l.wg.Add(1)

	go func() {
		l.mutex.Lock()

		errString := fmt.Sprintf(str+"\t"+"%v"+"\n", inputError)
		_, err := file.Write([]byte(errString))
		if err != nil {
			panic(err)
		}

		l.mutex.Unlock()
		file.Close()
		l.wg.Done()
	}()

	l.wg.Wait()
}

func (l *Logger) Delete() {
	os.Remove(l.path)
}

func (l *Logger) Err(str string) {
	_, err := os.Stderr.WriteString(str)
	if err != nil {
		panic(err)
	}
}
