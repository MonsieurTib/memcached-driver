package message

import "encoding/binary"

type Request struct {
	OpCode     CommandOpCode
	Cas        uint64
	Opaque     uint32
	Key        []byte
	Body       []byte
	Flags      uint32
	Expiration uint32
	VBucket    uint16
}

func (req *Request) setHeader() []byte {

	extraLength := 0
	if req.Expiration != 0 || req.Body != nil {
		extraLength = 8
	}
	data := make([]byte, HEADER_LENGTH+extraLength+len(req.Key)+len(req.Body))
	data[0] = MAGIC_REQUEST
	data[1] = byte(req.OpCode)
	binary.BigEndian.PutUint16(data[2:4], uint16(len(req.Key)))
	data[4] = byte(extraLength)
	data[5] = byte(0)
	binary.BigEndian.PutUint16(data[6:8], req.VBucket)
	binary.BigEndian.PutUint32(data[8:12], uint32(len(req.Body)+len(req.Key)+extraLength))
	binary.BigEndian.PutUint32(data[12:16], req.Opaque)
	binary.BigEndian.PutUint64(data[16:], req.Cas)
	copy(data[HEADER_LENGTH+extraLength:], req.Key)
	return data
}

func (req *Request) GetBytes() []byte {
	data := req.setHeader()
	if len(req.Body) > 0 {
		copy(data[HEADER_LENGTH+8+len(req.Key):], req.Body)
	}
	return data
}
