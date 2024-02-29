# webui starter for golang

## build
```bash
go get github.com/webui-dev/go-webui/v2/@2.4.1
./setup.sh

# or
make once
```

```bash
go build -ldflags -s -o bin/open-browser main.go
# -ldflags -s required for mac arm64
# or 
make build
```

## run
```bash
./bin/open-browser
# or
make run
```

## Reference
- [webui](https://github.com/webui-dev/webui)
- [go-webui](https://github.com/webui-dev/go-webui)