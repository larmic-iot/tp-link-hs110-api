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
	encryptedMessage := createByteArray(len(message))

	for _, x := range e.encrypt(message) {
		if e.printDebug {
			fmt.Printf("write %d to byte %b\n", x, byte(x))
		}

		encryptedMessage = append(encryptedMessage, byte(x))
	}

	return encryptedMessage
}

func createByteArray(messageLength int) []byte {
	var encryptedMessage []byte
	encryptedMessage = append(encryptedMessage, 0)
	encryptedMessage = append(encryptedMessage, 0)
	encryptedMessage = append(encryptedMessage, 0)
	encryptedMessage = append(encryptedMessage, 0)
	binary.BigEndian.PutUint32(encryptedMessage, uint32(messageLength))
	return encryptedMessage
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
