@echo off
go build demo.go
taskkill /F /IM demo.exe
start "demo" demo.exe web