# go-todo
#### This is a TODO app based on go lang and gRPC protocol

## pre-requirements
* Install go based on your system.
* Install requirements package with these comands:
    * `go get -u github.com/golang/protobuf/{proto,protoc-gen-go}`
    * `go get -u google.golang.org/grpc`


## Run the programm
* `go install ./cmd/todo`
* `todo add SOMETEXT`
* `todo list`

## See What is in db
* `hexdump -c mydb.pb`
* `cat mydb.pb | protoc --decode_raw`