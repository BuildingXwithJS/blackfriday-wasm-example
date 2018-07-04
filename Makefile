build:
	GOOS=js GOARCH=wasm go1.11beta1 build -o lib.wasm lib.go
