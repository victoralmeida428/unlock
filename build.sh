#!/usr/bin/env bash

echo "Build linux"
go build -o unlock-app-linux



echo "Build windows"
GOOS=windows GOARCH=amd64 go build -o unlock-app.exe

echo "Compliação concluída"

echo "- unlock-app-linux (binário Linux)"
echo "- unlock-app.exe (binário Windows)"