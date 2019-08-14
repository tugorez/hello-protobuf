package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/tugorez/protobuf/pb"
	"io/ioutil"
	"log"
)

func main() {
	fname := "address_book.txt"
	p := pb.Person{
		Id:    1234,
		Name:  "John Doe",
		Email: "jdoe@example.com",
		Phones: []*pb.Person_PhoneNumber{
			{Number: "555-4321", Type: pb.Person_HOME},
		},
	}
	people := []*pb.Person{&p}
	book := &pb.AddressBook{People: people}

	// Lets write the file
	out, err := proto.Marshal(book)
	if err != nil {
		log.Fatalln("Failed to encode address book")
	}
	if err := ioutil.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("Failed to write address book:", err)
	}

	// Lets read the previously written file
	in, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalln("Failed to read the adress book")
	}
	new_book := &pb.AddressBook{}
	if err := proto.Unmarshal(in, new_book); err != nil {
		log.Fatalln("Failed to decode the adress book")
	}
	fmt.Println(new_book)
}
