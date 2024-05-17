@echo off

set GOPROXY=https://mirrors.aliyun.com/goproxy/
set CGO_ENABLED=0
set GOOS=linux
set GOARCH=amd64

go build
