new:
	- go run ./cmd/main.go -new "$(FOLDER)"

build_windows:
	- go build -o $(NAME).exe ./cmd/main.go

build_linux:
	- go build -o $(NAME).sh ./cmd/main.go

new-git:
	- go run ./cmd/main.go -new "$(FOLDER)" -git "$(LINK)"

help:
	- go run ./cmd/main.go -help
