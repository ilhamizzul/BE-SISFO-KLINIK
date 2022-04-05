package pkg

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func PrintJson(data string) {
	src := []byte(data)

	dst := &bytes.Buffer{}
	if err := json.Indent(dst, src, "", "  "); err != nil {
		panic(err)
	}

	fmt.Println(dst.String())
}
