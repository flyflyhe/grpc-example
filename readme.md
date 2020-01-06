####生成对应go文件
````
protoc --proto_path=test --go_out=plugins=grpc:test test/test.proto
````
####测试
````
go test tests/test_test.go -v     
````

####生成php客户端
````
   ##生成client基类
   protoc --proto_path=user --php_out=php --grpc_out=php --plugin=protoc-gen-grpc=/data/grpc/bins/opt/grpc_php_plugin user/user.proto
   ##生成grpc信息
  protoc --proto_path=user --php_out=plugins=protoc-gen-grpc=/data/grpc/bins/opt/grpc_php_plugin:php user/user.proto
````