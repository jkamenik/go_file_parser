package main

import (
	"io"
	"os/exec"
)


func xzReader(r io.Reader) io.ReadCloser {
	rpipe, wpipe := io.Pipe()

	cmd := exec.Command("xz", "--decompress", "--stdout")
	cmd.Stdin = r
	cmd.Stdout = wpipe

	go func() {
			err := cmd.Run()
			wpipe.CloseWithError(err)
	}()

	return rpipe
}
