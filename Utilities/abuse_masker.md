# Designing Realtime Abuse Masker

Say, one video live stream is powered by one server.
All participants part of that live stream (max 100) are connected to the same server.



In this live stream, people can send text messages (broadcast).

We want to ensure abuses are masked.
`s***` / `f***`

## Understanding Architecture
![](/diagrams/livestream.png)

- live stream of creator is captured over RTMP
- participants are connected to the server over WebSocket
- participant can send message to all

### How socket.io works?
- Socket IO has notion of rooms.
- Simple Socket IO process is running on server.
- We create "room" for this live stream.
- Participants join the room.
- When message is sent to room, Socket IO sends to all.

## Abuse Dictionary

Assume that we have a list of words (abuses) stored in a text file on Blob storage like S3 (at a path).
Hence, when the live stream server starts up:
- downloads the abuse file
- loads it in memory -> *how and which data structure?*

## Masking the abuse

Given all the messages of one live stream will go through the same server, we have to find some logic that efficiently detects and masks.

### Approach 1: Tokenize and lookup

1. load all abuses in dictionary (hash table / set)
2. tokenize incoming text message
3. For each word, check if it is in abuse dict
4. if yes: mask and update the token
5. if no: copy the token

**This is a simple approach but not efficient**:
1. tokenize (text to list of tokens requires extra space)
2. string lookup is expensive

### Approach 2: Trie

We can build a trie out of abuses and in just `O(n)` traversal mask the abuses.

For each character in text we iterate through trie.
Every time we get `' '`, `','`, `'.'` etc we reset our trie iteration (to the top).

If the abuse is found and next char is EOS (End of String) or non-alphabet, we mask and write to the new string. Else the word is copied as is.

The final string is then sent / broadcasted to every participant.

## How this fits into the system?

We typically think of databases and all to hold and query, but that is very slow for this.
Plus, **no database exposes a Trie**.

So, what do we do?

### Really poor way to design this system
Have an "abuse masker" service that exposes an HTTP endpoint accepting the text and returns a masked version of it.
The web service will have the trie loaded in memory.
![](/diagrams/abusemaskerservice.png)

**Note**: Not everything needs to be a "service" -> *Think realistically.*

**Disadvantages of a separate service**:
- Making a network call for *every* text message will be massively slow.
  - Setting up a TCP connection every time requires a 3-way handshake.
  - Terminating the connection once we get a response requires a 2-way teardown.
- Even for a persistent connection, we still incur network I/O latency.

Hence, creating a separate web service to clean the text is a bad idea.


![](/diagrams/s3abuselist.png)
Live stream server itself can host the trie. It holds the trie. When the live stream server boots up it makes a call to s3 downloads the abuse list, builds a trie around it and starts serving req. Whenever user sends a message it checks for the abuse, it reconstructs another message (masking it).

## Configuring Abuse list

We need a way to update the list of abuses (for an internal team).

- we need a UI where abuses are listed
- we should update the abuses and push them to S3

We create an **Abuse Admin Service** to take care of this.
![](/diagrams/abuseadmin.png)
