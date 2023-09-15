new:
	- go run ./cmd/main.go -new "$(FOLDER)"

new-git:
	- go run ./cmd/main.go -new "$(FOLDER)" -git "$(LINK)"

help:
	- go run ./cmd/main.go -help
