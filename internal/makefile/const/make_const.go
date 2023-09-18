package _const

const MakefileBase = `run:
	go run cmd/main.go

run_docker:
	docker-compose build && docker-compose up

make_migrations:
	migrate create -ext sql -dir ./migration -seq init

migrate:
	migrate -path ./migration -database 'postgres://$(USER):$(PASS)@localhost:8001/postgres?sslmode=disable' up

rollback:
	migrate -path ./migration -database 'postgres://$(USER):$(PASS)@localhost:8001/postgres?sslmode=disable' up

path_reset:
	export PATH=$PATH:$(go env GOPATH)/bin`
