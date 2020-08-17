package crypto

import (
	"encoding/binary"
	"fmt"
)

type Encryptor struct {
	key        int32
	printDebug bool
}

func NewEncryptor(printDebug bool) *Encryptor {
	encryptor := Encryptor{
		key:        int32(0xAB),
		printDebug: printDebug,
	}
	return &encryptor
}

func (e *Encryptor) Encrypt(message string) []byte {
	var text []byte
	text = append(text, 0)
	text = append(text, 0)
	text = append(text, 0)
	text = append(text, 0)
	binary.BigEndian.PutUint32(text, uint32(len(message)))

	for _, x := range e.encrypt(message) {
		valur := byte(x)

		if e.printDebug {
			fmt.Printf("write %d to byte %b\n", x, valur)
		}

		text = append(text, valur)
	}

	return text
}

func (e *Encryptor) encrypt(message string) []int32 {
	var buffer []int32

	for pos, char := range message {
		value := char ^ e.key
		e.key = value
		buffer = append(buffer, value)

		if e.printDebug {
			fmt.Printf("character %c, %d starts at byte position %d\n", char, value, pos)
		}
	}

	return buffer
}
