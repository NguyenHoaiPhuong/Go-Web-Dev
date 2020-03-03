# Example #1

```
protoc pb/helloworld.proto --go_out=plugins=grpc:.

go run ./client/main.go

go run ./server/main.go
```
