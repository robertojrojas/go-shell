.PHONY: build
build:
	@go build -o go-shell cmd/main.go

.PHONY: run
run: build
	@./go-shell
