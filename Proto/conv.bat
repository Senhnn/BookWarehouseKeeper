@REM 使用 source_relative 则不会使用option go_package中指定的路径

@REM ServiceProto的生成
protoc --go_out=../ --go-grpc_out=../ ./ServiceProto.proto

@REM DBProto的生成
protoc --go_out=../ ./DBProto.proto