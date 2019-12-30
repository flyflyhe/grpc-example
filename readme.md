####生成对应go文件
````
protoc --proto_path=test --go_out=plugins=grpc:test test/test.proto
````
####测试
````
go test tests/test_test.go -v     
````