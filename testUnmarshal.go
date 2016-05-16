package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"log"
	"os"
	"proto/gen"
)


//unmarshaling error: proto: bad wiretype for field gen.DocV2.BackendId: got wiretype 2, want 0
//exit status 1
func main() {
	path := "/proto/data/1"
	str := read(path)
	// unmarshaling
	Payload := &gen.Payload{}
	fmt.Println(len(str))
	err := proto.Unmarshal(str, Payload)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}
	fmt.Println(Payload)
}

func read(path string) []byte {
	fi, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	return fd
}
