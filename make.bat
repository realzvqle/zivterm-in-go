@echo off
go build -o bin/zivterm.exe src/main.go src/cmdcheck.go src/checkError.go src/commands.go
start bin/zivterm.exe
rem go build src/*.go