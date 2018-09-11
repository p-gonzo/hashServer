# Simple Golang Hashing Server
## Description
### Simple Golang Hashing Server is an HTTP server written in Go using the standard library

#### Hashing

Strings sent via a `POST` request to the `/hash` endpoint with a form field named `password` are hashed using the SHA512 algorithm and encoded as a Base64 string. A response is served after a 5 second delay.

The server can handle multiple concurrent requests and will always return a hashed value after ~5 seconds.


#### Stats

A `GET` request can be made to the `/stats` endpoint and a JSON object will be returned with the total number of hashed passwords made since the server last started along with the average time to hash each password.

#### Graceful Shutdown

A request can be made to the `/shutdown` endpoint and the server will invoke a graceful shutdown without interrupting any active connections


## Installation

`go get github.com/p-gonzo/hashServer`

## Usage

Inside of the `$HOME/go/src/github.com/p-gonzo/hashServer` directory, run `./hashServer`

## Testing

Inside of the `$HOME/go/src/github.com/p-gonzo/hashServer/routes` directory, run `go test`