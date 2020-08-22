package ioutil

import (
	"bufio"
	"bytes"
	"io"
	"net"
)

// ioutil.ReadAll(...) is waiting up to 1 second to close connection.
// this reader counts opening and closing brackets and compare them.
// if bracket count is equal connection will be closed.
func ReadJson(conn net.Conn) (string, error) {
	reader := bufio.NewReader(conn)
	var key = int32(0x2B)
	var buffer bytes.Buffer
	var counter = 0
	var countOpenBrackets = 1
	var countCloseBrackets = 0
	buffer.WriteString("{")

	for {
		ba, err := reader.ReadByte()
		if err != nil {
			if err == io.EOF {
				break
			}
			return "", err
		}

		decryptedValue := string(int32(ba) ^ key)
		key = int32(ba)

		// ignore first 4 bytes
		if counter > 4 {
			buffer.WriteString(decryptedValue)
		}

		// count opened brackets
		if decryptedValue == "{" {
			countOpenBrackets++
		}

		// count closed bracket
		if decryptedValue == "}" {
			countCloseBrackets++
		}

		// stop reading connection if opened and closed brackets are equal
		if countOpenBrackets == countCloseBrackets {
			break
		}

		counter++
	}

	conn.Close()

	return buffer.String(), nil
}
