# gRPC-todo
#### This is a TODO app based on go lang and gRPC protocol

## pre-requirements
* Install go based on your system.
* Install requirements package with these comands:
    * `go get -u github.com/golang/protobuf/{proto,protoc-gen-go}`
    * `go get -u google.golang.org/grpc`


## Run the programm
* make proto file in go language
    * `protoc -I . todo.proto --go_out=plugins=grpc:.`    
* To run server
   * `cd ./cmd/server && go build` (`cd` used to server be in same folder ad db file)
   * `mv server ../.. && ./server`
* To run client
   * `cd ./cmd/todo`
   * `go run main.go list` (this option maybe not work; There is a problem in server side)
   * `go run main.go add SOMETEXT`

## See What is in db
* `hexdump -c mydb.pb`
* `cat mydb.pb | protoc --decode_raw`
