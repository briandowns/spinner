GO = go

.PHONY: deps
deps: go.mod

go.mod:
	go mod init
	go mod tidy

.PHONY: test
test:
	$(GO) test -v -cover ./...

.PHONY: check
check:
	if [ -d vendor ]; then cp -r vendor/* ${GOPATH}/src/; fi

.PHONY: clean
clean:
	$(GO) clean
