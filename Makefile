bl:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -v -o bin/bark-send-darwin

bl_linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o bin/bark-send-linux

