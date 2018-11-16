package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestHttpGet(t *testing.T){
	fmt.Println("test http get method")
	resp, err := http.Get("http://127.0.0.1:8081/test/testhandler")
	if err != nil {
		fmt.Println(err)
	}

	defer  resp.Body.Close()
	bytes, e := ioutil.ReadAll(resp.Body)

	if e != nil {
		fmt.Println(e)
	}

	fmt.Println(string(bytes))
}
