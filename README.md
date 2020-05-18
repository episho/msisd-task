# msisd-task

## RPC
I choose gRPC because grpc does not support only Golang, it is supported by many languages. 
The fact that jsonRPC can support other languages but not the HTTP method limits it's application in real life. 
Therefore for production environment, we use gRPC to overcome this. 
It is developed in Protobuf serialised protocol and supports popular languages such as Python, Java and Golang. 
Therefore i wrote a producer that can be used in other services no matter if they are written in golang or php.

### How to run
- to run the server: cd cmd/ `go run server.go`
- to run the client: `go run client.go` (http://localhost:1212/msisdn/{input}) 
 
