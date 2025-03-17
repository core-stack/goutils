package structutils

import (
	"bytes"
	"encoding/json"
)

func Transcode(in, out any) {
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(in)
	json.NewDecoder(buf).Decode(out)
}
