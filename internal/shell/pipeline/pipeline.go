package pipeline

import "github.com/sanurb/go-shell/internal/shell/commands"

type Pipeline struct {
	Commands []commands.Command
}

func (p *Pipeline) AddCommand(cmd commands.Command) {
	p.Commands = append(p.Commands, cmd)
}

func (p *Pipeline) Execute() error {
	composite := &commands.CompositeCommand{
		Commands: p.Commands,
	}
	return composite.Execute()
}
