package main

import (
	"bytes"
	"io"
	"os"
	"os/exec"
)

func main() {
	c0 := exec.Command("cat", "/Users/b4b4r07/.vimrc")
	c1 := exec.Command("pygmentize", "-O", "style=solarized", "-f", "console256", "-g")
	c2 := exec.Command("nkf")

	r, w := io.Pipe()
	c1.Stdout = w
	c2.Stdin = r

	var b2 bytes.Buffer
	c2.Stdout = &b2

	c1.Start()
	c2.Start()
	c1.Wait()
	w.Close()
	c2.Wait()
	io.Copy(os.Stdout, &b2)
	//あはは
}
