package main

import (
	"fmt"
	"log"
	"math/rand"

	"example.com/greetings"
	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello")
	fmt.Println(quote.Go())

	log.SetPrefix("greetings: ")
	log.SetFlags(5)
	msg, err := greetings.Hello("GoLand")
	if nil != err {
		log.Fatal(err)
	}
	fmt.Println(msg)

	fmt.Println(rand.Intn(20))

	formats := []string{
		"aaaaa, %v",
		"bbbbb, %v",
		"ccccc, %v",
	}

	fmt.Println(len(formats))
	str_fmt := formats[rand.Intn(len(formats))]
	result := fmt.Sprintf(str_fmt, "micro")
	fmt.Println(result)

	//use map func
	names := []string{"microsoft", "intel", "amd"}
	msgs, err := greetings.Hellos(names)
	if nil != err {
		log.Fatal(err)
	}
	fmt.Println(msgs)

	for _, msg := range msgs {
		fmt.Println(msg)
	}

	for i, msg := range msgs {
		str := fmt.Sprintf("key-> : [%v], value-> : [%v]", i, msg)
		fmt.Println(str)
	}
}
