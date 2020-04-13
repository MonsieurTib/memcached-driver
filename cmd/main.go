package main

import (
	"github.com/monsieurtib/memcached-driver/pkg"
)

func main() {

	var client = pkg.NewClient()
	client.Set("hello", []byte("world"))
	response, err := client.Get("hello2")
	if err == nil {
		println("GET RESULT", string(response))
	}
}
