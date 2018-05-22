package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	// http.Get will be returned a *Response:Struct and Erorr:Error
	// https://golang.org/pkg/net/http/#Client.Get
	res, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	/* 'Response' represents the response from an HTTP request.
	 * (Status, StatusCode, Header, BODY, etc)
	 *
	 * Body       is a io.ReadCloser type -> https://golang.org/pkg/net/http/#Response
	 * ReadCloser is a Interface type, containts : Reader & Closer interface -> https://golang.org/pkg/io/#ReadCloser
	 * Reader     is a Interface type, containts Read(p []byte) (n int, err error) -> https://golang.org/pkg/io/#Reader
	 *
	 */

	bs := make([]byte, 99999)
	res.Body.Read(bs)
	fmt.Println(string(bs))
}
