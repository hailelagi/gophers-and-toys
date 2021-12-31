package main

import (
	"github.com/hailelagi/gophers-and-toys/rock_paper_scissors/rps"
)

/*
Core challenge is working/learning TCP/IP sockets in go
if client A | B sends a message it waits for 30 secs before timeout
Then when the value of A & B is received game logic compares
and res sent to both clients

++++++++++++   socket   ++++++++++++   socket   +++++++++++++
+ Client A +   <--->    +  Server +    <---->   + Client B +
++++++++++++   message  ++++++++++++   message  +++++++++++++
*/

/*
TODO:
--prototype
1. spin up a TCP server and client
2. spin up a cli session intro asking a terminal session to connect
3. implement interface rules of game - should be simple enough

--improvements
4. Simple Vue.js client with emoji as input.
A new browser window should start a brand new game session
To connect to a new session you must possess a game ID
5. Implement game ID// game history - how to persist?
*/

func main() {
	// 1. spin up a TCP server and client
	rps.Server()
}
