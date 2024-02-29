default: build

once:
	@echo "Once Setting up..."
	@go get github.com/webui-dev/go-webui/v2/@latest
	@./setup.sh

build:
	@echo "Building..."
	@go build -ldflags -s -o bin/open-browser main.go
# -ldflags -s required for mac arm64

run:
	@echo "Running..."
	@go run main.go

clean:
	@echo "Cleaning..."
	@rm -rf bin
