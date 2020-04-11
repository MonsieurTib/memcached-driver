package cluster

import (
	"github.com/monsieurtib/memcached-driver/requests"
	"io"
	"net"
)

type ClusterNode interface {
	ExecuteCommand(cmd requests.CommandOpCode, key string, body []byte)
}

type Node struct {
	connection io.ReadWriteCloser
	alive      bool
}

func CreateNode(endPoint string) (node *Node, err error) {
	connection, e := net.Dial("tcp", endPoint)
	if e != nil {
		println("connexion OK")
	}

	return &Node{
		connection: connection,
		alive:      true,
	}, e
}

func (node *Node) ExecuteCommand(cmd requests.CommandOpCode, key string, body []byte) {

	request := requests.Request{
		OpCode: cmd,
		Cas:    0,
		Opaque: 1,
		Key:    []byte(key),
		Body:   body,
	}
	var data = request.GetBytes()

	/*for i, datum := range data {
		println(i,datum)
	}*/

	_, err := node.connection.Write(data)
	if err != nil {
		println("write failed")
	} else {
		println("rock n roll!!")
	}
	node.ReadResponse()

}
func (node *Node) ReadResponse() {

	buffer := make([]byte, 24)
	_, err := io.ReadFull(node.connection, buffer)
	if err != nil {
		println("read failed")
	}
	for i, b := range buffer {
		println(i, b)
	}

}
