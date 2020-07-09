GOPATH:=$(shell go env GOPATH)

.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64  go build -o apiServer main.go
	chmod +x apiServer
	mv apiServer ./build/

.PHONY: run
run:
	swag init # docs https://github.com/swaggo/swag
	go mod download
	bee run -gendoc=true -downdoc=true # docs https://beego.me/quickstart
	#go run main.go







