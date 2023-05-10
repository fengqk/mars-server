set GOOS=windows
set GOARCH=amd64
cd ./../
go build server.go
copy /y server.exe .\bin
del server.exe