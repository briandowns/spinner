GO = go

.PHONY: test
test:
	$(GO) test -v -cover ./...

.PHONY: clean
clean:
	$(GO) clean
	rm -f example

.PHONY: example
example: clean
	go build -o example ./example

