# udppeer Punch

## Motivation and Goal

***
## Deployment 

### Requirements 
- Needs a public accessible Server 
- Needs two clients in different networks, each protected by a firewall
- Executable needs to be build for different os architectures
```Bash
# Build a executable for local architecture 
go build 
```

### Usage

Server: 
```Bash
./udpeerPunch s
```

Client 1: 
```Bash
./udpeerPunch c <ip_of_server>:<port_of_server>
```

Client 2:
```Bash
./udpeerPunch c <ip_of_server>:<port_of_server>
```

## Credits & Collaborators