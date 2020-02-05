# For Python

```
python -m grpc_tools.protoc -I. --python_out=./python --grpc_python_out=./python nltk_service.proto
```

# For Golang

```
protoc -I. nltk_service.proto --go_out=plugins=grpc:golang
```