


build:
	GOOS=js GOARCH=wasm go build -v -o main.wasm .

run:
	GOOS=js GOARCH=wasm go run . -exec="$(go env GOROOT)/misc/wasm/go_js_wasm_exec"

valgrind:
	goexec 'http.ListenAndServe(":8080", http.FileServer(http.Dir(".")))'
