# Compile in linux and output executable binary file that can run in windows
cd ../../cmd
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ./GoWebPractic-windows.exe ./main.go
echo "build finish"
