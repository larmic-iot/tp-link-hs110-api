package crypto

import (
	"encoding/binary"
	"log"
)

const (
	cryptoKey = int32(0xAB)
)

type Encryptor struct {
	nextCryptoKey int32
	printDebug    bool
}

func NewEncryptor(printDebug bool) *Encryptor {
	return &Encryptor{
		nextCryptoKey: cryptoKey,
		printDebug:    printDebug,
	}
}

func (e *Encryptor) Encrypt(message string) []byte {
	encryptedMessage := createByteArray(len(message))

	for _, x := range e.encrypt(message) {
		if e.printDebug {
			log.Printf("write %d to byte %b\n", x, byte(x))
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
		value := char ^ e.nextCryptoKey
		e.nextCryptoKey = value
		buffer = append(buffer, value)

		if e.printDebug {
			log.Printf("character %c, %d starts at byte position %d\n", char, value, pos)
		}
	}

	return buffer
}
