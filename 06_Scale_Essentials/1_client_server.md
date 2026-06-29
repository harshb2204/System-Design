# Client Server Model

Most common way for two machines to talk to each other.

* **Client (C)**: Demands
* **Server (S)**: Does the job

Examples of requests:
- give me my profile info
- delete the post I made
- spin up one server

The communication happens over the common network connecting the two.
Two protocols to exchange data: **TCP** and **UDP**.
*(we almost 99.99% times use TCP)*

## Some important properties of TCP

1. TCP connection requires 3-way handshake for setup.
2. TCP connection requires 2-way handshake for teardown.
3. TCP connection does not break immediately after data is exchanged.
   - breaks because of network interruption.
   - breaks because server/client initiated it.

Hence connection remains open... almost 'forever'.

## Protocol over TCP

TCP does not dictate what data can be sent over it.
Common format agreed upon by client & server is called a protocol: **HTTP**

HTTP is just a format that client and server understands.
* You can define your own format, and make
  1. your client send data in it
  2. your server parses & processes it (e.g., `"GET_K\n"`)

## Properties of HTTP 1.1

There are many versions of it - HTTP 1.1 / HTTP 2 / HTTP 3
HTTP 1.1 is the most commonly used one.

1. For client and server to talk over HTTP 1.1, they need to establish a TCP connection.
2. Connection is typically terminated once response is sent to client.
3. almost new connection for every request / response.
   -> little expensive
4. Hence people pass `"Connection: Keep-alive"` header which tells client and server to not close the connection.
   * depends if server follows it or not

## WebSocket

WebSockets are meant to do bi-directional communication.
* **Key feature**: Server can proactively send data to client, without client asking for it.

Because there is no need for setting up TCP every single time, we get really low latency in communication.

Any where you need "realtime", "low latency" communication for your end user over the internet, think about WebSockets.
- e.g., Chat, Realtime likes on live stream, Stock market ticks
