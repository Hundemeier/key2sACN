#!/bin/bash
#make the assetFS filesystem that includes everything in the webgui/build folder into the executable
go-bindata-assetfs webgui/build/...

# build for following architectures:
GOOS=linux GOARCH=amd64 go build -o out/key2sACN_amd64
GOOS=linux GOARCH=386 go build -o out/key2sACN_386
GOOS=linux GOARCH=arm go build -o out/key2sACN_arm