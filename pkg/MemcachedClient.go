package pkg

import (
	"github.com/monsieurtib/memcached-driver/pkg/clustering"
	"github.com/monsieurtib/memcached-driver/pkg/message"
)

type MemcachedClient struct {
	cluster clustering.ICluster
}

func NewClient() *MemcachedClient {
	var cluster clustering.ICluster
	cluster = clustering.CreateCluster(clustering.EndPoint{Address: "127.0.0.1", Port: "11211"})
	return &MemcachedClient{
		cluster: cluster,
	}
}

func (client *MemcachedClient) Get(key string) (result []byte, err error) {
	node := client.cluster.Locate(key)
	return node.SendCommand(message.GET, key, nil)
}

func (client *MemcachedClient) Set(key string, body []byte) (result []byte, err error) {
	node := client.cluster.Locate(key)
	return node.SendCommand(message.SET, key, body)
}
