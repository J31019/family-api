all:run

run:
	go run main.go -addr 127.0.0.1:8080

build:
	rm ./release/api.exe
	GOOS=windows GOARCH=amd64 go build -o ./release/api.exe