all: build

build: clean
	GOARCH=wasm GOOS=js go build -o web/app.wasm ./cmd/html2goapp-frontend
	go run ./cmd/html2goapp-frontend

run:
	GOARCH=wasm GOOS=js go build -o web/app.wasm ./cmd/html2goapp-frontend
	go run ./cmd/html2goapp-frontend -serve

clean:
	rm -rf web/app.wasm
	rm -rf out
