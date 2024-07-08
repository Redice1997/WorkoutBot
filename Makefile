.PHONY:
.SILENT:

build:
	go build -o ./.bin/bot cmd/telegram/main.go

run: build
	./.bin/bot