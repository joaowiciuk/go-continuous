package diff

import "github.com/joaowiciuk/go-continuous/pkg/file"

type Differ interface {
	Diff(branchA, branchB string) ([]*file.ProjectFile, error)
}
