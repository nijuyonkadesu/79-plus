package greeting

import (
	"errors"
	"fmt"
	"math/rand"
)

func Okay(msg string) (string, error) {
	if msg == "" {
		return msg, errors.New("empty bhai")
	}

	response := fmt.Sprintf(randomFormat(), msg, msg)
	return response, nil
}

func randomFormat() string {
    formats := []string{
        "Hi, %v. Welcome!",
        "Great to see you, %v!",
        "Hii, %T, %+v, welc",
    }
    return formats[rand.Intn(len(formats))]
}

func AppendBytes(slice []byte, data ...byte) []byte {
    m := len(slice)
    n := m + len(data)
    if cap(slice) < n {
        newSlice := make([]byte, (n+1)*2)
        copy(newSlice, slice)
    }
    slice = slice[0:n]
    copy(slice[m:n], data)
    return slice
}

// only functions with Caps could be imported ???
// [2]string / [...]string = array. 
// []string = slice, and that is what most used
// copy => returns the number of elements copied.
// ioutil

// array points to the entire file!!, so copy []byte before returning it
// len, cap, copy, append, delete
// byte is universal?? interesting...

// check further reading at the bottom: https://go.dev/blog/slices-intro

