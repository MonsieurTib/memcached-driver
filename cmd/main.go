package main

import (
	"github.com/monsieurtib/memcached-driver/cluster"
	"github.com/monsieurtib/memcached-driver/requests"
)

func main() {
	var node cluster.ClusterNode
	node, _ = cluster.CreateNode("127.0.0.1:11211")
	node.ExecuteCommand(requests.SET, "pouet", []byte("zouzou"))

}
