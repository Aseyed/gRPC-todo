package main

import (
	"os"
	"io/ioutil"
	"github.com/golang/protobuf/proto"
	"encoding/binary"
	context "context"
	"fmt"
	"log"
	"net"
	"todo/todo"

	grpc "google.golang.org/grpc"
)



func main() {
	srv := grpc.NewServer()
	var tasks taskServer
	todo.RegisterTasksServer(srv, tasks)
	l, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalf("could not listen to 8888 %v", err)
	}
	log.Fatal(srv.Serve(l))
}
type taskServer struct {
}


type length int64
const (
	sizeOfLength = 8
	dbPath = "mydb.pb"
)
var endianness = binary.LittleEndian

// func  add(text string) error {
// 	task := &todo.Task {
// 		Text: text,
// 		Done: false,
// 	}

// 	b, err := proto.Marshal(task)
// 	if err != nil {
// 		return fmt.Errorf("Could not encode task: %v", err)
// 	}

// 	f, err := os.OpenFile(dbPath, os.O_WRONLY | os.O_CREATE | os.O_APPEND, 0666 )

// 	if err != nil {
// 		return fmt.Errorf("Could not open %s: %v", dbPath, err)
// 	}

// 	if err := binary.Write(f, endianness, length(len(b))); err != nil {
// 		return fmt.Errorf("Could not encode length of message %v", err)
// 	}
// 	_, err = f.Write(b)
// 	if err != nil {
// 		return fmt.Errorf("Could not write task to file: %v", err)
// 	}

// 	if err := f.Close(); err != nil {
// 		return fmt.Errorf("Could not close file %s: %v", dbPath, err)
// 	}

// 	// fmt.Println(task.GetDone())
// 	// fmt.Println(task.GetText())

// 	fmt.Println(proto.MarshalTextString(task))
// 	return nil
// }


func (s taskServer) List(ctx context.Context, void *todo.Void) (*todo.TaskList, error) {
	return nil, fmt.Errorf("not implemented")
}

func list() error {
	b, err := ioutil.ReadFile(dbPath)
	if err != nil {
		return fmt.Errorf("Could not read file %s: %v",dbPath, err)
	}

	for {
		if len(b) == 0 {
			return nil
		}else if len(b) < 4 {
			return fmt.Errorf("remaining odd %d bytes, what to do?", len(b))
		}

		var length int64
		if err := binary.Read(bytes.NewReader(b[:4])).Decode(&length); err != nil {
			return fmt.Errorf("Could not decode message length: %v", err)
		}
		b = b[4:]

		var task todo.Task
		if err := proto.Unmarshal(b[:length], &task); err != nil {
			return fmt.Errorf("Could not read task: %v", err)
		}
		b = b[length:]

		if task.Done {
			fmt.Printf ("[X]")
		}else{
			fmt.Printf("[ ]")
		}

		fmt.Printf(" %s\n",task.Text)
	}

}
