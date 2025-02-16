package dao

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"strings"
	"unicode"
)

// There is a bug with Postgres & libxml2 with UTF-8 special characters.
// At the time of this writing, this occurs with 3 of the 50 title XML exports.
// https://www.postgresql.org/message-id/18274-98d16bc03520665f%40postgresql.org

func removeSpecialCharacters(s string) string {
	var b strings.Builder
	b.Grow(len(s))
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsNumber(r) || unicode.IsSpace(r) {
			b.WriteRune(r)
		}
	}
	return b.String()
}

func scrubXML(b []byte) []byte {
	reader := bytes.NewReader(b)

	var output bytes.Buffer

	decoder := xml.NewDecoder(reader)
	encoder := xml.NewEncoder(&output)

	for {
		token, err := decoder.Token()
		if err != nil {
			break
		}

		switch t := token.(type) {
		case xml.CharData:
			cleaned := removeSpecialCharacters(string(t))
			encoder.EncodeToken(xml.CharData(cleaned))
		default:
			encoder.EncodeToken(t)
		}
	}

	if err := encoder.Flush(); err != nil {
		fmt.Printf("Error flushing XML encoder: %v", err)
	}

	return output.Bytes()
}
