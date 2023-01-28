DESTINATION_DIR=$$HOME/.local/bin

.PHONY: build
build:
	mkdir -p ./bin
	go build -o ./bin/kubectl-example

.PHONY: install
install: build
	cp ./bin/kubectl-example $(DESTINATION_DIR)
