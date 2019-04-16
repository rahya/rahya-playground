


build:
	GOOS=js GOARCH=wasm go build -v -o main.wasm .

run:
	go run .

valgrind:
	goexec 'http.ListenAndServe(":8080", http.FileServer(http.Dir(".")))'
