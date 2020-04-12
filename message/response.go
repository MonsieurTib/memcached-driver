package message

import "encoding/binary"

type Response struct {
	Test            string
	OpCode          CommandOpCode
	Cas             uint64
	Opaque          uint32
	KeyLength       uint16
	ExtraLength     byte
	DataType        byte
	Status          ResponseStatus
	Flags           uint32
	TotalBodyLength uint32
	Body            []byte
}

func (response *Response) TryParseHeader(data []byte) bool {

	if len(data) < HEADER_LENGTH {
		panic("unexpected header")
	}
	if data[0] != MAGIC_RESPONSE {
		panic("unexpected response")
	}
	/*for i, datum := range data {
		println(i, datum)
	}*/
	response.OpCode = CommandOpCode(data[1])
	response.KeyLength = binary.BigEndian.Uint16(data[2:])
	response.ExtraLength = byte(data[4])
	response.DataType = byte(data[5])
	response.Status = ResponseStatus(binary.BigEndian.Uint16(data[6:8]))
	response.TotalBodyLength = binary.BigEndian.Uint32(data[8:12])
	response.Opaque = binary.BigEndian.Uint32(data[12:16])
	response.Cas = binary.BigEndian.Uint64(data[16:])

	return true
}
func (response *Response) TryParseBody(data []byte) bool {

	println("BODY**************************")

	response.Flags = binary.BigEndian.Uint32(data[1:])
	index := response.KeyLength + uint16(response.ExtraLength)

	response.Body = make([]byte, response.TotalBodyLength-(uint32(response.KeyLength)+uint32(response.ExtraLength)))
	copy(response.Body, data[index:])
	/*for i, b := range response.Body {
		println(i,b)
	}*/
	return true
}
