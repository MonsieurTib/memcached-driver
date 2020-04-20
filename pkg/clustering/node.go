package clustering

import (
	"fmt"
	"github.com/monsieurtib/memcached-driver/pkg/message"
	"io"
	"net"
	"sync"
)

type INode interface {
	SendCommand(cmd message.CommandOpCode, key string, body []byte) (result []byte, err error)
	Start(endPoint EndPoint)
}

type Node struct {
	rw    io.ReadWriteCloser
	alive bool
	mux   sync.Mutex
}

func (node *Node) Start(endPoint EndPoint) {

	connection, e := net.Dial("tcp", endPoint.Address+":"+string(endPoint.Port))
	if e != nil {
		println("connexion OK")
	}
	node.rw = connection
}

func (node *Node) SendCommand(cmd message.CommandOpCode, key string, body []byte) (result []byte, err error) {

	node.mux.Lock()
	defer node.mux.Unlock()
	request := message.Request{
		OpCode: cmd,
		Cas:    0,
		Opaque: 1,
		Key:    []byte(key),
		Body:   body,
	}

	var data = request.GetBytes()

	_, error := node.rw.Write(data)
	if error != nil {
		println("write failed")
		return nil, error
	}
	response, err := node.readResponse()
	return response, err

}

func (node *Node) readResponse() (result []byte, err error) {

	buffer := make([]byte, message.HEADER_LENGTH)
	_, error := io.ReadFull(node.rw, buffer)
	if error != nil {
		println("read failed")
		return nil, error
	}

	response := message.Response{}
	if response.TryParseHeader(buffer) && response.TotalBodyLength > 0 {
		bodyBuffer := make([]byte, response.TotalBodyLength)
		if response.TotalBodyLength > 0 {
			if _, err := io.ReadFull(node.rw, bodyBuffer); err != nil {
				fmt.Println("error:", err)
			} else {
				response.TryParseBody(bodyBuffer)
			}
		}
		return response.Body, nil
	}
	return nil, nil
}
