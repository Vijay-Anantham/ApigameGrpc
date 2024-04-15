# Hello Grpc..
The repo serves as in grpc experimentation and learning archive
Here in this application we try to have one grpc server and multiple client interacting with it.
- server
    - Records the max score of a player and update it realtime to a dashboard client
- players
    - The players roll for a number and update the server if a number is taken
- server
    - the roller will get the available pool from the serer and allocates number to a player.
    - This should not allot repeated numbers.

## GRPC

### running protoc compiler to generate go program
The following command can be used to generate go program from the protofiles
```bash
protoc --go_out=./game --go_opt=paths=source_relative \
    --go-grpc_out=./game --go-grpc_opt=paths=source_relative \
    ./protofiles/game.proto
```
## Resources
[Medium Docs on HTTP2](https://cabulous.medium.com/http-2-and-how-it-works-9f645458e4b2)
[Protobuf guide to generate go functions](https://protobuf.dev/reference/go/go-generated/)
[go package from protobuf](https://github.com/protocolbuffers/protobuf-go)
[example for developing basic protobuf](https://github.com/grpc/grpc-go/blob/master/examples/)
[Generate code reference from grpc](https://grpc.io/docs/languages/go/generated-code/)
[different grpc communication definition](https://grpc.io/docs/languages/go/basics/)
[noted on http/grpc](https://medium.com/geekculture/how-to-create-grpc-microservices-with-jpa-b3e804b4d91e)