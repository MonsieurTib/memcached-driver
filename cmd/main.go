package main

import "github.com/monsieurtib/memcached-driver/pkg"

func main() {

	var client = pkg.NewClient()
	client.Set("hello", []byte("world"))

	result, _ := client.Get("hello")

	println("********************************", string(result))

}
