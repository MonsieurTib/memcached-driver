package main

import (
	"github.com/monsieurtib/memcached-driver/pkg"
	"testing"
)

func TestGet(t *testing.T) {
	var client = pkg.NewClient()
	get, _ := client.Get("bose")
	t.Log(string(get))
}
