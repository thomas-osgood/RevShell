BASE_FILES=main.go init.go globals.go revshell.go
BASE_OUT=revshell
LINUX_FILES=constants_linux.go
LINUX_OUT=${BASE_OUT}.out
MAIN_COMMAND=go build -o
WINDOWS_FILES=constants_windows.go
WINDOWS_OUT=${BASE_OUT}.exe


# this builds the .out file for linux. 
#
# NOTE: the "." after ${LINUX_OUT} is not a typo.
build_linux:
	GOOS=linux ${MAIN_COMMAND} ${LINUX_OUT} .

# this builds the .exe file for windows. 
#
# NOTE: the "." after ${WINDOWS_OUT} is not a typo.
build_windows:
	GOOS=windows ${MAIN_COMMAND} ${WINDOWS_OUT} .

# this calls both build_linux and build_windows to
# build the binaries for both windows and linux.
build_all: build_linux build_windows