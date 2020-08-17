package crypto

import (
	"bytes"
	"log"
)

type Decryptor struct {
	key        int32
	printDebug bool
}

func NewDecryptor(printDebug bool) *Decryptor {
	return &Decryptor{
		key:        int32(0x2B),
		printDebug: printDebug,
	}
}

func (d *Decryptor) Decrypt(message []byte) string {
	var buffer bytes.Buffer
	buffer.WriteString("{")

	for pos, char := range message {
		value := int32(char) ^ d.key
		d.key = int32(char)

		if d.printDebug {
			log.Printf("read character %d, %d, %s\n", int32(char), value, string(value))
		}

		if pos > 4 {
			buffer.WriteString(string(value))
		}
	}

	return buffer.String()
}
