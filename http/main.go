package main

import (
	"io"
	"fmt"
	"net/http"
	"os"
)

type logWriter struct {

}

func main(){
	resp, err := http.Get("https://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	
	lw := logWriter{}

	io.Copy(lw, resp.Body)
}

func (logWriter) Write(bs []byte) (int, error){
	fmt.Println(string(bs))
	fmt.Println("Just wrote this:", len(bs))
	return len(bs), nil
}