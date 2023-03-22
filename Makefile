BINPATH=./bin/
BINNAME=douyin-grab

ifeq ($(OS),Windows_NT)
	PLATFORM=windows
	BINNAME=douyin-grab.exe
else
	ifeq ($(shell uname),Darwin)
		PLATFORM=darwin
	else
		PLATFORM=linux
	endif
endif

all:
	go build -o $(BINPATH)$(BINNAME) main.go