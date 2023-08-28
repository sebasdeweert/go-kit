# Include and export .env file (the `-` prevents errors if no .env file is present)
-include .env
export

prepare:
	go get -u github.com/golang/mock/mockgen
	go get -u github.com/golang/dep/cmd/dep
	go get -u golang.org/x/tools/cmd/goimports
	go get -u github.com/go-swagger/go-swagger/cmd/swagger
	go get -u github.com/mgechev/revive
	make deps
	make mocks
	make deps

deps:
	GOPRIVATE=github.com/sebasdeweert
	go mod tidy

mocks:
	rm -rf mocks
	find . -name *_mocks.go -exec rm -f {} +
	mkdir -p mocks
	$$GOPATH/bin/mockgen -package=mocks github.com/sebasdeweert/go-kit/wrappers/http Client > mocks/http_mocks.go
	$$GOPATH/bin/mockgen -package=mocks net/http Handler > mocks/http_handler_mocks.go
	$$GOPATH/bin/mockgen -package mocks github.com/sebasdeweert/go-kit/encoding Encoder > mocks/encoding_mocks.go
	$$GOPATH/bin/mockgen -package mocks github.com/go-redis/redis Cmdable,Pipeliner > mocks/redis_mocks.go
	$$GOPATH/bin/mockgen -package mocks github.com/sirupsen/logrus Hook > mocks/logrus_mocks.go

fmt:
	$$GOPATH/bin/goimports -w $$(find . -name "*.go" | grep -v vendor | uniq)

unit:
	go test -v -short ./...

cover:
	go test ./... -v -race -coverprofile=coverage.temp.txt -covermode=atomic
	cat coverage.temp.txt | grep -v "_mocks.go" > coverage.txt
	rm coverage.temp.txt
	go tool cover -html=coverage.txt -o coverage.html
	go tool cover -func coverage.txt

lint:
	$$GOPATH/bin/revive -config vendor/github.com/sebasdeweert/go-kit/revive.toml -formatter friendly $$(find . -name "*.go" | grep -v vendor | grep -v service | grep -v mocks | uniq)

vet:
	go vet $$(go list ./...)
