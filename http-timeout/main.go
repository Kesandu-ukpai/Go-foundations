package main

import (
	"fmt"
	"net/http"
	"time"
)

//create the client like in the doc

func main() {

	client := http.Client{
		Timeout: 1 * time.Second,
	}
	resp, err := client.Get("https://api.github.com")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	fmt.Println(resp.Status)

}
