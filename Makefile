all:
	go build loader.go
	go build test.go
.PHONY: all
.DEFAULT: all