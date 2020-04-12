package clustering

import (
	"fmt"
	"github.com/monsieurtib/memcached-driver/message"
	"io"
	"net"
)

type IClusterNode interface {
	ExecuteCommand(cmd message.CommandOpCode, key string, body []byte)
}

type Node struct {
	connection io.ReadWriteCloser
	alive      bool
}

func (node *Node) Start(endPoint EndPoint) {
	connection, e := net.Dial("tcp", endPoint.Address+":"+string(endPoint.Port))
	if e != nil {
		println("connexion OK")
	}
	node.connection = connection

}

/*func CreateNode(endPoint EndPoint) *Node {
	address := endPoint.Address+":"+string(endPoint.Port)
	println(address)
	connection, e := net.Dial("tcp", address)
	if e != nil {
		println("connexion OK")
	}
	return &Node{
		connection: connection,
		alive:      true,
	}
}*/

func (node *Node) ExecuteCommand(cmd message.CommandOpCode, key string, body []byte) {

	request := message.Request{
		OpCode: cmd,
		Cas:    0,
		Opaque: 1,
		Key:    []byte(key),
		Body:   body,
	}

	var data = request.GetBytes()
	_, err := node.connection.Write(data)
	if err != nil {
		println("write failed")
	} else {
		println("rock n roll!!")
	}
	node.ReadResponse()

}
func (node *Node) ReadResponse() {

	println("ReadResponse")
	buffer := make([]byte, message.HEADER_LENGTH)
	_, err := io.ReadFull(node.connection, buffer)
	if err != nil {
		println("read failed")
	}
	response := message.Response{}
	response.TryParseHeader(buffer)

	bodyBuffer := make([]byte, response.TotalBodyLength)

	if _, err := io.ReadFull(node.connection, bodyBuffer); err != nil {
		fmt.Println("error:", err)
	} else {
		response.TryParseBody(bodyBuffer)
	}
	println("GET RESULT", string(response.Body))

}
