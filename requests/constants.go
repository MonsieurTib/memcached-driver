package requests

const (
	MAGIC_REQUEST  = 0x80
	MAGIC_RESPONSE = 0x81
	HEADER_LENGTH  = 24
)

type CommandOpCode uint8

const (
	GET = CommandOpCode(0x00)
	SET = CommandOpCode(0x01)
)
