package main

import (
	"github.com/monsieurtib/memcached-driver/pkg"
	"strconv"
	"sync"
	"testing"
	"time"
)

func TestParrallism(t *testing.T) {

	getFunc := func(client *pkg.MemcachedClient, from, to int, group *sync.WaitGroup, t *testing.T) {
		for i := from; i < to; i++ {
			index := strconv.Itoa(i)
			response, err := client.Get("hello_" + index)
			if err == nil {

				if string(response) != "world_"+index {
					t.Fatal("FAILED : hello_" + index)
				}
			} else {
				t.Fatal("FAILED", err.Error())
			}
		}
		group.Done()
	}

	var client = pkg.NewClient()

	for i := 0; i < 2000; i++ {
		index := strconv.Itoa(i)
		client.Set("hello_"+index, []byte("world_"+index))
	}

	var waitgroup sync.WaitGroup
	waitgroup.Add(2)

	println("start")
	start := time.Now()
	go getFunc(client, 0, 1000, &waitgroup, t)
	go getFunc(client, 1001, 2000, &waitgroup, t)

	waitgroup.Wait()
	duration := time.Since(start)
	println(duration.Milliseconds())
	println("complete")

}

func BenchmarkSetGet(t *testing.B) {

	var client = pkg.NewClient()
	client.Set("hello", []byte("world"))
	response, err := client.Get("hello")
	if err == nil {
		if string(response) != "world" {
			t.Fatal("FAILED")
		}
	} else {
		t.Fatal("FAILED", err.Error())
	}
}
