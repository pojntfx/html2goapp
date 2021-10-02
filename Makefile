# Variables
DESTDIR ?=
PREFIX ?= /usr/local
DST ?= out/cli/html2goapp-cli
WWWROOT ?= /var/www/html
WWWPREFIX ?= /html2goapp

all: build

# Build
build: build/cli build/pwa

build/cli:
	go build -o $(DST) ./cmd/html2goapp-cli

build/pwa:
	GOARCH=wasm GOOS=js go build -o web/app.wasm ./cmd/html2goapp-pwa
	go run ./cmd/html2goapp-pwa -prefix $(WWWPREFIX)
	cp -rf web/* out/pwa/web/web
	tar -cvzf out/pwa/html2goapp-pwa.tar.gz -C out/pwa/web .

# Install
install: install/cli install/pwa

install/cli:
	install -D -m 0755 $(DST) $(DESTDIR)$(PREFIX)/bin/html2goapp-cli

install/pwa:
	mkdir -p $(DESTDIR)$(WWWROOT)$(WWWPREFIX)
	cp -rf out/pwa/web/* $(DESTDIR)$(WWWROOT)$(WWWPREFIX)

# Uninstall
uninstall: uninstall/cli uninstall/pwa

uninstall/cli:
	rm $(DESTDIR)$(PREFIX)/bin/html2goapp-cli

uninstall/pwa:
	rm -rf $(DESTDIR)$(WWWROOT)$(WWWPREFIX)

# Run
run: run/cli run/pwa

run/cli: build
	$(DST)

run/pwa:
	GOARCH=wasm GOOS=js go build -o web/app.wasm ./cmd/html2goapp-pwa
	go run ./cmd/html2goapp-pwa -serve

# Clean
clean:
	rm -rf out web/app.wasm
