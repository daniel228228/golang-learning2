build_server:
	go build -race -o bin/server cmd/main.go

run_server:
	./bin/server

server: build_server run_server

clean:
	del /Q /F *.log
	rmdir /Q /S bin