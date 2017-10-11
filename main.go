package main

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"github.com/ghodss/yaml"
)

func check(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func source() (io.Reader, error) {
	if len(os.Args) > 1 {
		f, err := os.Open(os.Args[1])
		if err != nil {
			return nil, err
		}
		return f, nil
	}

	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		return os.Stdin, nil
	}

	return nil, errors.New("no input source")
}

func main() {
	src, err := source()
	check(err)

	in, err := ioutil.ReadAll(src)
	check(err)

	d, err := yaml.JSONToYAML(in)
	check(err)

	fmt.Print(string(d))
}
