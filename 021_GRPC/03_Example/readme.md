# Arithmetic service

```
protoc pb/arithmetic.proto --go_out=plugins=grpc:.

go run ./client/main.go

go run ./server/main.go
```