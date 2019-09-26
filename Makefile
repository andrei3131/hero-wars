GAME = hero-wars
GAMEBIN = ./bin/${GAME}
CMDDIR = ./cmd
TESTDIR = ./test
VERSION = 0.0.1-dev

test:
	# go generate ./...
	go test -v github.com/andrei3131/hero-wars/...

build-all: build-mac build-linux

build:
	go build -o ${GAMEBIN} ${CMDDIR}

build-linux:
	GOOS=linux GOARCH=amd64 go build -ldflags "-s -X main.Version=${VERSION}" -v -o ${GAMEBIN}-linux-amd64 ${CMDDIR}

build-mac:
	GOOS=darwin GOARCH=amd64 go build -ldflags "-s -X main.Version=${VERSION}" -v -o ${GAMEBIN}-darwin-amd64 ${CMDDIR}

clean:
	rm -f ${GAMEBIN}
	rm -f ${GAMEBIN}-linux-amd64
	rm -f ${GAMEBIN}-darwin-amd64

install: build
	sudo install -d /usr/local/bin
	sudo install -c ./bin/${GAME} /usr/local/bin/${GAME}

uninstall:
	sudo rm /usr/local/bin/${GAME}


.PHONY: all test build uninstall clean