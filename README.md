# Gochat

A simple chat program

## Building

```bash
git clone https://github.com/PioKozi/gochat.git
cd gochat
go build cmd/*
```

## Usage

When started, it will listen on a port

```bash
./main <port number>
```

For a client to send on a socket it must be introduced to it

```
/introduce <host>:<port>
```

## To do

- [ ] Make introduction to hosts easier, faster
- [ ] Can only receive messages from known hosts
- [ ] More commands for more functionality
- [ ] Encryption?
- [ ] Verification?
