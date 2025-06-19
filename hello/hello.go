package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"time"

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

	var a, b, c = 1, "hello", false
	fmt.Printf("a=%d,b=%v,c=%v\n", a, b, c)

	const n = 500000000
	const d = 3e20 / n
	fmt.Println(int64(d))
	fmt.Println(math.Sin(n))

	i := 1
	for i <= 3 {
		fmt.Println(i)
		i += 1
	}

	for j := 10; j < 16; j++ {
		fmt.Println(j)
	}

	for i := range 10 {
		fmt.Printf("out with for .. range loop: ->%d\n", i+1)
	}

	count := 0
	for {
		if count == 100 {
			break
		}
		count += 1
		if count%10 == 0 {
			fmt.Printf("%v. ->loop mod 10 \n", count)
		} else if count%5 == 0 {
			fmt.Printf("%v. ->loop mod 5 \n", count)
		} else {
			fmt.Printf("%v. ->loop not mod 10 not mod 5 \n", count)
		}
	}

	dev_typs := []string{
		"PC", "APP", "WEB",
	}

	for k, v := range dev_typs {
		fmt.Println(k, v)
		switch k {
		case 0:
			fmt.Printf("[%v]===>Desktop: %v\n", time.Now(), v)
		case 1:
			fmt.Printf("[%v]===>Mobile: %v\n", time.Now(), v)
		case 2:
			fmt.Printf("[%v]===>Browser: %v\n", time.Now(), v)
		}

	}

	greetings.WhatAmI(true)
	greetings.WhatAmI(9999)
	greetings.WhatAmI("hello")
	greetings.Test()

	//arrays test
	arr1 := [5]int{}
	fmt.Println(arr1)
	for _i := range 5 {
		fmt.Println("===>", _i)
		fmt.Println(_i, math.Pi,
			math.Round(math.Pi),
			float64((_i+1)*1e7)*math.Pi,
			int64(math.Round((float64((_i+1)*1e9) * math.Pi))))
		arr1[_i] = int(math.Round(float64((_i+1)*1e9) * math.Pi))

		fmt.Println(arr1[_i])
	}
	fmt.Println(arr1, len(arr1))

	//arr2 := [...]string{"a", "b", "c"} //[...] not support append
	arr2 := []string{"a", "b", "c"} //[] support append
	s := "HelloWorld"
	l := len(arr2)
	for k, v := range s {
		fmt.Println(arr2, k, v, l)
		arr2 = append(arr2, string(v))
	}
	fmt.Println("===>", arr2)

	var double_arr [2][3]int
	for i := range 2 {
		for j := range 3 {
			double_arr[i][j] = (i + j + 1) * 8
		}
	}
	fmt.Println(double_arr)

}
