.PHONY: build

build:
	GO_ENABLED=0 go build -a -installsuffix cgo -ldflags '-w -s -extldflags "-static"' -o paper-vault .
