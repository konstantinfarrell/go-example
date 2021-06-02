TEST_FILES = 

run:
	go run cmd/go-example/main.go

test:
	for f in $$(find . -name "*test.go" -type f) ; do \
		go test $${f} ; \
	done

mock:
	mockery --all
