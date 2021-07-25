# Go Continuous
CLI tool to help testing only relevant packages when integrating branches in a project

## How it works
go-continuous uses git to check for changed files between branches, assembles a package dependency tree based on them and calls go test only on packages whose files changed or depends on packages whose files did so.