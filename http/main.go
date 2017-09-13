package main

import "net/http"
import "fmt"
import "os"
import "io"

type logWriter struct {
}

func main() {
	resp, err := http.Get("https://www.google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	lw := logWriter{}

	io.Copy(lw, resp.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("bytes length", len(bs))
	return len(bs), nil
}
