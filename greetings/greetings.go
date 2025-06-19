package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {
	if "" == name {
		return name, errors.New("empty name")
	}
	msg := fmt.Sprintf(randomFormat(), name) //format output string with rand value
	return msg, nil
}

func Hellos(names []string) (map[string]string, error) {
	msgs := make(map[string]string) //init map
	for _, name := range names {
		msg, err := Hello(name) //call func Hello with arg name
		if nil != err {
			return nil, err
		}
		msgs[name] = msg //set map value in loop
	}
	return msgs, nil
}

func randomFormat() string {
	fmts := []string{
		"Dev->App : %v ",
		"Dev->Web : %v ",
		"Dev->Qt  : %v ",
	}

	n := rand.Intn(len(fmts)) //get rand number

	return fmts[n] //get string by rand index
}
