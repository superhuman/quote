package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/yuya-takeyama/argf"
)

const usage = `Usage: quote [-d] <file>

quote input suitable for pasting as javascript
`

func decode(reader io.Reader) error {
	var output string
	err := json.NewDecoder(reader).Decode(&output)
	if err != nil {
		return err
	}
	fmt.Println(string(output))
	return err
}

func encode(reader io.Reader) error {
	str, err := ioutil.ReadAll(reader)
	if err != nil {
		return err
	}
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetEscapeHTML(false)
	encoder.Encode(string(str))
	fmt.Println("")
	return nil
}

func run() error {
	var displayHelp bool
	var shouldDecode bool
	flag.BoolVar(&displayHelp, "h", false, "display help")
	flag.BoolVar(&shouldDecode, "d", false, "decode")
	flag.Parse()
	reader, err := argf.From(flag.Args())

	if err != nil {
		return err
	}

	if displayHelp {
		fmt.Println()

		return nil
	}

	if shouldDecode {
		return decode(reader)
	}

	return encode(reader)
}

func main() {
	err := run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
