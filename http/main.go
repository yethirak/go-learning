package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type Resp struct {
	Status        string
	StatusCode    int
	Proto         string
	ProtoMajor    int
	ProtoMinior   int
	Header        http.Header
	Body          io.ReadCloser
	ContentLength int64
}

func main() {
	resp, err := http.Get("https://google.com")
	if err != nil {
		fmt.Println("Error: ", err)
	}

	// //raw code how to read body from raw response
	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))
	// resp.Body.Close()

	// usually you will go with compination of writer/reader to read data
	io.Copy(os.Stdout, resp.Body)
}
