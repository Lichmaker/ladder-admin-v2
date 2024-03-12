#!/bin/bash

# go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
# go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

# 获取脚本文件的相对路径
relative_path="$0"

# 获取脚本文件的绝对路径
absolute_path=$(readlink -f "$relative_path")
current_dir=$(dirname "$absolute_path")
echo "cd $current_dir"
cd $current_dir

mkdir -p ../v2ray-manager/protobuf/manager
mkdir -p ../server/protobuf/manager
mkdir -p ../server/protobuf/vpsproxy
protoc --proto_path=src --go_out=../v2ray-manager/protobuf/manager --go-grpc_out=../v2ray-manager/protobuf/manager --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative  manager.proto
protoc --proto_path=src --go_out=../server/protobuf/manager --go-grpc_out=../server/protobuf/manager --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative  manager.proto
protoc --proto_path=src --go_out=../server/protobuf/vpsproxy --go-grpc_out=../server/protobuf/vpsproxy --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative  vpsproxy.proto