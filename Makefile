GO := go

build:
	@echo todo

test:
	$(GO) test -v ./...

generate:
	bash ./hack/client_codegen.sh
