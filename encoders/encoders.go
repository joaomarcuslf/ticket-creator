package encoders

// import (
// 	"errors"
// 	"math"
// 	"strings"
// )

import (
	b64 "encoding/base64"
)

func Encode(initialLink string) string {
	sEnc := b64.StdEncoding.EncodeToString([]byte(initialLink))

	return sEnc
}

func Decode(encodedLink string) (string, error) {
	sDec, _ := b64.StdEncoding.DecodeString(encodedLink)

	return string(sDec), nil
}
