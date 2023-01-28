.PHONY: build
build:
	mkdir -p ./bin
	go build -o ./bin/kubectl-example

.PHONY: install
install: build
	cp ./bin/kubectl-example $$HOME/.local/bin
