## Communication

### The usual communication
![](/diagrams/reqres.png)

---

### Short Polling
![](/diagrams/shortpolling.png)

The client repeatedly sends requests at fixed intervals to check if new data is available.

#### Examples
- Refreshing a live cricket score every 5 seconds.
- Checking whether a server has finished provisioning.

#### Disadvantages:

- Many unnecessary HTTP requests.
- Repeated request/response overhead.

---

### Long Polling
![](/diagrams/longpolling.png)
The client sends a request, but the server does not respond immediately. Instead, it keeps the connection open until:
- new data is available, or
- a timeout occurs.
Once the client receives the response, it immediately sends another long-poll request.

#### Examples
- Waiting for an EC2 instance to become ready.
- Waiting for the next ball in a cricket match instead of checking every few seconds.
- short polling sends response right away
- long polling sends response only when done
  - connection kept open for the entire duration

---

## WebSockets
![](/diagrams/websockets.png)

### Advantages
- realtime data transfer
- low communication overhead

### Applications
- realtime communication
- stock market ticker
- live experiences
- multiplayer games

---

## Server Sent Events
![](/diagrams/sse.png)

### Examples
- Deployment log streaming: You triggered a deployment on a server, it goes on a worker machine and it is deploying. The logs are streamed to the frontend. Client does not send anything after setting up connection. 
- stock market ticker
- Advertisments
- updates on twitter feeds


### Prototype
```java
@RestController
public class EventController {

    @GetMapping(value = "/events",
            produces = MediaType.TEXT_EVENT_STREAM_VALUE)
    public Flux<String> streamEvents() {

        return Flux.interval(Duration.ofSeconds(1))
                .map(i -> "Current Time : " + LocalTime.now());
    }
}

```

Flux is a class from Project Reactor, which is the reactive programming library used by Spring WebFlux.
`A stream of zero or more values that arrive over time.`

`Flux<String>` means "I will keep sending Strings."
```java
Flux.interval(Duration.ofSeconds(1))
```
This creates a Flux that emits a number every second.
Then
```java
.map(i -> "Current Time : " + LocalTime.now())
```
transforms each emitted number into a string.
---

## How realtime apps are generally structured?
![](/diagrams/realtimeapps.png)
