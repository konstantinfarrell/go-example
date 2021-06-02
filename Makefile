TEST_FILES = 

run:
	go run cmd/go-example/main.go

docker:
	docker build -t go-example .
	docker run -it --rm -p 9001:9001 go-example:latest

test:
	for f in $$(find . -name "*test.go" -type f) ; do \
		go test $${f} ; \
	done

mock:
	mockery --all
