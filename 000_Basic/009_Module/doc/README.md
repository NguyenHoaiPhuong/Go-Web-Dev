# GO MODULE

A module is a collection of Go packages stored in a file tree with a go.mod file at its root. The go.mod file defines the module’s module path, which is also the import path used for the root directory, and its dependency requirements, which are the other modules needed for a successful build. Each dependency requirement is written as a module path and a specific semantic version.

As of Go 1.11, the go command enables the use of modules when the current directory or any parent directory has a go.mod, provided the directory is outside $GOPATH/src. (Inside $GOPATH/src, for compatibility, the go command still runs in the old GOPATH mode, even if a go.mod is found. See the go command documentation for details.) Starting in Go 1.13, module mode will be the default for all development.

Purpose of this article:

- Creating a new module: ```go mod init```
- Adding a dependency.
- Upgrading dependencies.
- Adding a dependency on a new major version.
- Upgrading a dependency to a new major version.
- Removing unused dependencies.

# References:

- https://blog.golang.org/using-go-modules