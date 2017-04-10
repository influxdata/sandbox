# TICK Stack Sandbox on Windows

[Docker for Windows Homepage](https://www.docker.com/docker-windows "Docker for Windows")

There is a sandbox.cmd file that replaces the sandbox shell script for windows command line.

## What works

- Sandbox containers start and seem to work (make sure that your firewall allows port 445 connections)
- Enhanced docker-compose command so that sandbox does not use the directory name as part of the container names -> sandbox can be cloned to an arbitrarily named directory and the sandbox script still works :-)

## What does not work

- chronograf seems not to work, at least not with localhost:8888
- I have not yet found a replacement for /var/run/docker.sock docker API access

