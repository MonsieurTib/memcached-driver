package clustering

import (
	"fmt"
	"github.com/monsieurtib/memcached-driver/pkg/message"
	"io"
	"net"
)

type INode interface {
	SendCommand(cmd message.CommandOpCode, key string, body []byte) (result []byte, err error)
	Start(endPoint EndPoint)
}

type Node struct {
	connection io.ReadWriteCloser
	alive      bool
}

func NewNode(endPoint EndPoint) *Node {
	connection, e := net.Dial("tcp", endPoint.Address+":"+string(endPoint.Port))
	if e != nil {
		println("connexion OK")
	}
	return &Node{
		connection: connection,
	}
}

func (node *Node) Start(endPoint EndPoint) {

	connection, e := net.Dial("tcp", endPoint.Address+":"+string(endPoint.Port))
	if e != nil {
		println("connexion OK")
	}
	node.connection = connection
}

func (node *Node) SendCommand(cmd message.CommandOpCode, key string, body []byte) (result []byte, err error) {

	request := message.Request{
		OpCode: cmd,
		Cas:    0,
		Opaque: 1,
		Key:    []byte(key),
		Body:   body,
	}

	var data = request.GetBytes()

	_, error := node.connection.Write(data)
	if error != nil {
		println("write failed")
		return nil, error
	}
	return node.readResponse()

}
func (node *Node) readResponse() (result []byte, err error) {

	println("ReadResponse")
	buffer := make([]byte, message.HEADER_LENGTH)
	_, error := io.ReadFull(node.connection, buffer)
	if error != nil {
		println("read failed")
		return nil, error
	}
	response := message.Response{}
	if response.TryParseHeader(buffer) && response.TotalBodyLength > 0 {
		bodyBuffer := make([]byte, response.TotalBodyLength)
		if _, err := io.ReadFull(node.connection, bodyBuffer); err != nil {
			fmt.Println("error:", err)
		} else {
			response.TryParseBody(bodyBuffer)
		}
		return response.Body, nil
	}
	return nil, nil
}
