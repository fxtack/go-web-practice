:: You can use this script to compile to get executable binary file that can run on linux
cd ../../cmd
set CGO_ENABLED=0
set GOOS=linux
set GOARCH=amd64
go build -o cmd/GoWebPractic-linux ./cmd/main.go

pause