BASE_FILES=main.go init.go globals.go revshell.go
BASE_OUT=revshell
LINUX_FILES=constants_linux.go
LINUX_OUT=${BASE_OUT}.out
WINDOWS_FILES=constants_windows.go
WINDOWS_OUT=${BASE_OUT}.exe

build_linux:
	GOOS=linux go build . ${LINUX_OUT}

build_windows:
	GOOS=windows go build . -o ${WINDOWS_OUT}

build_all: build_linux build_windows