package git

import (
	"fmt"

	"github.com/joaowiciuk/go-continuous/pkg/file"
)

type Git struct {
}

type Option func(*Git) error

func New(options ...Option) (*Git, error) {
	git := new(Git)
	for _, option := range options {
		err := option(git)
		if err != nil {
			return nil, fmt.Errorf("%w: %s", ErrInit, err.Error())
		}
	}
	return git, nil
}

func (g *Git) Diff(branchA, branchB string) ([]*file.ProjectFile, error) {
	panic("not implemented")
}
