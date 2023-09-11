BINARY=snog
.DEFAULT_GOAL := install


run:
	go run main.go

install:
	go build -o snog
	mv $(BINARY) /usr/local/bin
	echo "Installation complete"

build:
	go build
