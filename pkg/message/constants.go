package message

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

type ResponseStatus uint16

const (
	SUCCESS           = ResponseStatus(0x00)
	KEY_NOT_FOUND     = ResponseStatus(0x01)
	KEY_EXISTS        = ResponseStatus(0x02)
	VALUE_TOO_LARGE   = ResponseStatus(0x03)
	INVALID_ARGUMENTS = ResponseStatus(0x04)
)
