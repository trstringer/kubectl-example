DESTINATION_DIR=$$HOME/.local/bin

.PHONY: build
build:
	mkdir -p ./bin
	go build -o ./bin/kubectl-example

.PHONY: install
install: build
	cp ./bin/kubectl-example $(DESTINATION_DIR)

.PHONY: lint
lint:
	docker run --rm -v $$(pwd):/app -w /app golangci/golangci-lint:v1.46.2 golangci-lint run -v
