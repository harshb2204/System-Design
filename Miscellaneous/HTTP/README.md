## HTTP Protocol
2 machines are connected within the network and they want to talk to each other. They establish a TCP connection between them.
But how would device B understand what device A is talking about. -> There needs to be a common language -> protocol
![](/diagrams/protocol.png)

We defined the specification (protocol) for a server that does Add, Sub ops.

1. space separated
2. first word is the command
3. subsequent word are arguments
4. in response, one word is result
This is req res(client server model)
HTTP is also a protocol where

1. client asks for something from server
2. server responds

But the question is, how does it ask ? → how server would know what client wants ?

In HTTP the requests and responses are encoded as "messages" → Human Readable

`GET /foo HTTP/1.1 \r\n`

literal text we send over TCP to the server

A server is a "web server" if it understands HTTP protocol → given the message ... it interprets what the client wants


