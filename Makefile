run: main.go
	build && ./judger

build: main.go
	go build -o judger main.go

test:  router_test.go
	sudo go test
