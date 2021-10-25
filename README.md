# Chitty-Chat

## How to run

First run the Server, then Clients

### Server

To start the server, run ´server.go´ in its directory with:
> `go run .`

The server can also be started with arguments to specify e.g. which port it has to be run on.
This can be done by writing:

> `go run . -port <port>`

E.g. to run the server on port 5000 write:

> `go run . -port 5000`

If no argument is provided it will start on a standard port.

### Client

To start a client, run `client.go` in its directory with:
> `go run .`

The client can also be started with arguments to specify e.g.

- The display name of the client in chat
- The password for the client
- The ip address of the server the client connects to
- The port the client have to connect through

> `go run . -name <name> -password <password> -ip <address> -port <port>`

E.g. to create a client with the name "Agata", with password "123", ip address of server being 192.168.0.1 on port 5000, write:
> `go run . -name Agata -password 123 -ip 192.168.0.1 -port 5000`

If no arguments are provided it will pick a standard name, password, and port. The ip address will be set to "localhost".

### Script (Windows only)

Chitty-Chat can also be run with a PowerShell script for easy creation of server and clients.

To run the predefined script run:
> `run.ps1`

You can also specify how many clients to connect to the server by writing:
> `run.ps1 -p <number of clients>`

If no argument is provided it will start the server with 4 clients.

---

## Mini Project 2 - Chitty-Chat

### Description

You have to implement Chitty-Chat a distributed system, that is providing a chatting service, and keeps track of logical time using Lamport Timestamps.

We call clients of the Chitty-Chat service Participants.

Clients in Chitty-Chat can Publish a valid chat message at any time the wish.

A valid message is a string of UTF-8 encoded text with a maximum length of 128 characters.

A client publishes a message by making an RPC call Publish() to Chitty-Chat.

The Chitty-Chat service has to Broadcast every published message, together with the current Lamport timestamp, to all participants in the system, by using the RPC call Broadcast().

When a client receives a broadcasted message, it has to write the message and the current Lamport timestamp to the log

Chat clients can join or drop out at any time.

A "Participant X  joined Chitty-Chat at Lamport time L" message is broadcast when client X joins to all Participant, including the new Participant.

A "Participant X left Chitty-Chat at Lamport time L" message is broadcast when Participant X leaves to all remaining Participants.

### Technical Requirements

1. Use gRPC for all message passing between nodes
2. Use Golang to implement the service and clients
3. Every client has to be deployed as a separate processes
4. Log all service calls (Publish, Broadcast, ...) using the log package
5. Demonstrate that the system can be started with at least 3 client nodes
6. Demonstrate that a client node can join the system
7. Demonstrate that a client node can leave the system
8. Optional: All elements of the Chitty-Chat service are deployed as Docker containers

### Hand-in requirements

1. Hand in a single report in a pdf file
2. Describe your system architecture - do you have a server-client architecture, peer to peer, or something else?
3. Describe what  RPC methods are implemented: Publish(), Broadcast(), any other ?
4. Describe how you have implemented calculation of the Lamport timestamps
5. Provide a diagram, that traces a sequence of RPC calls together with the Lamport timestamps, that corresponds to a chosen sequence of interactions: Client X joines, Client X Publishes, ..., Client X leaves. Include documentation (system logs) in your appendix.
6. Provide a link to a Git repo with your source code in the report
7. Include system logs, that document the requirements are met, in the appendix of your report

---
