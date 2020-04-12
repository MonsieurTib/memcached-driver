package main

import (
	"github.com/monsieurtib/memcached-driver/clustering"
	"github.com/monsieurtib/memcached-driver/message"
)

func main() {

	/*var node clustering.ClusterNode
	node = clustering.CreateNode(clustering.EndPoint{Address: "127.0.0.1", Port: 11211})
	//node.ExecuteCommand(message.SET, "poueto", []byte("zouzou"))
	node.ExecuteCommand(message.GET, "large", nil)*/

	var cluster clustering.ClusterNodeLocator

	cluster = clustering.CreateCluster(clustering.EndPoint{Address: "127.0.0.1", Port: "11211"})
	node := cluster.Locate("poueto")
	node.ExecuteCommand(message.GET, "large", nil)
}
