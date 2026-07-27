## How servers handle multiple connections ?

TCP is the most reliable way for two machines to talk to each other over the network

Logical Entity that wraps communication specifics → SOCKET
* TCP
* UDP

What is a TCP server then ?
It is a simple process that runs in a machine that 'listens' to a port which understands TCP.
Any server that wants to talk to the server has to connect over the port and establish the connection. 
#### Step1: Start listening on the port
When your process starts pick a port and start listening to it.

```go

net.Listen("http", ":1729")
```


#### Step2: Wait for a client to connect
Invoke the "accept" system call and wait for a client to connect. (this is a blocking call and your server would not proceed until some client connects)

```go

listener.Accept()
```

#### Step3: Read the request and send the response
Once the connection is established 
- invoke the read system call to read the request(blocking)
- invoke the write system to call to send the response (blocking)
- close the connection
```go

Conn.Read()
Conn.Write()
Conn.Close()
```

#### Step 4: Do this over and over again

Put this entire thing in an infinite for loop ...

1. continuously waiting for client to connect
2. reading the request
3. writing the response
4. closing the connection



#### Step 5: Multiple request concurrently
Currently the code will accept one connection, process it, and then accept another.
Sequential execution and handling

#### Step 6: Parallelize the processing

Once client connects, fork a thread to process and respond

let main thread come back to 'accept' as quickly as possible

```
for {
    accept()
    ~~> process(conn)
}
```


### Improvements
1. Limiting the number of threads
2. add thread pool to save on thread creation time
3. connection timeout (What if a client connects, you are waiting for something, and the client never sends you the request)
4. TCP baclog queue config