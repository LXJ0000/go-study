## 安装gRPC
`go get google.golang.org/grpc@latest`

## 安装Protocol Buffers v3
[Win](https://github.com/protocolbuffers/protobuf/releases/download/v3.20.1/protoc-3.20.1-win64.zip)  
[Linux](https://github.com/protocolbuffers/protobuf/releases/download/v3.20.1/protoc-3.20.1-linux-x86_64.zip)

> Win 需加入环境变量

## 安装插件
```bash
## go语言插件
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
## grpc插件
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

```
## CHECK
```
protoc --version
protoc-gen-go --version
protoc-gen-go-grpc --version
```

## GEN
```
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative pb/hello.proto
```