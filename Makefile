.PHONY: build
build:
	go build
test:
	go test
cov:
	go test -coverprofile cover.out && go tool cover -html=cover.out -o cover.html && open cover.html
install:
	dep ensure
run:
	go build && ./go-stock-checker
