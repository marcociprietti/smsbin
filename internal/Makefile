PROJECTNAME=smsbin

GOBASE=$(shell pwd)
GOBIN=$(GOBASE)/bin


go-build:
	@echo "  > Building binary"
	@rm -r $(GOBIN) || true
	@go build -o $(GOBIN)/$(PROJECTNAME)

go-clean:
	@echo "  > Cleaning build cache"
	@go clean

go-get:
	@echo "  > Checking missing dependencies"
	@go get

go-compile: go-clean go-get go-build

go-test:
	@echo "  > Starting tests"
	@go test ./...
