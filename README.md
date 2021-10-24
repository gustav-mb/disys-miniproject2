# Mini Project 2—Chitty-Chat

## Description ##
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

## Technical Requirements ##
1. Use gRPC for all message passing between nodes
2. Use Golang to implement the service and clients
3. Every client has to be deployed as a separate processes
4. Log all service calls (Publish, Broadcast, ...) using the log package
5. Demonstrate that the system can be started with at least 3 client nodes 
6. Demonstrate that a client node can join the system
7. Demonstrate that a client node can leave the system
8. Optional: All elements of the Chitty-Chat service are deployed as Docker containers

## Hand-in requirements ##
1. Hand in a single report in a pdf file
2. Describe your system architecture - do you have a server-client architecture, peer to peer, or something else?
3. Describe what  RPC methods are implemented: Publish(), Broadcast(), any other ?
4. Describe how you have implemented calculation of the Lamport timestamps
5. Provide a diagram, that traces a sequence of RPC calls together with the Lamport timestamps, that corresponds to a chosen sequence of interactions: Client X joines, Client X Publishes, ..., Client X leaves. Include documentation (system logs) in your appendix.
6. Provide a link to a Git repo with your source code in the report
7. Include system logs, that document the requirements are met, in the appendix of your report
