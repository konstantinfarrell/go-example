run:
	go run cmd/go-example/main.go

test:
	find . -name "*test.go" -type f | go test | xargs 
