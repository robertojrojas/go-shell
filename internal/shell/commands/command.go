package commands

import "io"

type Command interface {
	Execute() error
	SetStdin(io.Reader)
	SetStdout(io.Writer)
	StdinPipe() (io.WriteCloser, error)
	StdoutPipe() (io.ReadCloser, error)
}

type BaseCommand struct {
	Stdin  io.Reader
	Stdout io.Writer
}

func (c *BaseCommand) SetStdin(input io.Reader) {
	c.Stdin = input
}

func (c *BaseCommand) SetStdout(output io.Writer) {
	c.Stdout = output
}

func (c *BaseCommand) StdinPipe() (io.WriteCloser, error) {
	r, w := io.Pipe()
	c.Stdin = r
	return w, nil
}

func (c *BaseCommand) StdoutPipe() (io.ReadCloser, error) {
	r, w := io.Pipe()
	c.Stdout = w
	return r, nil
}
